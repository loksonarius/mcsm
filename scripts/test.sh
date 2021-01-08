#!/usr/bin/env bash
set -e

go test -timeout=30s -parallel=4 ./pkg/...

echo 'Tested'
