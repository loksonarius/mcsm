#!/usr/bin/env bash
set -e

docker build \
  --pull \
  -t loksonarius/mcsm-integration-image \
  -f integration/Dockerfile \
  .

echo 'Integration image built'
