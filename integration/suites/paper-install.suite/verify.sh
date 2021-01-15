#!/usr/bin/env bash
# verify.sh

if [[ ! -f "minecraft_server.jar" ]]; then
  echo "expected minecraft_server.jar to be present" && exit 1
fi

if [[ ! -d "plugins" ]]; then
  echo "expected plugins to be present" && exit 1
fi

EXPECTED_FILES="\
  plugins/Vault.jar
  plugins/Clearlag.jar
"

for f in ${EXPECTED_FILES}; do
  if [[ ! -f "${f}" ]]; then
    echo "expected '${f}' to be present" && exit 2
  fi
done
