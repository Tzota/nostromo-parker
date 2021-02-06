#!/usr/bin/env bash
set -euo pipefail

if [ $# -eq 0 ]
then
echo "scp-compatible path needed"
exit 1
fi

cd ../../../

pushd scripts/
./release_redis_stream.sh
popd

scp cmd/redis_stream/out/linux/arm/nostromo-parker-redis-stream $1
