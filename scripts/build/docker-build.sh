#!/usr/bin/env bash
set -euo pipefail

source ./_common/get_arch.sh || exit 1
pwd
pushd ../../build/package/docker/$ARCH/redis-stream

if [ -d "./app" ]
then
rm -rf ./app
fi

mkdir app

cp -r ../../../../../cmd/redis_stream/out/linux/$ARCH/* ./app
cp -r ../../../../../cmd/redis_stream/config.json ./app

docker build -t nostromo-parker-redis-stream-$ARCH:0.0.1 .

rm -rf ./app





