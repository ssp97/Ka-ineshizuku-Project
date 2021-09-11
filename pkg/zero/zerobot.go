package zero

import (
	"encoding/base64"
	"fmt"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/avoidExamine"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/fsUtils"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"io/ioutil"
	"net/http"
)

var defaultZero = ZeroBot.New()


func RunDefault(op ZeroBot.Config) {
	ZeroBot.Run(op)
}

func Default() *ZeroBot.Engine {
	return defaultZero
}

func ImageUrlMessage(url string)(message.MessageSegment){
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return message.Text("图片下载失败")
	}
	d, err := ioutil.ReadAll(res.Body)
	if err != nil{
		panic(err)
	}
	d = avoidExamine.PicByte(d)
	bs := base64.StdEncoding.EncodeToString(d)
	return ImageBase64Message(&bs)
}

func ImageFileMessage(path string)(message.MessageSegment){
	d := fsUtils.ReadFile(path)
	d = avoidExamine.PicByte(d)
	base64str := base64.StdEncoding.EncodeToString(d)
	return ImageBase64Message(&base64str)
}

func ImageBase64Message(b *string)(message.MessageSegment){
	return message.Image(fmt.Sprintf("base64://%s", *b))
}

func IsBot(id int64)bool{
	_,ok := ZeroBot.APICallers.Load(id)
	return ok
}

func IsGroupManager(ctx *ZeroBot.Ctx)bool{
	data := ctx.GetGroupMemberInfo(ctx.Event.GroupID, ctx.Event.SelfID, false)
	return data.Get("role").String() == "admin"
}