package setuai

import (
	"encoding/base64"
	"fmt"
	"github.com/ssp97/Ka-ineshizuku-Project/app/publicModels"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/zero"
	setuaiapi "github.com/ssp97/Ka-ineshizuku-Project/pkg/setuai"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

const SETU_AI_URL_KEY = "setu_ai_url"

func Init(){
	zero.Default().OnRegex("^!setuai_url:(.*)$", ZeroBot.SuperUserPermission).FirstPriority().SetBlock(true).Handle(func(ctx *ZeroBot.Ctx) {
		url := ctx.State["regex_matched"].([]string)[1]
		publicModels.SetSetting(SETU_AI_URL_KEY, url)

		ctx.SendChain(message.Text("setuai:url设置成功"))
	})

	zero.Default().OnRegex("^!setuai (.*)$").SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		tag := ctx.State["regex_matched"].([]string)[1]
		err, url := publicModels.GetSetting(SETU_AI_URL_KEY)
		if err!=nil{
			ctx.SendChain(message.Text("setuai:未设置url"))
		}
		print("tag = ", tag)
		ctx.SendChain(message.Text("少女祈祷中......"))
		img,txt := setuaiapi.Request(url, &tag, nil, nil, nil, nil, nil, nil, nil)
		imgB64 := "base64://" + base64.StdEncoding.EncodeToString(img)
		ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Image(imgB64), message.Text(txt))
	})

	zero.Default().OnRegex("^!setuai_b64 (.*)$").SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		tag := ctx.State["regex_matched"].([]string)[1]

		tagBytes,err := base64.StdEncoding.DecodeString(tag)
		if err!= nil{
			ctx.SendChain(message.Text(fmt.Sprintf("err:%s", err)))
		}
		tag = string(tagBytes)

		err, url := publicModels.GetSetting(SETU_AI_URL_KEY)
		if err!=nil{
			ctx.SendChain(message.Text("setuai:未设置url"))
		}
		print("tag = ", tag)
		ctx.SendChain(message.Text("少女祈祷中......"))
		img,txt := setuaiapi.Request(url, &tag, nil, nil, nil, nil, nil, nil, nil)
		imgB64 := "base64://" + base64.StdEncoding.EncodeToString(img)
		ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Image(imgB64), message.Text(txt))
	})
}
