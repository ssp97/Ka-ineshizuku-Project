go version
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GO111MODULE=auto
#go mod tidy
go build -ldflags="-s -w" -o Ka-ineshizuku-Project
./Ka-ineshizuku-Project
#go run main.go
