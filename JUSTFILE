export BINARY := "mcsm"
export VERSION := `test -z ${CI:-""} && echo "${USER:-'userless'}-dev" || git describe --always --abbrev=0 --dirty --tags`
export COMMIT := `git describe --always --abbrev=0 || echo "commit-less"`

alias c := check
alias cl := clean
alias b := build
alias t := test
alias co := coverage
alias i := integration

# Runs pre build checks to verify formatting, linting, and such
check:
  ./scripts/check.sh

# Cleans up build artifacts
clean:
  ./scripts/clean.sh

# Compile binary for current toolchain
build local="no":
  ./scripts/build.sh "{{local}}"

# Run all unit tests
test:
  ./scripts/test.sh

# Generates a test coverage report
coverage:
  ./scripts/coverage.sh

_build_integration_image: build
  ./scripts/build_integration_image.sh

# Generates a test coverage report
integration: _build_integration_image
  ./scripts/integration.sh
