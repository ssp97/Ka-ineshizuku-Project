package chatgpt

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/ssp97/Ka-ineshizuku-Project/app/publicModels"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/zero"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	chatgpt_go "github.com/zhan3333/chatgpt-go"
	log "github.com/sirupsen/logrus"
	"time"
)

const CHATGPT_TOKEN_CODE = "chatgpt_token"
var gC2 *chatgpt_go.ChatGPT
var gConversationDict = make(map[int64]*chatgpt_go.Conversation)
func Init(){

	//gC = chatgpt.NewChatGpt(chatgpt.NewClient(chatgpt.NewCredentials("0")))
	_, token  := publicModels.GetSetting(CHATGPT_TOKEN_CODE)
	if token != ""{
		var err error
		log.Info("init chatgpt")
		timeout := time.Second * 60
		gC2, err = chatgpt_go.NewChatGPT(token, chatgpt_go.ChatGPTOptions{
			Log:     logrus.NewEntry(logrus.StandardLogger()),
			Timeout: &timeout,
		})
		//gConversation = gC2.NewConversation("", "")
		if err != nil{
			log.Error(err)
		}
	}

	zero.Default().OnRegex("^!chatgpt_token:(.*)$", ZeroBot.SuperUserPermission).SetPriority(-100).SetBlock(true).Handle(func(ctx *ZeroBot.Ctx) {
		var err error
		token := ctx.State["regex_matched"].([]string)[1]
		publicModels.SetSetting(CHATGPT_TOKEN_CODE, token)

		//gC = chatgpt.NewChatGpt(chatgpt.NewClient(chatgpt.NewCredentials(token)))
		timeout := time.Second * 60
		gC2, err = chatgpt_go.NewChatGPT(token, chatgpt_go.ChatGPTOptions{
			Log:     logrus.NewEntry(logrus.StandardLogger()),
			Timeout: &timeout,
		})
		//gConversation = gC2.NewConversation("", "")

		if err != nil {
			ctx.SendChain(message.Text(fmt.Sprintf("设置失败，%s", err)))
		}else{
			ctx.SendChain(message.Text("设置成功"))
		}


	})

	zero.Default().OnRegex("^!chatgpt ((?:.|\n)*?)$").SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		text := ctx.State["regex_matched"].([]string)[1]

		uid := ctx.Event.GroupID
		if uid == 0{
			uid = ctx.Event.UserID
		}

		if _, ok := gConversationDict[uid]; !ok{
			gConversationDict[uid] = gC2.NewConversation("", "")
		}
		resp, err := gConversationDict[uid].SendMessage(text)

		if err != nil {
			// Handle err
			ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text(fmt.Sprintf("出错了, %s", err)))
			return
		}
		ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text(resp))

	})

	zero.Default().OnRegex("!chatgpt_clr").SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		uid := ctx.Event.GroupID
		if uid == 0{
			uid = ctx.Event.UserID
		}

		gConversationDict[uid] = gC2.NewConversation("", "")
		ctx.SendChain(message.Text("我好像...失忆了"))
	})
}
