# Integration Testing

Integration testing for `mcsm` is done by running suites in a controlled
environment that includes the necessary requirements to both run and verify
installations of Minecraft servers.

## Suites

Suites are basically just directories with some setup and verification scripts
that call some `mcsm` commands and then checks for expected conditions to be
true about the server installation, or command output, or any other user-facing
aspect of `mcsm`.

### Running through `just`

Runing suites is meant to be done using the `just` CLI. There is an
`integration` task defined in the [JUSTFILE](../JUSTFILE) that sets up the test
environment and runs suites within a container. This task is meant to cover the
really basic and frequent cycle of iterate, test, commit.

### Running inside the test environment

The test environment for integration tests is a Docker image built based off the
[Dockerfile](Dockerfile) that lives in this directory. To enter a bash shell
within the test environment, use the `test-env` task (see [JUSTFILE](JUSTFILE)
or `just --list` for details). This will start up a Docker container with the
`integration` directory mounted to `/tests` and the current `mcsm` Linux/AMD64
binary mounted in the enviornment's PATH (this means re-builds outside of the
enviornment will propagate to inside the environment).

Within the test environment, you can call the [suite.sh](suite.sh) script to
manage and run integration tests. `suite.sh` is set up as a rough CLI utility
that can run and initialize new integration test suites. `suite.sh` includes
much more detail on its operation in its usage docs -- to view these, simply
run:

```bash
./suite.sh help
```

### Directory layout

There is a [suite template](suites/.template) that defines some basic layout
suites should use for setup and execution.

#### server.yaml

This defines an `mcsm` server installation. The default server definition used
is a pretty bare vanilla server install. Feel free to reference the [server
definition spec](lol.dne) to see what you can control through here.

#### setup.sh

The setup script is meant to do things like run `mcsm install` or `mcsm run`.
This script is called before `verify.sh` and is expected to generally succeede
-- an error during setup won't stop `verify.sh` from running, but it'd be pretty
unexpected.

Some ideas for things to call during setup include:

- install a server using `mcsm install`
- print some config output to check later with `mcsm config > config.json`
- start a server in the background with `mcsm run; echo $! > server.pid`

#### verify.sh

Verification is meant to check for things like log file entries, files, config
options, command output formatting, and such. Really, just expect some list of
conditional statements followed by a brief error message and exit. Try
reading through the [`vanilla-run` suite's verify script for
reference](suites/vanilla-run.suite/verify.sh).
