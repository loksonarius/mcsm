#!/usr/bin/env bash
# verify.sh

if [[ "$(cat server.config | jq -r .Install.Kind)" != "bedrock" ]]; then
  echo "expected install kind to be bedrock" && exit 1
fi

if [[ "$(cat server.config | jq -r .Configs.vanilla)" != "null" ]]; then
  echo "expected vanilla config to be null" && exit 1
fi

if [[ "$(cat server.config | jq -r .Configs.bedrock.gamemode)" != "creative" ]]; then
  echo "expected gamemode to be creative" && exit 1
fi

if [[ "$(cat server.config | jq -r '.Configs.bedrock."player-movement-distance-threshold"')" != "0.9" ]]; then
  echo "expected player movement distance threshold to be 0.9" && exit 1
fi
