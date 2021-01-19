package server

import (
	"fmt"

	"github.com/loksonarius/mcsm/pkg/config"
)

type Server interface {
	Install() error
	Config() config.ConfigDict
	Configure() error
	Versions() ([]string, error)
	Run() error
}

type serverInitFunc func(def ServerDefinition) Server

var serverInitializers = make(map[InstallKind]serverInitFunc)

func registerServer(kind InstallKind, f serverInitFunc) {
	serverInitializers[kind] = f
}

func condenseConfig(i InstallOpts, r RuntimeOpts, cfs []config.ConfigFile) config.ConfigDict {
	cfg := make(map[string]interface{})

	cfg["install"] = i
	cfg["run"] = r

	configFiles := make(map[string]config.ConfigFile)
	for _, cf := range cfs {
		configFiles[cf.Path()] = cf
	}
	cfg["configs"] = configFiles

	return cfg
}

func GetServer(def ServerDefinition) (Server, error) {
	kind := def.Install.Kind
	init, ok := serverInitializers[kind]
	if !ok {
		return nil, fmt.Errorf("no registered server handler of kind %s", kind)
	}

	server := init(def)
	return server, nil
}
