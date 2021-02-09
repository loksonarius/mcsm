package cmd

import (
	"fmt"

	"github.com/loksonarius/mcsm/pkg/server"
)

var installUsage = `Usage:
	%s install {server_definition}

The install subcommand will install server binaries, plugins, and mods as
specified in the server definition at the path given by the 'server_definition'
argument. If no argument is given, then 'server_definition' defaults to
'server.yaml'.

Part of the installation process includes checking for the existence of the
specified server version requested if it is specified in the server definition
file. If none is specified, or it is set to 'latest', then installation will try
to pull the latest available binaries for the server.

If there are plugins or mods specified, then they will also be download from
remote sources or copied over from local paths. This will only occur for server
types that support either mods or plugins (or both).

After installing a server, use the 'run' subcommand to actually kick things off.

The install subcommand can be pretty destructive in regards to overwritting
server files depending on the installation method. For server types such as
'vanilla' and 'paper', installation will override the 'minecraft_server.jar'
file if it already exists. And any plugins or mods with filenames that match
plugins or mods being added during installation will also be overriden. For
'bedrock' servers, the entire server is installed from a package distributed by
Mojang. The 'permissions.json' and 'whitelist.json' files will be purposely
ignored during installation if they already exist, but any other files
distributed in Mojang's package will be placed without regard for previous
existence.

As for remote plugin and mod sources, their sources should be specified as the
full URL that links directly to the .jar file being expected. There may be some
error during download if the URL given requires some extra forwarding to resolve
to the target file.
`

var installCmd = Cmd{
	Name:    "install",
	Summary: "Install a Minecraft server",
	Usage:   installUsage,
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

		return srv.Install()
	},
}

func init() {
	registerSubcommand(installCmd)
}
