#!/usr/bin/env bash
set -e

if [[ -n "${CI}" ]]; then
  exit 0
fi

docker build \
  --pull \
  -t loksonarius/mcsm-integration-image \
  -f integration/Dockerfile \
  .

echo 'Integration image built'
