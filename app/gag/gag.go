package gag

import (
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/zero"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
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
		zero.Default().OnRegex(v, ZeroBot.OnlyGroup).SetBlock(true).SetPriority(40).
			Handle(func(ctx *ZeroBot.Ctx){
				ctx.SetGroupBan(
					ctx.Event.GroupID,
					ctx.Event.UserID, // 要禁言的人的qq
					60,
				)
				ctx.SendChain(message.Text("小小要求，可以满足"),message.At(ctx.Event.UserID)) //
				return
			})
	}
	ignoreMe := []string{"理我"}
	for _, v := range ignoreMe {
		zero.Default().OnRegex(v, ZeroBot.OnlyGroup).SetBlock(true).SetPriority(40).
			Handle(func(ctx *ZeroBot.Ctx){
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
