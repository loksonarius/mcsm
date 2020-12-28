package cmd

var runCmd = Cmd{
	Name:    "run",
	Summary: "Run the current directory's Minecraft server",
	Exec: func(args ...string) error {
		// parse server config path
		// ensure server is installed
		// build java opts
		return nil
	},
}

func init() {
	registerSubcommand(runCmd)
}
