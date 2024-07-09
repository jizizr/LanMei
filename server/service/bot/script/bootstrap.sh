#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=bot
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}