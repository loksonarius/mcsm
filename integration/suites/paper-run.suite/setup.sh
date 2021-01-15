#!/usr/bin/env bash
# setup.sh

set -e

mcsm install server.yaml

function runAndStopServer {
  SERVER_STARTUP_DONE_MESSAGE='For help, type "help"'
  mcsm run server.yaml 2>&1 > server.out &
  PID="${!}"
  while ! grep "${SERVER_STARTUP_DONE_MESSAGE}" server.out; do
    sleep 3
    echo "Waiting for server to finish startup"
  done

  echo "Server startup complete -- stopping now"
  # we wait /just/ a bit longer while plugins wrap up their post-init
  sleep 5 && kill "${PID}" && sleep 5
}
export -f runAndStopServer

timeout 90 bash -c runAndStopServer
