package server

import (
	"fmt"

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
			presets.BukkitYmlFromConfig(def.Configs),
			presets.SpigotYmlFromConfig(def.Configs),
			presets.PaperYmlFromConfig(def.Configs),
		},
	}
}

func init() {
	registerServer(PaperInstall, NewPaperServer)
}

func (ps *PaperServer) InstallOpts() InstallOpts {
	return ps.Definition.Install
}

func (ps *PaperServer) RuntimeOpts() RuntimeOpts {
	return ps.Definition.Run
}

func (ps *PaperServer) ConfigFiles() []config.ConfigFile {
	return ps.Configs
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

func (ps *PaperServer) getVersionDownloadURL(v string) (string, error) {
	var emptyURL string
	addr := "https://papermc.io/api/v2/projects/paper"

	var versionInfo struct{ Builds []int }
	addr = fmt.Sprintf("%s/versions/%s", addr, v)
	err := httpGetAndParseJSON(addr, &versionInfo)
	if err != nil {
		return emptyURL, err
	}

	build := versionInfo.Builds[len(versionInfo.Builds)-1]
	addr = fmt.Sprintf("%s/builds/%d", addr, build)
	var buildInfo struct {
		Downloads struct{ Application struct{ Name string } }
	}

	err = httpGetAndParseJSON(addr, &buildInfo)
	if err != nil {
		return emptyURL, err
	}

	download := buildInfo.Downloads.Application.Name
	addr = fmt.Sprintf("%s/downloads/%s", addr, download)

	return addr, nil
}

func (ps *PaperServer) Versions() ([]string, error) {
	var emptyResponse []string
	addr := "https://papermc.io/api/v2/projects/paper"
	var parsedResponse struct {
		VersionGroups []string
		Versions      []string
	}
	if err := httpGetAndParseJSON(addr, &parsedResponse); err != nil {
		return emptyResponse, err
	}

	return parsedResponse.Versions, nil
}

func (ps *PaperServer) Run() error {
	binaryPath := ps.ServerBinaryPath
	runOpts := ps.Definition.Run
	return runJavaServer(binaryPath, runOpts)
}
