#!/usr/bin/env bash
# suite.sh
#
# Run either a specifc list of integration test suites, or all of them.
# Expects a suite to include a server definition as well as both a setup.sh and
# verify.sh script. The return code of the verify script will be used to
# determine if the integration test passed or failed.
# The return code of this script will be the count of failed integration suites
# and the output will include a list of failing suites at termination.

USAGE="\
Expected usage:

  ${0} {subcommand} {args}

Subcommands:
  list          List available integration test suites
  run SUITES    Run a given suite list
  clean SUITES  Clean up suite server directories
  new SUITE     Create a new suite
  help          Print these usage docs

When running a suite, the SUITES argument can be either the empty string, the
literal 'all', or a comma-separated list of suite names. The list of available
suites can be displayed by the list command. If SUITE is omitted or 'all', then
all available suites will be run.

Suites runs are maintained after a run regardless of success, so look in the
'servers' directory for resulting server installs and runs. These runs will be
cleaned up before setting up a new suite run, but they can also be manually
cleaned up using the 'clean' subcommand.

New suites can be generated from a basic template using the 'new' subcommand.
'new' takes a single argument SUITE that determines the name of the new suite
that will be generated. The new suite will located in the 'suites' directory.
"

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
SERVERS_DIR="${SCRIPT_DIR}/servers"
SUITES_DIR="${SCRIPT_DIR}/suites"
AVAILABLE_SUITES="$(find "${SUITES_DIR}" -name '*.suite' -type d -exec basename {} .suite \;)"
COMMAND="${1:-help}"

SUITES="${2:-all}"
if [ "${SUITES}" == "all" ]; then
  SUITES="${AVAILABLE_SUITES}"
fi
SUITES="${SUITES//,/ }"

case "${COMMAND}" in
  list)
    echo "${AVAILABLE_SUITES}"
    ;;
  run)
    FAILS=0
    for SUITE in ${SUITES}; do
      SUITE_DIR="${SUITES_DIR}/${SUITE}.suite"
      SERVER_DIR="${SERVERS_DIR}/${SUITE}"

      rm -rf "${SERVER_DIR}"
      mkdir -p "${SERVER_DIR}"
      cp -r "${SUITE_DIR}"/* "${SERVER_DIR}"

      pushd "${SERVER_DIR}" 2>&1 > /dev/null
        "${SERVER_DIR}/setup.sh"
        if [ $? -eq 0 ]; then
          "${SERVER_DIR}/verify.sh"
          if [ $? -ne 0 ]; then
            echo "Suite '${SUITE}' failed verify!"
            FAILS=$((FAILS+1))
          fi
        else
          echo "Suite '${SUITE}' failed setup!"
          FAILS=$((FAILS+1))
        fi
      popd 2>&1 > /dev/null
    done

    if [[ "${FAILS}" -eq 0 ]]; then
      echo "All suites passed"
    else
      echo "${FAILS} suite(s) failed"
    fi
    exit "${FAILS}"
    ;;
  clean)
    for SUITE in ${SUITES}; do
      SERVER_DIR="${SERVERS_DIR}/${SUITE}"
      rm -rf "${SERVER_DIR}"
    done
    ;;
  new)
    SUITE="${2}"
    SUITE_DIR="${SUITES_DIR}/${SUITE}.suite"
    if [[ -z "${SUITE}" ]]; then
      echo "Missing suite argument"
      exit 1
    fi

    if [[ -d "${SUITE_DIR}" ]]; then
      echo "Suite '${SUITE}' already exists"
      exit 1
    fi

    cp -r suites/.template "${SUITE_DIR}"
    ;;
  help)
    echo "${USAGE}"
    ;;
  *)
    echo "${USAGE}"
    exit 1
    ;;
esac
