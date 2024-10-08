#!/bin/bash

ldd_output=$(ldd --version 2>&1)
if echo "$ldd_output" | grep -iq 'musl'; then
    SYSTEM_LIBC="musl"
elif echo "$ldd_output" | grep -iq 'glibc'; then
    SYSTEM_LIBC="glibc"
else
    echo "Unknown C library. Exiting..."
    exit 1
fi
ROOT_OUTPUT_DIR="output/$SYSTEM_LIBC"

# 定义pids文件位置
PID_FILE="$ROOT_OUTPUT_DIR/pids"

# 检查pids文件是否存在
if [ ! -f "$PID_FILE" ]; then
  echo "PID file not found! No processes to terminate."
  exit 1
fi

# 读取pids文件并终止进程
while IFS= read -r pid; do
  if kill -0 "$pid" >/dev/null 2>&1; then
    echo "Terminating process $pid"
    kill "$pid"
  else
    echo "Process $pid not found or already terminated."
  fi
done < "$PID_FILE"

# 删除pids文件
rm -f "$PID_FILE"

echo "All processes have been terminated."
