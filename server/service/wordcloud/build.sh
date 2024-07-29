#!/usr/bin/env bash
RUN_NAME="rpc"
mkdir -p output/bin output/conf output/lib output/data
cp script/* output/
cp -r conf/* output/conf
cp -r data/* output/data
cp -r biz/lib/libwcloud.* output/lib
chmod +x output/bootstrap.sh
go build -ldflags="-r ./lib" -o output/bin/${RUN_NAME}
