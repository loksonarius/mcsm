package cmd

import (
	"fmt"
	"strings"
)

const USAGE_DOCS = `Minecraft Server Manager
Usage:
	%s {subcommand} {arguments}

Subcommands:
%s

Consider using 'help' to explore  the available subcommand's, and their specific
options.
`

func printUsageDocs() {
	sub_docs := make([]string, 0)
	for _, c := range commands {
		sub_docs = append(
			sub_docs,
			fmt.Sprintf("\t%s\t\t%s", c.Name, c.Summary),
		)
	}

	fmt.Printf(USAGE_DOCS, cli, strings.Join(sub_docs, "\n"))
}

var helpCmd = Cmd{
	Name:    "help",
	Summary: "Print this usage information",
	Exec: func(args ...string) error {
		printUsageDocs()
		return nil
	},
}

func init() {
	registerSubcommand(helpCmd)
}
