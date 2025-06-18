#!/bin/bash

# 切换到上一级目录
cd ..

# 运行 swag init 生成 Swagger 文档
# swag init -g ./cmd/main.go

# 进入 cmd 目录
cd cmd/main/

# 编译 Go 程序，输出到上一级目录的 gin-server
go run main.go

cd ../saveData
go run main.go


# 检查 tmux 会话是否存在，如果不存在则创建
tmux has-session -t gin-server 2>/dev/null
if [ $? != 0 ]; then
    tmux new-session -d -s gin-server
fi

# 向 tmux 会话发送 Ctrl+C，终止当前运行的进程（如果有）
tmux send-keys -t gin-server C-c

# 向 tmux 会话发送命令，运行 ./gin-server
tmux send-keys -t gin-server "./gin-server" C-m

# 附着到 tmux 会话（可选，注释掉则在后台运行）
tmux attach -t gin-server