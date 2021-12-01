go version
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GO111MODULE=auto
#go mod tidy
go build -ldflags="-s -w" -gcflags="-l -l -l -l"  -o Ka-ineshizuku-Project
# tinygo build -ldflags="-s -w" -gcflags="-l -l -l -l"  -o Ka-ineshizuku-Project
./Ka-ineshizuku-Project
