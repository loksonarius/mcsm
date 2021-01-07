package cmd

import (
	"fmt"

	"github.com/loksonarius/mcsm/pkg/server"
)

var runCmd = Cmd{
	Name:    "run",
	Summary: "run the current directory's Minecraft server",
	Exec: func(args ...string) error {
		if len(args) > 1 {
			return fmt.Errorf("expected only 1 argument")
		}

		path := "./server.yaml"
		if len(args) == 1 {
			path = args[0]
		}

		def, err := server.DefinitionFromPath(path)
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
