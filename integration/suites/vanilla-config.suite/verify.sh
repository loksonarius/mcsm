#!/usr/bin/env bash
# verify.sh

if [[ "$(cat server.config | jq -r .install.Kind)" != "vanilla" ]]; then
  echo "expected install kind to be vanilla" && exit 1
fi

if [[ "$(cat server.config | jq -r .run.MaxMemory)" != "1073741824" ]]; then
  echo "expected max memory to be 1073741824" && exit 1
fi

if [[ "$(cat server.config | jq -r '.configs."eula.txt".Accepted')" != "false" ]]; then
  echo "expected eula to be unaccepted" && exit 1
fi

if [[ "$(cat server.config | jq -r '.configs."server.properties".Gamemode')" != "survival" ]]; then
  echo "expected gamemode to be survival" && exit 1
fi

if [[ "$(cat server.config | jq -r '.configs."server.properties".PVP')" != "false" ]]; then
  echo "expected PVP to be false" && exit 1
fi

if [[ "$(cat server.config | jq -r '.configs."server.properties".RconPort')" != "42" ]]; then
  echo "expected RCON port to be 42" && exit 1
fi
