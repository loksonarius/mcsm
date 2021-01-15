package server

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type JarSource struct {
	Local string
	URL   *url.URL
}

func (js *JarSource) UnmarshalYAML(value *yaml.Node) error {
	source := value.Value
	format := struct {
		File string
		URL  string
	}{}

	err := value.Decode(&format)
	if err != nil {
		return err
	}

	if p, err := filepath.Abs(format.File); err == nil && format.File != "" {
		js.Local = p
		return nil
	}

	if u, err := url.Parse(format.URL); err == nil && u.Host != "" {
		js.URL = u
		return nil
	}

	errMsg := fmt.Sprintf("could not parse %s as URL nor file path", source)
	return &yaml.TypeError{
		Errors: []string{errMsg},
	}
}

func (js *JarSource) storeToDirectory(dir string) error {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	var source io.Reader
	var filename string
	if js.Local != "" { // copy local file
		osfile, err := os.Open(js.Local)
		if err != nil {
			return err
		}
		defer osfile.Close()
		source = osfile
		filename = filepath.Base(js.Local)
	} else if js.URL != nil { // download remote file
		resp, err := http.Get(js.URL.String())
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		source = resp.Body
		filename = filepath.Base(js.URL.Path)
	} else {
		return fmt.Errorf("can't download from undefined source")
	}

	if err = os.MkdirAll(dir, os.ModeDir|0755); err != nil {
		return err
	}

	dest := filepath.Join(dir, filename)
	_, err = os.Stat(dest)
	if err == nil {
		if !os.IsNotExist(err) {
			if err = os.Remove(dest); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	target, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer target.Close()

	_, err = io.Copy(target, source)
	return err
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
