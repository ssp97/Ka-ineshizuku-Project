package main

import (
	"fmt"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/zero"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"io"
	"net/http"
	"time"
)

/*
	正则 ((b23|acg)\.tv|bili2233.cn)\/[0-9a-zA-Z]+

*/


func getBilibiliInfo(){

}



func init() {
	zero.Default().OnRegex(`^>user info\s(.{1,25})$`).
		Handle(func(ctx *ZeroBot.Ctx) {

		})
}

// 搜索api：通过把触发指令传入的昵称找出uid返回
func bilibiliGetBv(u string) () {
	reader := *new(io.Reader)
	req, _ := http.NewRequest("GET",u,reader)
	c := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 30 * time.Second,
	}
	resp, err := c.Do(req)
	if err != nil{
		fmt.Println("err",err)
	}
	url := resp.Header.Get("location")
	fmt.Println(url)
}

func main() {
	bilibiliGetBv("https://b23.tv/EMX1nV")
}