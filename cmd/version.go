package cmd

import "fmt"

var (
	Version string
	Commit  string
)

var versionCmd = Cmd{
	Name:    "version",
	Summary: "Print version info",
	Exec: func(args ...string) error {
		fmt.Println(Version)
		fmt.Printf("commit: %s\n", Commit)
		return nil
	},
}

func init() {
	registerSubcommand(versionCmd)
}
