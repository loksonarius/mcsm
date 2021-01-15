package server

import (
	"fmt"
	"regexp"

	"github.com/loksonarius/mcsm/pkg/config"
	"github.com/loksonarius/mcsm/pkg/config/presets"
)

const VanillaInstall = InstallKind("vanilla")

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
	registerServer(VanillaInstall, NewVanillaServer)
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

type vanillaVersionMetadata struct {
	ID          string
	Type        string
	URL         string
	Time        string
	ReleaseTime string
}

func (vs VanillaServer) getAvailableVersions() ([]vanillaVersionMetadata, error) {
	var versions []vanillaVersionMetadata
	addr := "https://launchermeta.mojang.com/mc/game/version_manifest.json"
	var parsedResponse struct {
		Latest struct {
			Release  string
			Snapshot string
		}
		Versions []vanillaVersionMetadata
	}
	if err := httpGetAndParseJSON(addr, &parsedResponse); err != nil {
		return versions, err
	}

	// we only care about SemVer releases
	semverFormat := regexp.MustCompile(`^\d+\.\d+(\.\d)?((-rc|-pre)\d+)?$`)
	for _, v := range parsedResponse.Versions {
		if semverFormat.MatchString(v.ID) {
			versions = append(versions, v)
		}
	}

	return versions, nil
}

func (vs VanillaServer) getVersionDownloadURL(version vanillaVersionMetadata) (string, error) {
	var parsedResponse struct {
		Downloads map[string]struct {
			Sha1 string
			URL  string
		}
	}
	if err := httpGetAndParseJSON(version.URL, &parsedResponse); err != nil {
		return "", err
	}

	if entry, ok := parsedResponse.Downloads["server"]; ok {
		return entry.URL, nil
	}

	return "", fmt.Errorf("no server download found for version %s", version.ID)
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
	binaryPath := vs.ServerBinaryPath
	runOpts := vs.Definition.Run
	return runJavaServer(binaryPath, runOpts)
}
