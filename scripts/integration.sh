#!/usr/bin/env bash
set -e

suite="${1:-all}"

if [[ -z "${CI}" ]]; then
  docker run \
    -it --rm \
    -v $PWD/build/mcsm-linux-amd64:/usr/local/bin/mcsm \
    -v $PWD/integration:/tests \
    loksonarius/mcsm-integration-image run "${suite}"
else
  cp build/mcsm-linux-amd64 build/mcsm
  ./integration/suite.sh run "${suite}"
fi

echo "Integration suite run"
