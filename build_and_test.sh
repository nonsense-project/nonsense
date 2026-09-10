#!/usr/bin/env bash
set -euo pipefail
go version
go mod verify
parallel=2
if [ -n "${NO_PARALLEL:-}" ]; then parallel=1; fi
go test -p "$parallel" -parallel "$parallel" -timeout 30m ./...
mkdir -p bin
go build -mod=readonly -trimpath -o bin/nonsensed .
go build -mod=readonly -trimpath -o bin/nonsensewallet ./cmd/nonsensewallet
