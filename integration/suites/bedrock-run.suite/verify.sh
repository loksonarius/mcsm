#!/usr/bin/env bash
# verify.sh

EXPECTED_DIRS="\
  resource_packs
  minecraftpe
  worlds
"

for d in ${EXPECTED_DIRS}; do
  if [[ ! -d "${d}" ]]; then
    echo "expected '${d}' to be present" && exit 1
  fi
done

EXPECTED_FILES="\
  server.properties
  bedrock_server
  whitelist.json
  permissions.json
"

for f in ${EXPECTED_FILES}; do
  if [[ ! -f "${f}" ]]; then
    echo "expected '${f}' to be present" && exit 3
  fi
done

if ! grep 'Difficulty: 1 EASY' server.out; then
  echo 'Expected easy difficulty setting log' && exit 4
fi

if ! grep 'Game mode: 1 Creative' server.out; then
  echo 'Expected creative game mode setting log' && exit 5
fi

if ! grep 'Level Name: test-world' server.out; then
  echo 'Expected test-level level name setting log' && exit 6
fi
