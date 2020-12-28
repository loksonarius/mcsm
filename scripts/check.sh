#!/usr/bin/env bash
set -e

fmtlist="$(gofmt -l .)"
if [[ -n "$fmtlist" ]]; then
  echo 'Following files need formatting:'
  echo "$fmtlist"
  exit 1
fi

go mod verify
go vet

echo 'Checked'
