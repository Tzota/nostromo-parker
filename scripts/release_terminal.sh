#!/usr/bin/env bash
set -euo pipefail

cd ../cmd/terminal
rm -rf ./out

GOARM=7 GOARCH=arm GOOS=linux go build -o out/linux/arm/nostromo-parker-terminal .

echo "cmd/terminal has been built"
