package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/loksonarius/mcsm/pkg/server"
)

var configCmd = Cmd{
	Name:    "config",
	Summary: "Print the parsed server definition",
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

		out, err := json.MarshalIndent(def, "", "\t")
		if err != nil {
			return err
		}

		fmt.Println(string(out))
		return nil
	},
}

func init() {
	registerSubcommand(configCmd)
}
