#!/usr/bin/env bash
# verify.sh

if [[ "$(cat server.config | jq -r .install.Kind)" != "paper" ]]; then
  echo "expected install kind to be vanilla" && exit 1
fi

if [[ "$(cat server.config | jq -r .run.MaxMemory)" != "2147483648" ]]; then
  echo "expected max memory to be 2147483648" && exit 1
fi

if [[ "$(cat server.config | jq -r '.configs."eula.txt".Accepted')" != "false" ]]; then
  echo "expected eula to be unaccepted" && exit 1
fi

if [[ "$(cat server.config | jq -r '.configs."server.properties".Gamemode')" != "creative" ]]; then
  echo "expected gamemode to be bar" && exit 1
fi

if [[ "$(cat server.config | jq -r '.configs."bukkit.yml".Settings.AllowEnd')" != "false" ]]; then
  echo "expected end dimension to be unallowed" && exit 1
fi

if [[ "$(cat server.config | jq -r '.configs."bukkit.yml".Settings.ShutdownMessage')" != "Go home!" ]]; then
  echo "expected 'Go home!' shutdown message" && exit 1
fi

if [[ "$(cat server.config | jq -r '.configs."bukkit.yml".SpawnLimits.Ambient')" != "2" ]]; then
  echo "expected ambient spawn limit of 2" && exit 1
fi

if [[ "$(cat server.config | jq -r '.configs."bukkit.yml".ChunkGC.PeriodInTicks')" != "1200" ]]; then
  echo "expected chunk gc period of 1200" && exit 1
fi
