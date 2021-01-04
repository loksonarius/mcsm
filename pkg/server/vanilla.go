package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/loksonarius/mcsm/pkg/config"
	"github.com/loksonarius/mcsm/pkg/config/presets"
)

type VanillaServer struct {
	Definition       ServerDefinition
	ServerBinaryPath string
	Configs          []config.ConfigFile
}

func NewVanillaServer(def ServerDefinition) Server {
	return &VanillaServer{
		Definition:       def,
		ServerBinaryPath: "minecraft_server.jar",
		Configs: []config.ConfigFile{
			presets.ServerPropertiesFromConfig(def.Configs),
			presets.EulaTxtFromConfig(def.Configs),
		},
	}
}

func init() {
	serverInitializers[VanillaInstall] = NewVanillaServer
}

func (vs *VanillaServer) Install() error {
	versions, err := vs.getAvailableVersions()
	if err != nil {
		return err
	}

	// assume we want latest version
	targetVersion := versions[0]
	reqVersion := vs.Definition.Install.Version
	if reqVersion != "latest" {
		found := false
		for _, v := range versions {
			if v.ID == reqVersion {
				targetVersion = v
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("no vanilla server version %s", reqVersion)
		}
	}

	versionURL, err := vs.getVersionDownloadURL(targetVersion)

	// vanilla doesn't support mods nor plugins, so we're done at this point
	return downloadFileToPath(versionURL, vs.ServerBinaryPath)
}

func (vs *VanillaServer) Configure() error {
	for _, cfg := range vs.Configs {
		if err := cfg.Write(); err != nil {
			return err
		}
	}

	return nil
}

type versionMetadata struct {
	ID          string
	Type        string
	URL         string
	Time        string
	ReleaseTime string
}

func (vs VanillaServer) getAvailableVersions() ([]versionMetadata, error) {
	var versions []versionMetadata
	versionsURL, err := url.Parse("https://launchermeta.mojang.com/mc/game/version_manifest.json")
	if err != nil {
		return versions, err
	}

	resp, err := http.Get(versionsURL.String())
	if err != nil {
		return versions, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return versions, err
	}

	var parsedResponse struct {
		Latest struct {
			Release  string
			Snapshot string
		}
		Versions []versionMetadata
	}
	err = json.Unmarshal(body, &parsedResponse)

	// we only care about SemVer releases
	semverFormat := regexp.MustCompile(`^\d+\.\d+(\.\d)?((-rc|-pre)\d+)?$`)
	for _, v := range parsedResponse.Versions {
		if semverFormat.MatchString(v.ID) {
			versions = append(versions, v)
		}
	}
	return versions, err
}

func (vs VanillaServer) getVersionDownloadURL(version versionMetadata) (string, error) {
	versionURL, err := url.Parse(version.URL)
	if err != nil {
		return "", err
	}

	resp, err := http.Get(versionURL.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	type arbitraryJson map[interface{}]interface{}
	var parsedResponse struct {
		// let json ignore the excess keys, we only care about downloads
		Downloads map[string]struct {
			Sha1 string
			Size string
			URL  string
		}
	}
	err = json.Unmarshal(body, &parsedResponse)
	entry, ok := parsedResponse.Downloads["server"]
	if !ok {
		return "", fmt.Errorf(
			"no server download found for version %s", version.ID)
	}

	if _, err := url.Parse(entry.URL); err != nil {
		return "", err
	}

	return entry.URL, nil
}

func (vs *VanillaServer) Versions() ([]string, error) {
	var resp []string
	versions, err := vs.getAvailableVersions()
	if err != nil {
		return resp, err
	}

	for _, v := range versions {
		resp = append(resp, v.ID)
	}

	return resp, nil
}

func (vs *VanillaServer) Run() error {
	if _, err := os.Stat(vs.ServerBinaryPath); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("server not installed, refusing to run")
		}

		return err
	}

	path, err := exec.LookPath("java")
	if err != nil {
		return err
	}
	args := javaArgs(vs.ServerBinaryPath, vs.Definition.Run)
	cmd := exec.Command(path, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
		Pgid:    0,
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	fmt.Printf("Starting '%s %s'\n", path, strings.Join(args, " "))
	if err := cmd.Start(); err != nil {
		return err
	}

	go func() {
		io.Copy(in, os.Stdin)
	}()

	cmdErrChan := make(chan error, 1)
	go func() {
		cmdErrChan <- cmd.Wait()
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	timeoutChan := make(chan int, 1)

	for {
		select {
		case err := <-cmdErrChan:
			return err
		case <-sigChan:
			if _, err := in.Write([]byte("stop\n")); err != nil {
				// ru-roh
			}

			go time.AfterFunc(6*time.Second, func() { timeoutChan <- 1 })
		case <-timeoutChan:
			if err := cmd.Process.Kill(); err != nil {
				// woah, what a hardy process
			}

			return nil
		}
	}
}
