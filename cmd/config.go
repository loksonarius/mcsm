package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/loksonarius/mcsm/pkg/server"
)

var configUsage = `Usage:
	%s config {server_definition}

The config subcommand prints out whatever config has been parsed from the server definition
specified at the path given by the 'server_definition' argument. If no argument
is given, then 'server_definition' defaults to 'server.yaml'.

This subcommand primarily exists to help with testing configuration parsing and
verifying config values with some other JSON-parsing tooling.
`

var configCmd = Cmd{
	Name:    "config",
	Summary: "Print the parsed server definition",
	Usage:   configUsage,
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

		out, err := json.MarshalIndent(srv.Config(), "", "\t")
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
