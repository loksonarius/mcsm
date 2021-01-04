package cmd

import (
	"github.com/loksonarius/mcsm/pkg/server"
)

var installCmd = Cmd{
	Name:    "install",
	Summary: "Install the current directory's Minecraft server",
	Exec: func(args ...string) error {
		def, err := server.DefinitionFromPath("./test.yaml")
		if err != nil {
			return err
		}

		srv, err := server.GetServer(def)
		if err != nil {
			return err
		}

		return srv.Install()
	},
}

func init() {
	registerSubcommand(installCmd)
}
