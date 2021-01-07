#!/usr/bin/env bash

suite="${1:-all}"

docker run \
  -it --rm \
  -v $PWD/build/mcsm-linux-amd64:/usr/local/bin/mcsm \
  -v $PWD/integration:/tests \
  loksonarius/mcsm-integration-image run "${suite}"

echo "Integration suite run"
