#!/bin/bash

# 切换到上级目录
cd ..

# 执行 swag init 命令
swag init -g ./cmd/main.go

# 进入 cmd 目录
cd cmd/

# 编译 Go 文件并输出到上级目录的 gin-server
go build -o ../gin-server main.go

# 运行生成的可执行文件
./../gin-server