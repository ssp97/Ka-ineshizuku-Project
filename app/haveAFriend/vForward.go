package haveAFriend

import (
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/zero"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"strings"
)

// 虚拟合并转发

func init() {
	zero.Default().OnRegex("^!伪造").SetPriority(40).SetBlock(true).Handle(func(ctx *ZeroBot.Ctx) {
		msg := message.Message{}
		userId := int64(0)
		user := ctx.GetGroupMemberInfo(ctx.Event.GroupID, ctx.Event.UserID, false)
		for _, segment := range ctx.Event.Message {
			if segment.Type == "at"{
				userId = strToInt(segment.Data["qq"])
				user = ctx.GetGroupMemberInfo(ctx.Event.GroupID, userId, false)
			}else{
				if userId == 0{
					continue
				}
				if segment.Type == "text" && strings.ReplaceAll(segment.Data["text"], " ", "") == ""{
					continue
				}

				msg = append(msg, message.CustomNode(user.Get("nickname").String(), userId, segment.String()) )

			}
		}
		ctx.SendGroupForwardMessage(ctx.Event.GroupID,msg)
	})

}