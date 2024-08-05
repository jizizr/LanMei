#!/bin/bash

# 定义APP_PATH，使用已经设置的环境变量或者默认值
APP_PATH="server/service"
ROOT_OUTPUT_DIR="output"
# 获取APP_PATH下的一级子目录
DIRS=$(find "$APP_PATH" -mindepth 1 -maxdepth 1 -type d)

# 遍历每个子目录
for dir in $DIRS; do
  echo "Building in directory: $dir"

  # 检查子目录下是否有build.sh脚本，如果有则执行
  if [ -f "$dir/build.sh" ]; then
    pushd "$dir" > /dev/null
    bash build.sh
    popd > /dev/null

    # 定义输出目录名（子目录的名称）
    subdir_name=$(basename "$dir")
    # 目标目录是根目录下的 output/subdir_name
    target_dir="$ROOT_OUTPUT_DIR/$subdir_name"
    echo $target_dir
    # 如果目标目录不存在，则创建它
    if [ ! -d "$target_dir" ]; then
      mkdir -p "$target_dir"
    fi

    # 将 server/service/output/* 复制到根目录下的 output/subdir_name/
    if [ -d "$dir/output" ]; then
      cp -r -f "$dir/output/"* "$target_dir/"
      echo "Copied output files to $target_dir"
      strip -s $target_dir/bin/rpc
    else
      echo "No output directory in $dir, skipping copy..."
    fi
  else
    echo "No build.sh in $dir, skipping..."
  fi
done
