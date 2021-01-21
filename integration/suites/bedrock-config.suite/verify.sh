#!/usr/bin/env bash
# verify.sh

if [[ "$(cat server.config | jq -r .install.Kind)" != "bedrock" ]]; then
  echo "expected install kind to be bedrock" && exit 1
fi

if [[ "$(cat server.config | jq -r '.configs."server.properties".ServerPort')" != "19132" ]]; then
  echo "expected port to be 19132" && exit 1
fi

if [[ "$(cat server.config | jq -r '.configs."server.properties".Gamemode')" != "creative" ]]; then
  echo "expected gamemode to be creative" && exit 1
fi

if [[ "$(cat server.config | jq -r '.configs."server.properties".PlayerMovementDistanceThreshold')" != "0.9" ]]; then
  echo "expected player movement distance threshold to be 0.9" && exit 1
fi
