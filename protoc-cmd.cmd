@echo off
REM 记录当前目录
set CUR_DIR=%cd%
set /p ROOT_DIR="Enter the root directory: "
cd %ROOT_DIR%
set /p PROTO_FILE="Enter the proto file name (with extension): "
protoc --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative %PROTO_FILE%
echo "Done!"
cd %CUR_DIR%