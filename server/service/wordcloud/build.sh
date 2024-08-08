#!/usr/bin/env bash
RUN_NAME="rpc"
ldd_output=$(ldd --version 2>&1)
if echo "$ldd_output" | grep -iq 'musl'; then
    SYSTEM_LIBC="musl"
elif echo "$ldd_output" | grep -iq 'glibc'; then
    SYSTEM_LIBC="glibc"
else
    echo "Unknown C library. Exiting..."
    exit 1
fi
mkdir -p output/bin output/conf output/data
cp script/* output/
cp -r conf/* output/conf
cp -r data/* output/data
cp biz/lib/${SYSTEM_LIBC}/libwcloud.a biz/lib
chmod +x output/bootstrap.sh
go build -o output/bin/${RUN_NAME}
rm biz/lib/libwcloud.a
