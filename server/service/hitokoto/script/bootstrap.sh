#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/rpc"
exec "$CURDIR/bin/rpc"