#!/usr/bin/env bash
set -euo pipefail

source ./_common/get_arch.sh || exit 1
pushd ../../build/package/docker/$ARCH/redis-stream

if [ -d "./app" ]
then
rm -rf ./app
fi

mkdir app

cp -r ../../../../../cmd/redis_stream/out/linux/$ARCH/* ./app
cp -r ../../../../../cmd/redis_stream/config.json ./app
SEMVER="0.0.3"
docker build -t nostromo-parker-redis-stream-$ARCH:$SEMVER .

docker image tag nostromo-parker-redis-stream-$ARCH:$SEMVER tzota/nostromo-parker-redis-stream-$ARCH:$SEMVER
docker image tag nostromo-parker-redis-stream-$ARCH:$SEMVER tzota/nostromo-parker-redis-stream-$ARCH:latest

docker push --all-tags tzota/nostromo-parker-redis-stream-$ARCH

rm -rf ./app





