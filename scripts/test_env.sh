#!/usr/bin/env bash
set -e

if [[ -n "${CI}" ]]; then
  exit 0
fi

docker run \
  -it --rm \
  -v $PWD/build/mcsm-linux-amd64:/usr/local/bin/mcsm \
  -v $PWD/integration:/tests \
  --entrypoint /bin/bash \
  loksonarius/mcsm-integration-image
