#!/usr/bin/env bash
set -e

OPEN_BROWSER="${1:-no}"

go test \
  -timeout=30s \
  -parallel=4 \
  -coverprofile=coverage.out \
  ./pkg/...
go tool cover -func coverage.out

case "${OPEN_BROWSER}" in
  yes|YES|y)
    go tool cover -html coverage.out
    ;;
  *)
    go tool cover -html coverage.out -o coverage.html
    ;;
esac

echo 'Coverage reported'
