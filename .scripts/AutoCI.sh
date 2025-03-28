#!/bin/bash

# 切换到上一级目录
cd ..

# 初始化 Swagger 文档，指定入口文件为 ./cmd/main.go
swag init -g ./cmd/main.go

# 进入 cmd 目录
cd cmd/

# 构建 Go 程序，输出到 ../gin-server
go build -o ../gin-server main.go

# 检查 tmux 会话是否存在并尝试连接
if tmux has-session -t gin-server 2>/dev/null; then
    # 如果会话存在，尝试连接
    tmux attach -t gin-server
else
    # 如果会话不存在，创建一个新会话并运行 gin-server
    tmux new -s gin-server ../gin-server
fi

# 如果 tmux attach 退出（例如会话结束），直接运行 gin-server
# 注意：这一步可能不需要，取决于你的意图，这里保留你的原始逻辑
./gin-server