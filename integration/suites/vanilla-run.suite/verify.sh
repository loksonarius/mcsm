#!/usr/bin/env bash
# verify.sh

EXPECTED_DIRS="\
  logs
  world
"

for d in ${EXPECTED_DIRS}; do
  if [[ ! -d "${d}" ]]; then
    echo "expected '${d}' to be present" && exit 1
  fi
done

EXPECTED_FILES="\
  server.properties
  eula.txt
  minecraft_server.jar
  ops.json
  whitelist.json
  banned-ips.json
  banned-players.json
  logs/latest.log
"

for f in ${EXPECTED_FILES}; do
  if [[ ! -f "${f}" ]]; then
    echo "expected '${f}' to be present" && exit 2
  fi
done

if ! grep 'Default game type: CREATIVE' logs/latest.log; then
  echo 'Expected creative mode log entry' && exit 3
fi
