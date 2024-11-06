#!/bin/bash

# 记录当前目录
CUR_DIR=$(pwd)

# 输入根目录
read -p "Enter the root directory: " ROOT_DIR
cd "$ROOT_DIR" || { echo "Directory not found"; exit 1; }

# 输入 proto 文件名
read -p "Enter the proto file name (with extension): " PROTO_FILE

# 执行 protoc 命令
protoc --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative "$PROTO_FILE"

echo "Done!"

# 返回原目录
cd "$CUR_DIR" || exit 1
