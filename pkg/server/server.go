package server

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/loksonarius/mcsm/pkg/config"
)

type Server interface {
	Install() error
	InstallOpts() InstallOpts
	RuntimeOpts() RuntimeOpts
	ConfigFiles() []config.ConfigFile
	Versions() ([]string, error)
	Run() error
}

type serverInitFunc func(def ServerDefinition) Server

var serverInitializers = make(map[InstallKind]serverInitFunc)

func registerServer(kind InstallKind, f serverInitFunc) {
	serverInitializers[kind] = f
}

func Validate(s Server) []error {
	errs := make([]error, 0, 0)

	for _, cfg := range s.ConfigFiles() {
		if err := cfg.Validate(); err != nil {
			if verr, ok := err.(validator.ValidationErrors); ok {
				for _, verr := range verr {
					errs = append(errs, verr)
				}
			} else {
				errs = append(errs, err)
			}
		}
	}

	return errs
}

func Configure(s Server) error {
	for _, cfg := range s.ConfigFiles() {
		if err := cfg.Write(); err != nil {
			return err
		}
	}

	return nil
}

func GetConfig(s Server) config.ConfigDict {
	cfg := make(map[string]interface{})

	cfg["install"] = s.InstallOpts()
	cfg["run"] = s.RuntimeOpts()

	configFiles := make(map[string]config.ConfigFile)
	for _, cf := range s.ConfigFiles() {
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
