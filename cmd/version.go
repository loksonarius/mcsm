package cmd

var (
	Version string
	Commit  string
)

var versionUsage = `Usage:
	%s version

The version subcommand will print out version information about the running
binary. The version info includes the semantic version being used as well as the
commit used to build it.

Local development builds may have a non-semantic version string of $USER-dev.
`

var VersionCmd = Cmd{
	Name:    "version",
	Summary: "Print version info",
	Usage:   versionUsage,
	Exec: func(args ...string) error {
		Log.Println(Version)
		Log.Printf("commit: %s\n", Commit)
		return nil
	},
}

func init() {
	registerSubcommand(VersionCmd)
}
