package cmd

import (
	"github.com/loksonarius/mcsm/pkg/server"
)

var runCmd = Cmd{
	Name:    "run",
	Summary: "run the current directory's Minecraft server",
	Exec: func(args ...string) error {
		def, err := server.DefinitionFromPath("./test.yaml")
		if err != nil {
			return err
		}

		srv, err := server.GetServer(def)
		if err != nil {
			return err
		}

		if err := srv.Configure(); err != nil {
			return err
		}

		return srv.Run()
	},
}

func init() {
	registerSubcommand(runCmd)
}
