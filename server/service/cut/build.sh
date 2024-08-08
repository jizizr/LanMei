#!/usr/bin/env bash
RUN_NAME="rpc"
mkdir -p output/bin output/conf output/data
cp script/* output/
cp -r conf/* output/conf
cp -r data/* output/data
chmod +x output/bootstrap.sh
go build -o output/bin/${RUN_NAME}
