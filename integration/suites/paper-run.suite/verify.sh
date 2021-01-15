#!/usr/bin/env bash
# verify.sh

EXPECTED_DIRS="\
  logs
  plugins
  world
  world_nether
"

for d in ${EXPECTED_DIRS}; do
  if [[ ! -d "${d}" ]]; then
    echo "expected '${d}' to be present" && exit 1
  fi
done

if [[ -d "world_the_end" ]]; then
  echo "expected end dimension dir to not be present" && exit 2
fi

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
    echo "expected '${f}' to be present" && exit 3
  fi
done

if ! grep 'Reloading ResourceManager: Default, bukkit' logs/latest.log; then
  echo 'Expected bukkit ResourceManager reload log' && exit 4
fi

if ! grep 'Loading ClearLag' logs/latest.log; then
  echo 'Expected ClearLag plugin init log' && exit 5
fi

if ! grep 'Stopping the server' logs/latest.log; then
  echo 'Expected graceful shutdown log' && exit 6
fi

if ! grep 'ambient: 2' bukkit.yml; then
  echo 'Expected ambient spawn limit of 2 in bukkit config' && exit 7
fi

if ! grep 'allow-end: false' bukkit.yml; then
  echo 'Expected end dimension to be unallowed in bukkit config' && exit 8
fi
