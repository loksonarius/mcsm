#!/usr/bin/env bash
# verify.sh

if [[ "$(cat server.config | jq -r .Install.Kind)" != "vanilla" ]]; then
  echo "expected install kind to be vanilla" && exit 1
fi

if [[ "$(cat server.config | jq -r .Run.MaxMemory)" != "1073741824" ]]; then
  echo "expected max memory to be 1073741824" && exit 1
fi

if [[ "$(cat server.config | jq -r .Configs.eula.accepted)" != "false" ]]; then
  echo "expected eula to be unaccepted" && exit 1
fi

if [[ "$(cat server.config | jq -r .Configs.bedrock)" != "null" ]]; then
  echo "expected bedrock config to be null" && exit 1
fi

if [[ "$(cat server.config | jq -r .Configs.vanilla.gamemode)" != "nap" ]]; then
  echo "expected gamemode to be nap" && exit 1
fi

if [[ "$(cat server.config | jq -r .Configs.vanilla.pvp)" != "false" ]]; then
  echo "expected PVP to be false" && exit 1
fi

if [[ "$(cat server.config | jq -r '.Configs.vanilla."rcon.port"')" != "42" ]]; then
  echo "expected RCON port to be 42" && exit 1
fi
