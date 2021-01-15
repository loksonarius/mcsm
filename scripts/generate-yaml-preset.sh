#!/usr/bin/env bash
set -eu

file="${1}"
component="${2:-type}"

if ! command -v ruby &> /dev/null; then
  echo "Need Ruby installed to run generate script!"
  exit 1
else
  ./scripts/generate-yaml-preset.rb "${file}" "${component}"
fi
