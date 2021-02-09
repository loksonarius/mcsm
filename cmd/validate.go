package cmd

import (
	"fmt"

	"github.com/loksonarius/mcsm/pkg/server"
)

var validateUsage = `Usage:
	%s validate {server_definition}

The validate subcommand will validate the config parsed from the server
definition specified at the path given by the 'server_definition' argument. If
no argument is given, then 'server_definition' defaults to 'server.yaml'.

Validation will looks for any possible misconfigurations such as invalid
literals or out-of-range values. These validation rules go beyond checking for
appropriate type use and actually verify documented limits according to the docs
for each config file preset.

If any validation errors are found, they will be printed out and the process
will terminate with a non-zero exit code. Otherwise, a confirmation message will
be printed out and the process will terminate with an exit code of 0.
`

var validateCmd = Cmd{
	Name:    "validate",
	Summary: "Validate config file values for a server definition",
	Usage:   validateUsage,
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

		errs := server.Validate(srv)
		if len(errs) == 0 {
			fmt.Println("no errors found")
			return nil
		}

		for _, err := range errs {
			fmt.Println(err.Error())
		}
		return fmt.Errorf("found %d errors", len(errs))
	},
}

func init() {
	registerSubcommand(validateCmd)
}
