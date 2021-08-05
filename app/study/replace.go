package study

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"strings"
)

func replaceSpecialStr(ctx *zero.Ctx, str string)(result string){
	result = strings.ReplaceAll(str, "[name]", message.At(ctx.Event.UserID).String())
	result = strings.ReplaceAll(result, "[enter]", "\r\n")
	return
}
