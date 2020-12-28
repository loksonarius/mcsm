#!/usr/bin/env bash

go test -timeout=30s -parallel=4 -coverprofile=coverage.out \
  ./internal/... ./pkg/...
go tool cover -html coverage.out

echo 'Coverage reported'
