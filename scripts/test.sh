#!/usr/bin/env bash

go test -timeout=30s -parallel=4 ./pkg/...

echo 'Tested'
