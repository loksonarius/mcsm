#!/usr/bin/env bash
# verify.sh

#!/usr/bin/env bash
# verify.sh

if [[ "$(cat server.config | jq -r .Install.Kind)" != "paper" ]]; then
  echo "expected install kind to be vanilla" && exit 1
fi

if [[ "$(cat server.config | jq -r .Run.MaxMemory)" != "2147483648" ]]; then
  echo "expected max memory to be 2147483648" && exit 1
fi

if [[ "$(cat server.config | jq -r .Configs.eula.accepted)" != "false" ]]; then
  echo "expected eula to be unaccepted" && exit 1
fi

if [[ "$(cat server.config | jq -r .Configs.bedrock)" != "null" ]]; then
  echo "expected bedrock config to be null" && exit 1
fi

if [[ "$(cat server.config | jq -r .Configs.vanilla.gamemode)" != "bar" ]]; then
  echo "expected gamemode to be bar" && exit 1
fi

if [[ "$(cat server.config | jq -r '.Configs.bukkit.settings."allow-end"')" != "false" ]]; then
  echo "expected end dimension to be unallowed" && exit 1
fi

if [[ "$(cat server.config | jq -r '.Configs.bukkit.settings."shutdown-message"')" != "Go home!" ]]; then
  echo "expected 'Go home!' shutdown message" && exit 1
fi

if [[ "$(cat server.config | jq -r '.Configs.bukkit."spawn-limits".ambient')" != "2" ]]; then
  echo "expected ambient spawn limit of 2" && exit 1
fi

if [[ "$(cat server.config | jq -r '.Configs.bukkit."chunk-gc"."period-in-ticks"')" != "1200" ]]; then
  echo "expected chunk gc period of 1200" && exit 1
fi
