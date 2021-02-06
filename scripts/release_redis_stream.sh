#!/usr/bin/env bash
set -euo pipefail

cd ../cmd/redis_stream
rm -rf ./out

GOARM=7 GOARCH=arm GOOS=linux go build -o out/linux/arm/nostromo-parker-redis-stream .

echo "cmd/redis_stream has been built"
