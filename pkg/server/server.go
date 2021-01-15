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

type serverInitFunc func(def ServerDefinition) Server

var serverInitializers = make(map[InstallKind]serverInitFunc)

func registerServer(kind InstallKind, f serverInitFunc) {
	serverInitializers[kind] = f
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
