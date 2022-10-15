package setuai

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/ssp97/Ka-ineshizuku-Project/app/publicModels"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/zero"
	setuaiapi "github.com/ssp97/Ka-ineshizuku-Project/pkg/setuai"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/extension/shell"
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
	
	zero.Default().OnCommand("setuai_cmd").SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		fset := flag.FlagSet{}
		var(
			tag string
			uc string
			seed string
			width string
			height string
			scale string
			steps string
		)
		var n uint32
		binary.Read(rand.Reader, binary.LittleEndian, &n)

		fset.StringVar(&tag, "tag", "", "")
		fset.StringVar(&uc, "utag", "lowres, bad anatomy, bad hands, text, error, missing fingers, extra digit, fewer digits, cropped, worst quality, low quality, normal quality, jpeg artifacts, signature, watermark, username, blurry, lowres, bad anatomy, bad hands, text,error, missing fngers,extra digt ,fewer digits,cropped, wort quality ,low quality,normal quality, jpeg artifacts,signature,watermark, username, blurry, bad feet", "")
		fset.StringVar(&seed, "seed", fmt.Sprintf("%d", n), "")
		fset.StringVar(&width, "width", "512", "")
		fset.StringVar(&height, "height", "768", "")
		fset.StringVar(&scale, "scale", "11", "")
		fset.StringVar(&steps, "steps", "20", "")

		arguments := shell.Parse(ctx.State["args"].(string))
		err := fset.Parse(arguments)
		if err != nil {
			ctx.SendChain(message.Reply(ctx.Event.MessageID),message.Text(fmt.Sprintf("参数解析失败啦：%s", err)))
			return
		}

		if tag == ""{
			ctx.SendChain(message.Reply(ctx.Event.MessageID),message.Text(fmt.Sprintf("怎么说也得先给点tag吧")))
			return
		}

		err, url := publicModels.GetSetting(SETU_AI_URL_KEY)
		if err!=nil{
			ctx.SendChain(message.Text("setuai:未设置url"))
		}
		ctx.SendChain(message.Text("少女祈祷中......"))
		img,txt := setuaiapi.Request(url, &tag, &width, &height, &scale, nil, &steps, &seed, &uc)
		imgB64 := "base64://" + base64.StdEncoding.EncodeToString(img)
		id := ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Image(imgB64), message.Text(txt))
		if id == 0{
			ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text("图片，好像被吃掉了呢"))
		}
	})
}
