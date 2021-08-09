go version
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GO111MODULE=auto
go mod tidy
go build -ldflags="-s -w" -o Ka-ineshizuku-Project.exe
Ka-ineshizuku-Project.exe
#go run main.go
pause
