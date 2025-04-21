cd ..
cd cmd/
go build -o ../gin-server main.go
tmux attach -t gin-server
^C
./gin-server
