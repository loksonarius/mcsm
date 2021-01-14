package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/loksonarius/mcsm/pkg/config"
	"github.com/loksonarius/mcsm/pkg/config/presets"
)

const PaperInstall = InstallKind("paper")

type PaperServer struct {
	Definition       ServerDefinition
	ServerBinaryPath string
	PluginDirectory  string
	Configs          []config.ConfigFile
}

func NewPaperServer(def ServerDefinition) Server {
	return &PaperServer{
		Definition:       def,
		ServerBinaryPath: "minecraft_server.jar",
		PluginDirectory:  "plugins",
		Configs: []config.ConfigFile{
			presets.ServerPropertiesFromConfig(def.Configs),
			presets.EulaTxtFromConfig(def.Configs),
		},
	}
}

func init() {
	registerServer(PaperInstall, NewPaperServer)
}

func (ps *PaperServer) Install() error {
	versions, err := ps.Versions()
	if err != nil {
		return err
	}

	// assume we want latest version
	targetVersion := versions[len(versions)-1]
	reqVersion := ps.Definition.Install.Version
	if reqVersion != "latest" {
		found := false
		for _, v := range versions {
			if v == reqVersion {
				targetVersion = v
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("no paper server version %s", reqVersion)
		}
	}

	versionURL, err := ps.getVersionDownloadURL(targetVersion)
	if err != nil {
		return err
	}

	err = downloadFileToPath(versionURL, ps.ServerBinaryPath)
	if err != nil {
		return err
	}

	for _, source := range ps.Definition.Install.Plugins {
		source.storeToDirectory(ps.PluginDirectory)
	}

	return nil
}

func (ps *PaperServer) Configure() error {
	for _, cfg := range ps.Configs {
		if err := cfg.Write(); err != nil {
			return err
		}
	}

	return nil
}

type paperVersionsList struct {
	ProjectName   string
	VersionGroups []string
	Versions      []string
}

type paperVersionInfo struct {
	ProjectName string
	Version     string
	Builds      []int
}

type paperBuildInfo struct {
	ProjectName string
	Version     string
	Build       int
	Downloads   struct {
		Application struct {
			Name   string
			Sha256 string
		}
	}
}

func (ps *PaperServer) getVersionDownloadURL(v string) (string, error) {
	var emptyURL string

	versionURL, err := url.Parse(fmt.Sprintf(
		"https://papermc.io/api/v2/projects/paper/versions/%s",
		v,
	))
	if err != nil {
		return emptyURL, err
	}

	resp, err := http.Get(versionURL.String())
	if err != nil {
		return emptyURL, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return emptyURL, err
	}

	var versionInfo paperVersionInfo
	err = json.Unmarshal(body, &versionInfo)
	if err != nil {
		return emptyURL, err
	}

	build := versionInfo.Builds[len(versionInfo.Builds)-1]
	buildURL, err := url.Parse(fmt.Sprintf(
		"https://papermc.io/api/v2/projects/paper/versions/%s/builds/%d",
		v,
		build,
	))
	if err != nil {
		return emptyURL, err
	}

	resp, err = http.Get(buildURL.String())
	if err != nil {
		return emptyURL, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return emptyURL, err
	}

	var buildInfo paperBuildInfo
	err = json.Unmarshal(body, &buildInfo)
	if err != nil {
		return emptyURL, err
	}

	download := buildInfo.Downloads.Application.Name
	downloadURL, err := url.Parse(fmt.Sprintf(
		"https://papermc.io/api/v2/projects/paper/versions/%s/builds/%d/downloads/%s",
		v,
		build,
		download,
	))
	if err != nil {
		return emptyURL, err
	}

	return downloadURL.String(), nil
}

func (ps *PaperServer) Versions() ([]string, error) {
	var emptyResponse []string
	versionsURL, err := url.Parse("https://papermc.io/api/v2/projects/paper")
	if err != nil {
		return emptyResponse, err
	}

	resp, err := http.Get(versionsURL.String())
	if err != nil {
		return emptyResponse, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return emptyResponse, err
	}

	var parsedResponse struct {
		VersionGroups []string
		Versions      []string
	}
	err = json.Unmarshal(body, &parsedResponse)

	return parsedResponse.Versions, err
}

func (ps *PaperServer) Run() error {
	binaryPath := ps.ServerBinaryPath
	runOpts := ps.Definition.Run
	return runJavaServer(binaryPath, runOpts)
}
