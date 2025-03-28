cd ..
swag init -g ./cmd/main.go
cd cmd/
go build -o ../gin-server main.go
tmux attach -t gin-server
./gin-server
