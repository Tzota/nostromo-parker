#!/usr/bin/env bash
set -euo pipefail

source ./_common/get_arch.sh || exit 1

docker images -a | grep -E "nostromo-parker-redis-stream-$ARCH" | awk '{ print $1":"$2 }' | xargs --no-run-if-empty docker rmi
