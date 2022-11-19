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
const SETU_AI_URL_TYPE = "setu_ai_type"

func Init(){
	zero.Default().OnRegex("^!setuai_url:(.*)$", ZeroBot.SuperUserPermission).FirstPriority().SetBlock(true).Handle(func(ctx *ZeroBot.Ctx) {
		url := ctx.State["regex_matched"].([]string)[1]
		publicModels.SetSetting(SETU_AI_URL_KEY, url)

		ctx.SendChain(message.Text("setuai:url设置成功"))
	})

	zero.Default().OnRegex("^!setuai_type:(.*)$", ZeroBot.SuperUserPermission).FirstPriority().SetBlock(true).Handle(func(ctx *ZeroBot.Ctx) {
		t := ctx.State["regex_matched"].([]string)[1]

		if t == "sd" || t == "naifu"{
			publicModels.SetSetting(SETU_AI_URL_TYPE, t)
			ctx.SendChain(message.Text("setuai:type设置成功"))
		}else{
			ctx.SendChain(message.Text("setuai:type设置失败"))
		}


	})

	zero.Default().OnRegex("^!setuai (.*)$").SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		tag := ctx.State["regex_matched"].([]string)[1]
		err, url := publicModels.GetSetting(SETU_AI_URL_KEY)
		if err!=nil{
			ctx.SendChain(message.Text("setuai:未设置url"))
		}

		err, t := publicModels.GetSetting(SETU_AI_URL_TYPE)
		if err!=nil{
			ctx.SendChain(message.Text("setuai:未设置type"))
		}

		print("tag = ", tag)
		ctx.SendChain(message.Text("少女祈祷中......"))

		if t == "naifu" {
			img,txt := setuaiapi.NaifuRequest(url, &tag, nil, nil, nil, nil, nil, nil, nil)
			imgB64 := "base64://" + base64.StdEncoding.EncodeToString(img)
			ctx.SendChain(message.Reply(ctx.Event.MessageID), zero.ImageBase64Message(imgB64), message.Text(txt))
		}else if t == "sd"{
			img,txt := setuaiapi.SdRequest(url, &tag, nil, nil, nil, nil, nil, nil, nil)
			imgB64 := "base64://" + base64.StdEncoding.EncodeToString(img)
			ctx.SendChain(message.Reply(ctx.Event.MessageID), zero.ImageBase64Message(imgB64), message.Text(txt))
		}else{
			ctx.SendChain(message.Text("未设置参数"))
		}
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

		err, t := publicModels.GetSetting(SETU_AI_URL_TYPE)
		if err!=nil{
			ctx.SendChain(message.Text("setuai:未设置type"))
		}

		print("tag = ", tag)
		ctx.SendChain(message.Text("少女祈祷中......"))

		if t == "naifu" {
			img,txt := setuaiapi.NaifuRequest(url, &tag, nil, nil, nil, nil, nil, nil, nil)
			imgB64 := "base64://" + base64.StdEncoding.EncodeToString(img)
			ctx.SendChain(message.Reply(ctx.Event.MessageID), zero.ImageBase64Message(imgB64), message.Text(txt))
		}else if t == "sd"{
			img,txt := setuaiapi.SdRequest(url, &tag, nil, nil, nil, nil, nil, nil, nil)
			imgB64 := "base64://" + base64.StdEncoding.EncodeToString(img)
			ctx.SendChain(message.Reply(ctx.Event.MessageID), zero.ImageBase64Message(imgB64), message.Text(txt))
		}else{
			ctx.SendChain(message.Text("未设置参数"))
		}
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
			samper string
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
		fset.StringVar(&samper, "samper", "DDIM", "")

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
		err, t := publicModels.GetSetting(SETU_AI_URL_TYPE)
		if err!=nil{
			ctx.SendChain(message.Text("setuai:未设置type"))
		}


		ctx.SendChain(message.Text("少女祈祷中......"))

		if t == "naifu" {
			img,txt := setuaiapi.NaifuRequest(url, &tag, &width, &height, &scale, &samper, &steps, &seed, &uc)
			imgB64 := "base64://" + base64.StdEncoding.EncodeToString(img)
			id := ctx.SendChain(message.Reply(ctx.Event.MessageID), zero.ImageBase64Message(imgB64), message.Text(txt))
			if id == 0{
				ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text("图片，好像被吃掉了呢"))
			}
		}else if t == "sd"{
			img,txt := setuaiapi.SdRequest(url, &tag, &width, &height, &scale, &samper, &steps, &seed, &uc)
			imgB64 := "base64://" + base64.StdEncoding.EncodeToString(img)
			id := ctx.SendChain(message.Reply(ctx.Event.MessageID), zero.ImageBase64Message(imgB64), message.Text(txt))
			if id == 0{
				ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text("图片，好像被吃掉了呢"))
			}
		}else{
			ctx.SendChain(message.Text("未设置参数"))
		}
	})
}
