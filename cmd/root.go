package cmd

import (
	"fmt"
	"os"
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

	Log.Printf(USAGE_DOCS, cli, strings.Join(sub_docs, "\n"))
}

type Cmd struct {
	Name    string
	Summary string
	Usage   string
	Exec    func(args ...string) error
}

// These config values are initialized in setting.go
var (
	cli      string
	commands = make(map[string]Cmd)
)

func init() {
	cli = os.Args[0]
}

func registerSubcommand(c Cmd) {
	commands[c.Name] = c
}

func Execute(version, commit string) {
	Version = version
	Commit = commit

	command := ""
	args := os.Args

	if len(os.Args) > 1 {
		command = args[1]
		args = args[2:]
	}

	if command == "" {
		printUsageDocs()
		os.Exit(0)
	}

	if c, ok := commands[command]; ok {
		if err := c.Exec(args...); err != nil {
			Log.Fatalf("error: %s\n", err)
		}
	} else {
		Log.Fatalf("%s is not a valid subcommand!\n", command)
	}
}
