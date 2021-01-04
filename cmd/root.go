package cmd

import (
	"fmt"
	"os"
)

type Cmd struct {
	Name    string
	Summary string
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

	command := "help"
	args := os.Args

	if len(os.Args) > 1 {
		command = args[1]
		args = args[2:]
	}

	if f, ok := commands[command]; ok {
		if err := f.Exec(args...); err != nil {
			fmt.Printf("error: %s\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("%s is not a valid subcommand!\n", command)
		os.Exit(1)
	}
}
