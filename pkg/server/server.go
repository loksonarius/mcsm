package server

import (
	"fmt"
)

type Server interface {
	Install() error
	Configure() error
	Versions() ([]string, error)
	Run() error
}

var serverInitializers = make(map[InstallKind]func(def ServerDefinition) Server)

func GetServer(def ServerDefinition) (Server, error) {
	kind := def.Install.Kind
	init, ok := serverInitializers[kind]
	if !ok {
		return nil, fmt.Errorf("no registered server handler of kind %s", kind)
	}

	server := init(def)
	return server, nil
}
