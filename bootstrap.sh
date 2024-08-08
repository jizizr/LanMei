#!/bin/sh

ldd_output=$(ldd --version 2>&1)
if echo "$ldd_output" | grep -iq 'musl'; then
    SYSTEM_LIBC="musl"
elif echo "$ldd_output" | grep -iq 'glibc'; then
    SYSTEM_LIBC="glibc"
else
    echo "Unknown C library. Exiting..."
    exit 1
fi
# 定义ROOT_OUTPUT_DIR
ROOT_OUTPUT_DIR="output/$SYSTEM_LIBC"

# 获取output目录下的所有子目录
DIRS=$(find "$ROOT_OUTPUT_DIR" -mindepth 1 -maxdepth 1 -type d)

# 创建或清空pids文件
PID_FILE="$ROOT_OUTPUT_DIR/pids"
> "$PID_FILE"
# 遍历每个子目录并启动bootstrap.sh
for dir in $DIRS; do
  # 检查子目录下是否有bootstrap.sh脚本
  if [ -f "$dir/bootstrap.sh" ]; then
    echo "Starting bootstrap.sh in directory: $dir"
    # 以后台方式运行bootstrap.sh，并确保获取正确的PID
    cd "$dir"
    sh bootstrap.sh &
    cd -
    echo $! >> "$PID_FILE"
  else
    echo "No bootstrap.sh in $dir, skipping..."
  fi
done

echo "All processes have been started. PIDs are stored in $PID_FILE."
