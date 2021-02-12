package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/loksonarius/mcsm/pkg/server"
)

var configUsage = `Usage:
	%s config {server_definition}

The config subcommand prints out whatever config has been parsed from the server
definition specified at the path given by the 'server_definition' argument. If
no argument is given, then 'server_definition' defaults to 'server.yaml'.

The output of this subcommand will include a JSON-formatted struct with the
install, runtime, and config files for a server as they will be used during
operations such as 'run' and 'install'. The struct includes not just what was
parsed directly from the server definition, but also any defaults not set for
any server files and options.

The rendered struct has the following form:

{
	"install": {},	// install time options including server type and version
	"run": {},		// run time options including memory restrictions
	"configs": {	// server files that will be rendered during configuration
		"path.extension": {},	// these include various files like
		...						// server.properties, eula.txt, or spigot.yml
	}
}

This subcommand primarily exists to help with testing configuration parsing and
verifying config values with some other JSON-parsing tooling.
`

var ConfigCmd = Cmd{
	Name:    "config",
	Summary: "Print a parsed server definition",
	Usage:   configUsage,
	Exec: func(args ...string) (string, error) {
		if len(args) > 1 {
			return "", fmt.Errorf("expected only 1 argument")
		}

		path := "./server.yaml"
		if len(args) == 1 {
			path = args[0]
		}

		def, err := server.DefinitionFromPath(path)
		if err != nil {
			return "", err
		}

		srv, err := server.GetServer(def)
		if err != nil {
			return "", err
		}

		out, err := json.MarshalIndent(server.GetConfig(srv), "", "\t")
		if err != nil {
			return "", err
		}

		return string(out), nil
	},
}

func init() {
	registerSubcommand(ConfigCmd)
}
