#!/usr/bin/env bash
# verify.sh

if [[ ! -f "bedrock_server" ]]; then
  echo "expected bedrock_server to be present" && exit 1
fi
