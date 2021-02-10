#!/usr/bin/env bash
set -euo pipefail

source ./_common/get_arch.sh || exit 1

NAME="tzota/nostromo-parker-redis-stream-$ARCH"
docker run --rm --name foo -it --env REDIS_SERVER=127.0.0.1 --network=host $NAME:0.0.1
