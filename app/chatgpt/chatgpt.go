package chatgpt

import (
	"fmt"
	"context"
	"github.com/ssp97/Ka-ineshizuku-Project/app/publicModels"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/zero"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	openai "github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
)

const CHATGPT_TOKEN_CODE = "chatgpt_token"
const CHATGPT_DEFAULT_TEXT = "chatgpt_default_text"


var gC2 *openai.Client
var gChatgptCtxList = make(map[int64][]openai.ChatCompletionMessage)
func Init(){

	_, token  := publicModels.GetSetting(CHATGPT_TOKEN_CODE)
	gC2 = openai.NewClient(token)

	zero.Default().OnRegex("^!chatgpt_token:(.*)$", ZeroBot.SuperUserPermission).SetPriority(-100).SetBlock(true).Handle(func(ctx *ZeroBot.Ctx) {
		var err error
		token := ctx.State["regex_matched"].([]string)[1]
		publicModels.SetSetting(CHATGPT_TOKEN_CODE, token)

		gC2 = openai.NewClient(token)

		if err != nil {
			ctx.SendChain(message.Text(fmt.Sprintf("设置失败，%s", err)))
		}else{
			ctx.SendChain(message.Text("设置成功"))
		}
	})

	//zero.Default().OnRegex("^!chatgpt_timeout:(.*)$", ZeroBot.SuperUserPermission).SetPriority(-100).SetBlock(true).Handle(func(ctx *ZeroBot.Ctx) {
	//	string := ctx.State["regex_matched"].([]string)[1]
	//	second, err := strconv.Atoi(string)
	//	if err != nil{
	//		ctx.SendChain(message.Text("设置失败"))
	//		return
	//	}
	//	// publicModels.SetSetting(CHATGPT_TIMEOUT, string) // 不保存
	//	gC2.Timeout = time.Second * time.Duration(second)
	//
	//	ctx.SendChain(message.Text(fmt.Sprintf("设置成功，现在超时时间为%d秒", second)))
	//
	//})

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

		if gC2 == nil{
			ctx.SendChain(message.Text("未登录(各种原因"))
			return
		}

		if _, ok := gChatgptCtxList[uid]; !ok{
			gChatgptCtxList[uid] = []openai.ChatCompletionMessage{}
		}

		chatMsg := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		}

		resp, err := gC2.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: append(gChatgptCtxList[uid], chatMsg),
		})

		if err != nil {
			// Handle err
			log.Error(err)
			ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text(fmt.Sprintf("出错了, %s", err)))
			return
		}
		gChatgptCtxList[uid] = append(gChatgptCtxList[uid], chatMsg)
		gChatgptCtxList[uid] = append(gChatgptCtxList[uid], resp.Choices[0].Message)
		ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text(resp.Choices[0].Message.Content))

	})

	zero.Default().OnRegex("!chatgpt_clr").SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		uid := ctx.Event.GroupID
		if uid == 0{
			uid = ctx.Event.UserID
		}

		if gC2 == nil{
			ctx.SendChain(message.Text("未登录(各种原因"))
			return
		}

		gChatgptCtxList[uid] = []openai.ChatCompletionMessage{}

		err, defaultText  := publicModels.GetSetting(CHATGPT_DEFAULT_TEXT)

		if defaultText != "" && err != nil{
			chatMsg := openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleUser,
				Content: defaultText,
			}
			resp, err := gC2.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: append(gChatgptCtxList[uid], chatMsg),
			})
			if err != nil {
				// Handle err
				ctx.SendChain(message.Text("我好像...失忆了"))
				return
			}
			gChatgptCtxList[uid] = append(gChatgptCtxList[uid], chatMsg)
			gChatgptCtxList[uid] = append(gChatgptCtxList[uid], resp.Choices[0].Message)
			ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text(resp))
		}else{
			ctx.SendChain(message.Text("我好像...失忆了"))
		}

	})
}
