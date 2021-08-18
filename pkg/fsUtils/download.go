package fsUtils

import (
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"io"
	"net/http"
	"os"
)

func Download(ctx *ZeroBot.Ctx,file string, url string)(err error){
	f,err := os.Create(file)
	if err!= nil{
		return
	}
	defer f.Close()
	res, err := http.Get(url)
	if err!= nil{
		return
	}
	io.Copy(f, res.Body)
	return
}
