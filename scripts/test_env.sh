#!/usr/bin/env bash

docker run \
  -it --rm \
  -v $PWD/build/mcsm-linux-amd64:/usr/local/bin/mcsm \
  -v $PWD/integration:/tests \
  --entrypoint /bin/bash \
  loksonarius/mcsm-integration-image
