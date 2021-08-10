package study

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"strings"
)

func replaceAnswerSpecialStr(ctx *zero.Ctx, str string)(result string){
	result = strings.ReplaceAll(str, "[name]", message.At(ctx.Event.UserID).String())
	result = strings.ReplaceAll(result, "[enter]", "\r\n")
	result = strings.ReplaceAll(result, "[bot]", BotName)
	return
}

func replaceAskSpecialStr(ctx *zero.Ctx, str string)(result string){
	result = strings.ReplaceAll(str, BotName, "[bot]")
	return
}