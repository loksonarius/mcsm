package server

type JarSource struct {
	Source string
}

func (js *JarSource) storeToDirectory(dest string) error {
	// ensure dest directory exists
	// check if source is file path or URL
	// if file path, ensure file exists w/ R perms, and try copy
	// if url, try HTTP GET and write file to dest
	return nil
}

type InstallCfg struct {
	Kind    InstallKind
	Version string
	Mods    []JarSource
	Plugins []JarSource
}

func newInstallCfg(kind InstallKind, version string, mods, plugins []JarSource) InstallCfg {
	return InstallCfg{
		Kind:    kind,
		Version: version,
		Mods:    mods,
		Plugins: plugins,
	}
}

func defaultInstallCfg() InstallCfg {
	return newInstallCfg(VanillaInstall, "latest", []JarSource{}, []JarSource{})
}

type ServerInstaller interface {
	Install(path string, cfgc InstallCfg) error
	Versions() ([]string, error)
}
