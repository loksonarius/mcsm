#!/usr/bin/env bash
# verify.sh

EXPECTED_DIRS="\
  logs
  plugins
  world
  world_nether
  world_the_end
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

if ! grep 'Reloading ResourceManager: Default, bukkit' logs/latest.log; then
  echo 'Expected bukkit ResourceManager reload log' && exit 3
fi

if ! grep 'Loading ClearLag' logs/latest.log; then
  echo 'Expected ClearLag plugin init log' && exit 4
fi

if ! grep 'Stopping the server' logs/latest.log; then
  echo 'Expected graceful shutdown log' && exit 5
fi
