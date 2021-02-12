package integration

import (
	"fmt"
	"testing"

	"github.com/loksonarius/mcsm/cmd"
)

func TestVanillaConfig(t *testing.T) {
	asTest(t, func() error {
		return withSuiteDir("vanilla-config", func() error {
			out, err := cmd.ConfigCmd.Exec("server.yaml")
			if err != nil {
				return fmt.Errorf("failed to get server config: %s", err)
			}

			tests := []jsonFieldCheck{
				{path: `install/Kind`, value: `"vanilla"`},
				{path: `run/DebugGC`, value: "false"},
				{path: `run/MaxMemory`, value: "1073741824"},
				{path: `configs/eula.txt/Accepted`, value: "false"},
				{path: `configs/eula.txt/Accepted`, value: "false"},
				{path: `configs/server.properties/Gamemode`, value: `"survival"`},
				{path: `configs/server.properties/PVP`, value: "false"},
				{path: `configs/server.properties/RconPort`, value: "42"},
			}

			if err := checkJsonFields(out, tests); err != nil {
				return fmt.Errorf("failed to validate server config: %s", err)
			}

			return nil
		})
	})
}

func TestVanillaRun(t *testing.T) {
	asTest(t, func() error {
		return withSuiteDir("vanilla-install", func() error {
			_, err := cmd.InstallCmd.Exec("server.yaml")
			if err != nil {
				return fmt.Errorf("failed to install server: %s", err)
			}

			tests := []fileStateCheck{
				{"minecraft_server.jar", isfile},
			}

			if err := checkFileStates(tests); err != nil {
				return fmt.Errorf("install missing files: %s", err)
			}

			return nil
		})
	})
}

func TestVanillaInstall(t *testing.T) {
	asTest(t, func() error {
		return withSuiteDir("vanilla-run", func() error {
			// no clue what to actually due here to simultaneously:
			// - run server looking for a "startup complete" log
			// - mantain a timeout for the server process
			// - send signals to the server process to cleanly terminate
			// - store server logs to check them
			_, err := cmd.InstallCmd.Exec("server.yaml")
			if err != nil {
				return fmt.Errorf("failed to install server: %s", err)
			}

			tests := []fileStateCheck{
				{"minecraft_server.jar", isfile},
			}

			if err := checkFileStates(tests); err != nil {
				return fmt.Errorf("install missing files: %s", err)
			}

			return nil
		})
	})
}
