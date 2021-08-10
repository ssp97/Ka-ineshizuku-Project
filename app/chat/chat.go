/*
Package chat
对话插件 example
*/
package chat

import (
	"math/rand"
	"strconv"
	"time"

	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/extension/rate"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var poke = rate.NewManager(time.Minute*5, 8) // 戳一戳

func init() { // 插件主体
	// 被喊名字
	ZeroBot.OnFullMatch("", ZeroBot.OnlyToMe, ZeroBot.OnlyGroup).SetBlock(true).FirstPriority().
		Handle(func(ctx *ZeroBot.Ctx) {
			var nickname = ZeroBot.BotConfig.NickName[0]
			time.Sleep(time.Second * 1)
			ctx.SendChain(message.Text(
				[]string{
					nickname + "在此，有何贵干~",
					"(っ●ω●)っ在~",
					"这里是" + nickname + "(っ●ω●)っ",
					nickname + "不在呢~",
				}[rand.Intn(4)],
			))
		})
	// 戳一戳
	ZeroBot.On("notice/notify/poke", ZeroBot.OnlyToMe).SetBlock(true).FirstPriority().
		Handle(func(ctx *ZeroBot.Ctx) {
			var nickname = ZeroBot.BotConfig.NickName[0]
			switch {
			case poke.Load(ctx.Event.UserID).AcquireN(3):
				// 5分钟共8块命令牌 一次消耗3块命令牌
				time.Sleep(time.Second * 1)
				ctx.SendChain(message.Text("请不要戳", nickname, " >_<"))
			case poke.Load(ctx.Event.UserID).Acquire():
				// 5分钟共8块命令牌 一次消耗1块命令牌
				time.Sleep(time.Second * 1)
				ctx.SendChain(message.Text("喂(#`O′) 戳", nickname, "干嘛！"))
			default:
				ctx.SetGroupBan(
					ctx.Event.GroupID,
					ctx.Event.UserID, // 要禁言的人的qq
					60,
				)
				time.Sleep(time.Second * 1)
				ctx.SendChain(message.Text("生气了！"))
				// 频繁触发，不回复
			}
			return
		})
	// 群空调
	var AirConditTemp = map[int64]int{}
	var AirConditSwitch = map[int64]bool{}
	ZeroBot.OnFullMatch("空调开").SetBlock(true).FirstPriority().
		Handle(func(ctx *ZeroBot.Ctx) {
			AirConditSwitch[ctx.Event.GroupID] = true
			ctx.SendChain(message.Text("❄️哔~"))
		})
	ZeroBot.OnFullMatch("空调关").SetBlock(true).FirstPriority().
		Handle(func(ctx *ZeroBot.Ctx) {
			AirConditSwitch[ctx.Event.GroupID] = false
			delete(AirConditTemp, ctx.Event.GroupID)
			ctx.SendChain(message.Text("💤哔~"))
		})
	ZeroBot.OnRegex(`设置温度(\d+)`).SetBlock(true).FirstPriority().
		Handle(func(ctx *ZeroBot.Ctx) {
			if _, exist := AirConditTemp[ctx.Event.GroupID]; !exist {
				AirConditTemp[ctx.Event.GroupID] = 26
			}
			if AirConditSwitch[ctx.Event.GroupID] {
				temp := ctx.State["regex_matched"].([]string)[1]
				AirConditTemp[ctx.Event.GroupID], _ = strconv.Atoi(temp)
				ctx.SendChain(message.Text(
					"❄️风速中", "\n",
					"群温度 ", AirConditTemp[ctx.Event.GroupID], "℃",
				))
			} else {
				ctx.SendChain(message.Text(
					"💤", "\n",
					"群温度 ", AirConditTemp[ctx.Event.GroupID], "℃",
				))
			}
		})
	ZeroBot.OnFullMatch(`群温度`).SetBlock(true).FirstPriority().
		Handle(func(ctx *ZeroBot.Ctx) {
			if _, exist := AirConditTemp[ctx.Event.GroupID]; !exist {
				AirConditTemp[ctx.Event.GroupID] = 26
			}
			if AirConditSwitch[ctx.Event.GroupID] {
				ctx.SendChain(message.Text(
					"❄️风速中", "\n",
					"群温度 ", AirConditTemp[ctx.Event.GroupID], "℃",
				))
			} else {
				ctx.SendChain(message.Text(
					"💤", "\n",
					"群温度 ", AirConditTemp[ctx.Event.GroupID], "℃",
				))
			}
		})
}
