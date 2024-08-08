#!/bin/sh

APP_PATH="server/service"
# 检查 libc 版本
ldd_output=$(ldd --version 2>&1)
if echo "$ldd_output" | grep -iq 'musl'; then
    SYSTEM_LIBC="musl"
elif echo "$ldd_output" | grep -iq 'glibc'; then
    SYSTEM_LIBC="glibc"
else
    echo "Unknown C library. Exiting..."
    exit 1
fi
uid=$(id -u)
ROOT_OUTPUT_DIR="output/$SYSTEM_LIBC"

# 定义处理每个子目录的函数
process_directory() {
  local dir="$1"
  echo "Building in directory: $dir"

  # 检查子目录下是否有 build.sh 脚本，如果有则执行
  if [ -f "$dir/build.sh" ]; then
    cd "$dir"
    sh build.sh
    if [ $uid -eq 0 ]; then
      chmod -R 777 "output"
    fi
    cd -

    # 定义输出目录名（子目录的名称）
    subdir_name=$(basename "$dir")
    target_dir="$ROOT_OUTPUT_DIR/$subdir_name"
    echo "Target directory: $target_dir"

    # 如果目标目录不存在，则创建它
    if [ ! -d "$target_dir" ]; then
      mkdir -p "$target_dir"
    fi

    # 将 server/service/output/* 复制到根目录下的 output/subdir_name/
    if [ -d "$dir/output" ]; then
      cp -r "$dir/output/"* "$target_dir/"
      echo "Copied output files to $target_dir"
    else
      echo "No output directory in $dir, skipping copy..."
    fi
    if [ $SYSTEM_LIBC = "musl" ]; then
      if [ ! -f "$dir/Dockerfile" ]; then
      # 创建 Dockerfile
      cat <<EOF > "$target_dir/Dockerfile"
FROM alpine:latest
WORKDIR /app
COPY . .
CMD ["sh","bootstrap.sh"]
EOF
      else
        cp "$dir/Dockerfile" "$target_dir/"
      fi
      echo "Dockerfile has been created successfully."
    fi
  else
    echo "No build.sh in $dir, skipping..."
  fi
}

# 如果提供了目录作为参数，则处理该目录
if [ $# -eq 1 ]; then
  process_directory "$APP_PATH/$1"
else
  # 获取APP_PATH下的一级子目录
  DIRS=$(find "$APP_PATH" -mindepth 1 -maxdepth 1 -type d)

  # 遍历每个子目录并处理
  for dir in $DIRS; do
    process_directory "$dir"
  done
fi
if [ $uid -eq 0 ]; then
  chmod -R 777 "output"
fi

#!/bin/sh

APP_PATH="server/service"
# 检查 libc 版本
ldd_output=$(ldd --version 2>&1)
uid=$(id -u)
if echo "$ldd_output" | grep -iq 'musl'; then
    SYSTEM_LIBC="musl"
elif echo "$ldd_output" | grep -iq 'glibc'; then
    SYSTEM_LIBC="glibc"
else
    echo "Unknown C library. Exiting..."
    exit 1
fi

ROOT_OUTPUT_DIR="output/$SYSTEM_LIBC"

# 定义处理每个子目录的函数
process_directory() {
  local dir="$1"
  echo "Building in directory: $dir"

  # 检查子目录下是否有 build.sh 脚本，如果有则执行
  if [ -f "$dir/build.sh" ]; then
    cd "$dir"
    sh build.sh
    cd -

    # 定义输出目录名（子目录的名称）
    subdir_name=$(basename "$dir")
    target_dir="$ROOT_OUTPUT_DIR/$subdir_name"
    echo "Target directory: $target_dir"

    # 如果目标目录不存在，则创建它
    if [ ! -d "$target_dir" ]; then
      mkdir -p "$target_dir"
    fi

    # 将 server/service/output/* 复制到根目录下的 output/subdir_name/
    if [ -d "$dir/output" ]; then
      cp -r "$dir/output/"* "$target_dir/"
      rm -rf "$dir/output"
      echo "Copied output files to $target_dir"
    else
      echo "No output directory in $dir, skipping copy..."
    fi
  else
    echo "No build.sh in $dir, skipping..."
  fi
}

# 如果提供了目录作为参数，则处理该目录
if [ $# -eq 1 ]; then
  process_directory "$APP_PATH/$1"
else
  # 获取APP_PATH下的一级子目录
  DIRS=$(find "$APP_PATH" -mindepth 1 -maxdepth 1 -type d)

  # 遍历每个子目录并处理
  for dir in $DIRS; do
    process_directory "$dir"
  done
fi
if [ $uid -eq 0 ]; then
  chmod -R 777 "output"
fi
