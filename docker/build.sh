#!/usr/bin/env sh
set -eu
repo_root=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
docker build -t nonsense-node:local -f "$repo_root/docker/Dockerfile" "$repo_root"
