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
	"strconv"
	"time"
)

const CHATGPT_TOKEN_CODE = "chatgpt_token"
const CHATGPT_DEFAULT_TEXT = "chatgpt_default_text"
const CHATGPT_TIMEOUT = "chatgpt_timeout"
const CHATGPT_USER_AGENT = "chatgpt_user_agent"
const CHATGPT_CLEARANCE_TOKEN = "chatgpt_clearance_token"

var gC2 *chatgpt_go.ChatGPT
var gConversationDict = make(map[int64]*chatgpt_go.Conversation)
func Init(){

	//gC = chatgpt.NewChatGpt(chatgpt.NewClient(chatgpt.NewCredentials("0")))
	//_, token  := publicModels.GetSetting(CHATGPT_TOKEN_CODE)
	//if token != ""{
	//	var err error
	//	log.Info("init chatgpt")
	//	timeout := time.Second * 120
	//	gC2, err = chatgpt_go.NewChatGPT(token, chatgpt_go.ChatGPTOptions{
	//		Log:     logrus.NewEntry(logrus.StandardLogger()),
	//		Timeout: &timeout,
	//	})
	//	//gConversation = gC2.NewConversation("", "")
	//	if err != nil{
	//		log.Error(err)
	//	}
	//}

	go func() {
		for{
			if gC2 != nil{
				if err := gC2.RefreshAccessToken(); err != nil {
					log.Warn(fmt.Errorf("refresh access token: %w", err) )
				}
				time.Sleep(time.Minute * 1)
			}
		}

	}()

	zero.Default().OnRegex("^!chatgpt_user_agent:(.*)$", ZeroBot.SuperUserPermission).SetPriority(-100).SetBlock(true).Handle(func(ctx *ZeroBot.Ctx) {
		var err error
		token := ctx.State["regex_matched"].([]string)[1]
		publicModels.SetSetting(CHATGPT_USER_AGENT, token)

		if err != nil {
			ctx.SendChain(message.Text(fmt.Sprintf("设置失败，%s", err)))
		}else{
			ctx.SendChain(message.Text("设置成功"))
		}
	})

	zero.Default().OnRegex("^!chatgpt_clearance_token:(.*)$", ZeroBot.SuperUserPermission).SetPriority(-100).SetBlock(true).Handle(func(ctx *ZeroBot.Ctx) {
		var err error
		token := ctx.State["regex_matched"].([]string)[1]
		publicModels.SetSetting(CHATGPT_CLEARANCE_TOKEN, token)

		if err != nil {
			ctx.SendChain(message.Text(fmt.Sprintf("设置失败，%s", err)))
		}else{
			ctx.SendChain(message.Text("设置成功"))
		}
	})

	zero.Default().OnRegex("^!chatgpt_token:(.*)$", ZeroBot.SuperUserPermission).SetPriority(-100).SetBlock(true).Handle(func(ctx *ZeroBot.Ctx) {
		var err error
		token := ctx.State["regex_matched"].([]string)[1]
		publicModels.SetSetting(CHATGPT_TOKEN_CODE, token)
		_, cToken  := publicModels.GetSetting(CHATGPT_CLEARANCE_TOKEN)
		_, userAgent  := publicModels.GetSetting(CHATGPT_USER_AGENT)

		//gC = chatgpt.NewChatGpt(chatgpt.NewClient(chatgpt.NewCredentials(token)))
		timeout := time.Second * 120
		gC2, err = chatgpt_go.NewChatGPT(chatgpt_go.ChatGPTOptions{
			SessionToken: token,
			ClearanceToken: cToken,
			UserAgent: userAgent,
			Log:     logrus.NewEntry(logrus.StandardLogger()),
			Timeout: &timeout,
		})

		if err != nil {
			ctx.SendChain(message.Text(fmt.Sprintf("设置失败，%s", err)))
		}else{
			ctx.SendChain(message.Text("设置成功"))
		}
	})

	zero.Default().OnRegex("^!chatgpt_timeout:(.*)$", ZeroBot.SuperUserPermission).SetPriority(-100).SetBlock(true).Handle(func(ctx *ZeroBot.Ctx) {
		string := ctx.State["regex_matched"].([]string)[1]
		second, err := strconv.Atoi(string)
		if err != nil{
			ctx.SendChain(message.Text("设置失败"))
			return
		}
		// publicModels.SetSetting(CHATGPT_TIMEOUT, string) // 不保存
		gC2.Timeout = time.Second * time.Duration(second)

		ctx.SendChain(message.Text(fmt.Sprintf("设置成功，现在超时时间为%d秒", second)))

	})

	zero.Default().OnRegex("^!chatgpt_default_text:((?:.|\n)*?)$", ZeroBot.SuperUserPermission).SetPriority(-100).SetBlock(true).Handle(func(ctx *ZeroBot.Ctx) {
		token := ctx.State["regex_matched"].([]string)[1]
		publicModels.SetSetting(CHATGPT_DEFAULT_TEXT, token)


		ctx.SendChain(message.Text("设置成功"))

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

		err, defaultText  := publicModels.GetSetting(CHATGPT_DEFAULT_TEXT)
		if defaultText != "" && err != nil{
			resp, err := gConversationDict[uid].SendMessage(defaultText)
			if err != nil {
				// Handle err
				ctx.SendChain(message.Text("我好像...失忆了"))
				return
			}
			ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text(resp))
		}else{
			ctx.SendChain(message.Text("我好像...失忆了"))
		}


	})
}
