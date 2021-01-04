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
		def, err := server.DefinitionFromPath("./test.yaml")
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
