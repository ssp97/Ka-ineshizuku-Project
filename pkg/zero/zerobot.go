package zero

import (
	"encoding/base64"
	"fmt"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/avoidExamine"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/fsUtils"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var defaultZero = ZeroBot.New()


func RunDefault(op ZeroBot.Config) {
	ZeroBot.Run(op)
}

func Default() *ZeroBot.Engine {
	return defaultZero
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