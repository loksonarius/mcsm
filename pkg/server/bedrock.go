package server

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"syscall"

	"github.com/loksonarius/mcsm/pkg/config"
	"github.com/loksonarius/mcsm/pkg/config/presets"
)

const BedrockInstall = InstallKind("bedrock")

type BedrockServer struct {
	Definition       ServerDefinition
	ServerBinaryPath string
	Configs          []config.ConfigFile
}

func NewBedrockServer(def ServerDefinition) Server {
	return &BedrockServer{
		Definition:       def,
		ServerBinaryPath: "bedrock_server",
		Configs: []config.ConfigFile{
			presets.BedrockServerPropertiesFromConfig(def.Configs),
		},
	}
}

func init() {
	registerServer(BedrockInstall, NewBedrockServer)
}

func (bs *BedrockServer) InstallOpts() InstallOpts {
	return bs.Definition.Install
}

func (bs *BedrockServer) RuntimeOpts() RuntimeOpts {
	return bs.Definition.Run
}

func (bs *BedrockServer) ConfigFiles() []config.ConfigFile {
	return bs.Configs
}

func (bs *BedrockServer) Install() error {
	versions, err := bs.Versions()
	if err != nil {
		return err
	}

	// assume we want latest version
	targetVersion := versions[len(versions)-1]
	reqVersion := bs.Definition.Install.Version
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
			return fmt.Errorf("no bedrock server version %s", reqVersion)
		}
	}

	versionURL := fmt.Sprintf(
		"https://minecraft.azureedge.net/bin-linux/bedrock-server-%s.zip",
		targetVersion,
	)

	zipDir := filepath.Dir(bs.ServerBinaryPath)
	zipPath := filepath.Join(zipDir, "bedrock.zip")
	if err := downloadFileToPath(versionURL, zipPath); err != nil {
		return err
	}

	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	ignore_paths := []string{
		"whitelist.json",
		"permissions.json",
	}
	for _, f := range r.File {
		info := f.FileHeader.FileInfo()
		mode := info.Mode()

		ignore := false
		for _, ip := range ignore_paths {
			if f.Name == ip {
				// check if file already exists
				if _, err := os.Stat(ip); err != nil {
					if os.IsNotExist(err) { // doesn't exist
						break
					}

					return err // exists but might have strict permissions
				}

				// exists, so ignore
				ignore = true
				break
			}
		}
		if ignore {
			continue
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path, err := filepath.Abs(filepath.Join(zipDir, f.Name))
		if err != nil {
			return err
		}

		if info.IsDir() {
			if err := os.MkdirAll(path, mode); err != nil {
				return err
			}
		} else {
			out, err := os.Create(path)
			if err != nil {
				return err
			}
			defer out.Close()

			if err := out.Chmod(mode); err != nil {
				return err
			}

			_, err = io.Copy(out, rc)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (bs *BedrockServer) Versions() ([]string, error) {
	emptyResponse := []string{}

	addr := "https://raw.githubusercontent.com/loksonarius/bedrock-server-versions/main/versions.json"
	var parsedResponse struct {
		LastUpdated string
		Versions    []string
	}
	if err := httpGetAndParseJSON(addr, &parsedResponse); err != nil {
		return emptyResponse, err
	}

	return parsedResponse.Versions, nil
}

func (bs *BedrockServer) Run() error {
	if _, err := os.Stat(bs.ServerBinaryPath); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("server not installed, refusing to run")
		}

		return err
	}

	path, err := filepath.Abs(bs.ServerBinaryPath)
	if err != nil {
		return err
	}

	args := []string{filepath.Base(path)}

	ld_path := filepath.Dir(bs.ServerBinaryPath)
	env := append(os.Environ(), fmt.Sprintf("LD_LIBRARY_PATH=%s", ld_path))

	fmt.Printf("Starting '%s'\n", bs.ServerBinaryPath)
	return syscall.Exec(path, args, env)
}
