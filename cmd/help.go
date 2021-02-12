package cmd

import (
	"fmt"
)

var helpUsage = `Usage:
	%s help {subcommand}

The help subcommand will print out detailed usage information for the given
subcommand including what arguments the subcommand expects, if any.
`

var HelpCmd = Cmd{
	Name:    "help",
	Summary: "Print specifc usage information for a subcommand",
	Usage:   helpUsage,
	Exec: func(args ...string) (string, error) {
		if len(args) > 1 {
			return "", fmt.Errorf("expected only 1 argument")
		}

		subcommand := "help"
		if len(args) >= 1 {
			subcommand = args[0]
		}

		if c, ok := commands[subcommand]; ok {
			return fmt.Sprintf(c.Usage, cli), nil
		} else {
			return "", fmt.Errorf("%s is not a valid subcommand!\n", subcommand)
		}
	},
}

func init() {
	registerSubcommand(HelpCmd)
}
