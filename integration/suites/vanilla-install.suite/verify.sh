#!/usr/bin/env bash
# verify.sh

if [[ ! -f "minecraft_server.jar" ]]; then
  echo "expected minecraft_server.jar to be present" && exit 1
fi
