package cmd

import (
	"fmt"

	"github.com/loksonarius/mcsm/pkg/server"
)

var runUsage = `Usage:
	%s run {server_definition}

The run subcommand will execute the necessary binaries to run the server pointed
to by the 'server_definition' argument. If no argument is given, then
'server_definition' defaults to 'server.yaml'.

Right before actually starting the server process, all configuration files will
be generated and placed in their appropriate locations. This will replace any
currently existing instances of those files. These are files such as 'eula.txt',
'server.properties', 'bukkit.yml', and such.

The server will then be started attached to the current stdin, stdout, and
stderr used to call the run subcommand. This means servers will be started with
interactive sessions when possible (such as the Vanilla and Paper Java servers).
Meaning that commands can be sent to the server as if it had been run directly.

Graceful termination of the server will be handled by the run subcommand when
possible by catching SIGTERM, SIGKILL, and SIGINT signals and calling the
server's respective 'stop' commands. This will ensure worlds are saved before
the server process is killed. There is an internal timeout for this graceful
termination of 30 seconds that will force kill any server process that isn't
done shutting down.

NOTE: the bedrock server type doesn't have any available 'stop' commands to
handle, so kill and end signals will be forwarded right along to it!

The run subcommand relies the following expectations:
- the install subcommand has already been successfully run for the given server
- if the server requires the Java runtime, then the runtime should already be
  installed with the 'java' executable in the system's PATH
- if the server is of the bedrock server type, then the system should likely be
  an Ubuntu (18.04 or greater) server instance -- Mojang currently only supports
  running the Bedrock server's binary on Ubuntu systems
`

var RunCmd = Cmd{
	Name:    "run",
	Summary: "Start an installed Minecraft server",
	Usage:   runUsage,
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

		if err := server.Configure(srv); err != nil {
			return "", err
		}

		return "", srv.Run()
	},
}

func init() {
	registerSubcommand(RunCmd)
}
