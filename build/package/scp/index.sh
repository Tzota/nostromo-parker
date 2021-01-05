#!/usr/bin/env bash
set -euo pipefail

if [ $# -eq 0 ]
then
echo "scp-compatible path needed"
exit 1
fi

cd ../../../

pushd scripts/
./release_terminal.sh
popd

scp cmd/terminal/out/linux/arm/nostromo-parker-terminal $1
