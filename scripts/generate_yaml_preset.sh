#!/usr/bin/env bash
set -eu

file="${1}"
component="${2:-type}"

if ! command -v ruby &> /dev/null; then
  echo "Need Ruby installed to run generate script!"
  exit 1
else
  ./scripts/generate_yaml_preset.rb "${file}" "${component}"
fi
