#!/usr/bin/env bash
set -euo pipefail

source ./_common/get_arch.sh || exit 1

cd ../../cmd/redis_stream
rm -rf ./out

CGO_ENABLED=0 GOARCH=$ARCH GOOS=linux go build -o out/linux/$ARCH/nostromo-parker-redis-stream .

echo "cmd/redis_stream for $ARCH has been built"
