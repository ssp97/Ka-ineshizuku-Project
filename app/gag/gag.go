package gag

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"strconv"
)

type Config struct {
	Enable bool
}

func Init(c Config){

	if c.Enable == false{
		return
	}

	wantQuiet := []string{"想静静","要静静","自闭了","想自闭","要自闭"};

	for _, v := range wantQuiet {
		zero.OnRegex(v,zero.OnlyGroup).SetBlock(true).SetPriority(40).
			Handle(func(ctx *zero.Ctx){
				ctx.SetGroupBan(
					ctx.Event.GroupID,
					ctx.Event.UserID, // 要禁言的人的qq
					60,
				)
				ctx.SendChain(message.Text("小小要求，可以满足"),message.At(ctx.Event.UserID)) //
				return
			})
	}
	ignoreMe := []string{"理我"};
	for _, v := range ignoreMe {
		zero.OnRegex(v,zero.OnlyGroup).SetBlock(true).SetPriority(40).
			Handle(func(ctx *zero.Ctx){
				ctx.SetGroupBan(
					ctx.Event.GroupID,
					ctx.Event.UserID, // 要禁言的人的qq
					20,
				)
				ctx.SendChain(message.Text("当然可以"),message.At(ctx.Event.UserID)) //
				return
			})
	}

}

func strToInt(str string) int64 {
	val, _ := strconv.ParseInt(str, 10, 64)
	return val
}
