package manager

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/dbManager"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/ssp97/Ka-ineshizuku-Project/pkg/zero"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"

	timer "github.com/FloatTech/ZeroBot-Plugin-Timer"
)

type Config struct {
	Enable bool
}

type Group struct {
	ID uint64		`json:"id" form:"id" gorm:"primary_key;"`
	Enable bool	`json:"enable" form:"enable"`
}

type BlackList struct {
	ID 		uint64
	Time 	uint64
	Forever bool
}

var db *dbManager.ORM

func GroupSwitchControl(ctx *ZeroBot.Ctx) bool{

	if ZeroBot.OnlyGroup(ctx) == false{
		return true
	}

	groupId := ctx.Event.GroupID
	var group Group
	result := db.DB.First(&group, groupId)
	if result.Error == gorm.ErrRecordNotFound {
		log.Debugln("------------------->创建记录")
		db.DB.Create(&Group{
			ID: uint64(groupId),
			Enable: false,
		})
		return false
	}
	return group.Enable
}

func Init(config Config) { // 插件主体
	db = dbManager.GetDb(dbManager.DEFAULT_DB_NAME)
	db.DB.AutoMigrate(Group{})

	//zero.UsePreHandler(GroupSwitchControl)

	ZeroBot.OnFullMatch("开启",ZeroBot.AdminPermission).SetBlock(true).FirstPriority().Handle(func(ctx *ZeroBot.Ctx) {
		db.DB.Table("groups").Where("id = ?", ctx.Event.GroupID).Update("enable",true)
		ctx.SendChain(message.Text("群开关已开启"))
	})

	ZeroBot.OnFullMatch("关闭",ZeroBot.AdminPermission).SetBlock(true).FirstPriority().Handle(func(ctx *ZeroBot.Ctx) {
		db.DB.Table("groups").Where("id = ?", ctx.Event.GroupID).Update("enable",false)
		ctx.SendChain(message.Text("群开关已关闭"))
	})

	zero.Default().OnFullMatch("群开关测试",GroupSwitchControl).SetBlock(true).FirstPriority().Handle(func(ctx *ZeroBot.Ctx) {
		ctx.SendChain(message.Text("已开启"))
	})

	zero.Default().UsePreHandler(GroupSwitchControl)


	if config.Enable == false{
		return
	}

	// 菜单
	zero.Default().OnFullMatch("群管系统", ZeroBot.AdminPermission).SetBlock(true).FirstPriority().
		Handle(func(ctx *ZeroBot.Ctx) {
			ctx.SendChain(message.Text(
				"====群管====", "\n",
				"- 禁言@QQ 1分钟", "\n",
				"- 解除禁言 @QQ", "\n",
				"- 我要自闭 1分钟", "\n",
				"- 开启全员禁言", "\n",
				"- 解除全员禁言", "\n",
				"- 升为管理@QQ", "\n",
				"- 取消管理@QQ", "\n",
				"- 修改名片@QQ XXX", "\n",
				"- 修改头衔@QQ XXX", "\n",
				"- 申请头衔 XXX", "\n",
				"- 踢出群聊@QQ", "\n",
				"- 退出群聊 1234", "\n",
				"- 群聊转发 1234 XXX", "\n",
				"- 私聊转发 0000 XXX",
			))
		})
	// 升为管理
	zero.Default().OnRegex(`^升为管理.*?(\d+)`, ZeroBot.OnlyGroup, ZeroBot.SuperUserPermission).SetBlock(true).SetPriority(40).
		Handle(func(ctx *ZeroBot.Ctx) {
			ctx.SetGroupAdmin(
				ctx.Event.GroupID,
				strToInt(ctx.State["regex_matched"].([]string)[1]), // 被升为管理的人的qq
				true,
			)
			nickname := ctx.GetGroupMemberInfo( // 被升为管理的人的昵称
				ctx.Event.GroupID,
				strToInt(ctx.State["regex_matched"].([]string)[1]), // 被升为管理的人的qq
				false,
			).Get("nickname").Str
			ctx.SendChain(message.Text(nickname + " 升为了管理~"))
		})
	// 取消管理
	zero.Default().OnRegex(`^取消管理.*?(\d+)`, ZeroBot.OnlyGroup, ZeroBot.SuperUserPermission).SetBlock(true).SetPriority(40).
		Handle(func(ctx *ZeroBot.Ctx) {
			ctx.SetGroupAdmin(
				ctx.Event.GroupID,
				strToInt(ctx.State["regex_matched"].([]string)[1]), // 被取消管理的人的qq
				false,
			)
			nickname := ctx.GetGroupMemberInfo( // 被取消管理的人的昵称
				ctx.Event.GroupID,
				strToInt(ctx.State["regex_matched"].([]string)[1]), // 被取消管理的人的qq
				false,
			).Get("nickname").Str
			ctx.SendChain(message.Text("残念~ " + nickname + " 暂时失去了管理员的资格"))
		})
	// 踢出群聊
	zero.Default().OnRegex(`^踢出群聊.*?(\d+)`, ZeroBot.OnlyGroup, ZeroBot.AdminPermission).SetBlock(true).SetPriority(40).
		Handle(func(ctx *ZeroBot.Ctx) {
			ctx.SetGroupKick(
				ctx.Event.GroupID,
				strToInt(ctx.State["regex_matched"].([]string)[1]), // 被踢出群聊的人的qq
				false,
			)
			nickname := ctx.GetGroupMemberInfo( // 被踢出群聊的人的昵称
				ctx.Event.GroupID,
				strToInt(ctx.State["regex_matched"].([]string)[1]), // 被踢出群聊的人的qq
				false,
			).Get("nickname").Str
			ctx.SendChain(message.Text("残念~ " + nickname + " 被放逐"))
		})
	// 退出群聊
	zero.Default().OnRegex(`^退出群聊.*?(\d+)`, ZeroBot.OnlyGroup, ZeroBot.AdminPermission).SetBlock(true).SetPriority(40).
		Handle(func(ctx *ZeroBot.Ctx) {
			ctx.SetGroupLeave(
				strToInt(ctx.State["regex_matched"].([]string)[1]), // 要退出的群的群号
				true,
			)
		})
	// 开启全体禁言
	zero.Default().OnRegex(`^开启全员禁言$`, ZeroBot.OnlyGroup, ZeroBot.AdminPermission).SetBlock(true).SetPriority(40).
		Handle(func(ctx *ZeroBot.Ctx) {
			ctx.SetGroupWholeBan(
				ctx.Event.GroupID,
				true,
			)
			ctx.SendChain(message.Text("全员自闭开始~"))
		})
	// 解除全员禁言
	zero.Default().OnRegex(`^解除全员禁言$`, ZeroBot.OnlyGroup, ZeroBot.AdminPermission).SetBlock(true).SetPriority(40).
		Handle(func(ctx *ZeroBot.Ctx) {
			ctx.SetGroupWholeBan(
				ctx.Event.GroupID,
				false,
			)
			ctx.SendChain(message.Text("全员自闭结束~"))
		})
	// 禁言
	zero.Default().OnRegex(`^禁言.*?(\d+).*?\s(\d+)(.*)`, ZeroBot.OnlyGroup, ZeroBot.AdminPermission).SetBlock(true).SetPriority(40).
		Handle(func(ctx *ZeroBot.Ctx) {
			duration := strToInt(ctx.State["regex_matched"].([]string)[2])
			switch ctx.State["regex_matched"].([]string)[3] {
			case "分钟":
				//
			case "小时":
				duration = duration * 60
			case "天":
				duration = duration * 60 * 24
			default:
				//
			}
			if duration >= 43200 {
				duration = 43199 // qq禁言最大时长为一个月
			}
			ctx.SetGroupBan(
				ctx.Event.GroupID,
				strToInt(ctx.State["regex_matched"].([]string)[1]), // 要禁言的人的qq
				duration*60,                                        // 要禁言的时间（分钟）
			)
			ctx.SendChain(message.Text("小黑屋收留成功~"))
		})
	// 解除禁言
	zero.Default().OnRegex(`^解除禁言.*?(\d+)`, ZeroBot.OnlyGroup, ZeroBot.AdminPermission).SetBlock(true).SetPriority(40).
		Handle(func(ctx *ZeroBot.Ctx) {
			ctx.SetGroupBan(
				ctx.Event.GroupID,
				strToInt(ctx.State["regex_matched"].([]string)[1]), // 要解除禁言的人的qq
				0,
			)
			ctx.SendChain(message.Text("小黑屋释放成功~"))
		})
	// 自闭禁言
	zero.Default().OnRegex(`^我要自闭.*?(\d+)(.*)`, ZeroBot.OnlyGroup).SetBlock(true).SetPriority(40).
		Handle(func(ctx *ZeroBot.Ctx) {
			duration := strToInt(ctx.State["regex_matched"].([]string)[1])
			switch ctx.State["regex_matched"].([]string)[2] {
			case "分钟":
				//
			case "小时":
				duration = duration * 60
			case "天":
				duration = duration * 60 * 24
			default:
				//
			}
			if duration >= 43200 {
				duration = 43199 // qq禁言最大时长为一个月
			}
			ctx.SetGroupBan(
				ctx.Event.GroupID,
				ctx.Event.UserID,
				duration*60, // 要自闭的时间（分钟）
			)
			ctx.SendChain(message.Text("那我就不手下留情了~"))
		})
	// 修改名片
	zero.Default().OnRegex(`^修改名片.*?(\d+).*?\s(.*)`, ZeroBot.OnlyGroup, ZeroBot.AdminPermission).SetBlock(true).SetPriority(40).
		Handle(func(ctx *ZeroBot.Ctx) {
			ctx.SetGroupCard(
				ctx.Event.GroupID,
				strToInt(ctx.State["regex_matched"].([]string)[1]), // 被修改群名片的人
				ctx.State["regex_matched"].([]string)[2],           // 修改成的群名片
			)
			ctx.SendChain(message.Text("嗯！已经修改了"))
		})
	// 修改头衔
	zero.Default().OnRegex(`^修改头衔.*?(\d+).*?\s(.*)`, ZeroBot.OnlyGroup, ZeroBot.AdminPermission).SetBlock(true).SetPriority(40).
		Handle(func(ctx *ZeroBot.Ctx) {
			ctx.SetGroupSpecialTitle(
				ctx.Event.GroupID,
				strToInt(ctx.State["regex_matched"].([]string)[1]), // 被修改群头衔的人
				ctx.State["regex_matched"].([]string)[2],           // 修改成的群头衔
			)
			ctx.SendChain(message.Text("嗯！已经修改了"))
		})
	// 申请头衔
	zero.Default().OnRegex(`^申请头衔(.*)`, ZeroBot.OnlyGroup).SetBlock(true).SetPriority(40).
		Handle(func(ctx *ZeroBot.Ctx) {
			ctx.SetGroupSpecialTitle(
				ctx.Event.GroupID,
				ctx.Event.UserID,                         // 被修改群头衔的人
				ctx.State["regex_matched"].([]string)[1], // 修改成的群头衔
			)
			ctx.SendChain(message.Text("嗯！不错的头衔呢~"))
		})
	// 群聊转发
	zero.Default().OnRegex(`^群聊转发.*?(\d+)\s(.*)`, ZeroBot.SuperUserPermission).SetBlock(true).SetPriority(40).
		Handle(func(ctx *ZeroBot.Ctx) {
			// 对CQ码进行反转义
			content := ctx.State["regex_matched"].([]string)[2]
			content = strings.ReplaceAll(content, "&#91;", "[")
			content = strings.ReplaceAll(content, "&#93;", "]")
			ctx.SendGroupMessage(
				strToInt(ctx.State["regex_matched"].([]string)[1]), // 需要发送的群
				content,                                            // 需要发送的信息
			)
			ctx.SendChain(message.Text("📧 --> " + ctx.State["regex_matched"].([]string)[1]))
		})
	// 私聊转发
	zero.Default().OnRegex(`^私聊转发.*?(\d+)\s(.*)`, ZeroBot.SuperUserPermission).SetBlock(true).SetPriority(40).
		Handle(func(ctx *ZeroBot.Ctx) {
			// 对CQ码进行反转义
			content := ctx.State["regex_matched"].([]string)[2]
			content = strings.ReplaceAll(content, "&#91;", "[")
			content = strings.ReplaceAll(content, "&#93;", "]")
			ctx.SendPrivateMessage(
				strToInt(ctx.State["regex_matched"].([]string)[1]), // 需要发送的人的qq
				content,                                            // 需要发送的信息
			)
			ctx.SendChain(message.Text("📧 --> " + ctx.State["regex_matched"].([]string)[1]))
		})

	// 定时提醒
	zero.Default().OnRegex(`^在(.{1,2})月(.{1,3}日|每?周.?)的(.{1,3})点(.{1,3})分时(用.+)?提醒大家(.*)`, ZeroBot.SuperUserPermission).SetBlock(true).SetPriority(40).
		Handle(func(ctx *ZeroBot.Ctx) {
			if ctx.Event.GroupID > 0 {
				dateStrs := ctx.State["regex_matched"].([]string)
				ts := timer.GetFilledTimeStamp(dateStrs, false)
				ts.Grpid = uint64(ctx.Event.GroupID)
				if ts.Enable {
					go timer.RegisterTimer(ts, true)
					ctx.Send("记住了~")
				} else {
					ctx.Send("参数非法!")
				}
			}
		})
	// 取消定时
	zero.Default().OnRegex(`^取消在(.{1,2})月(.{1,3}日|每?周.?)的(.{1,3})点(.{1,3})分的提醒`, ZeroBot.SuperUserPermission).SetBlock(true).SetPriority(40).
		Handle(func(ctx *ZeroBot.Ctx) {
			if ctx.Event.GroupID > 0 {
				dateStrs := ctx.State["regex_matched"].([]string)
				ts := timer.GetFilledTimeStamp(dateStrs, true)
				ts.Grpid = uint64(ctx.Event.GroupID)
				ti := timer.GetTimerInfo(ts)
				t, ok := (*timer.Timers)[ti]
				if ok {
					t.Enable = false
					delete(*timer.Timers, ti) //避免重复取消
					_ = timer.SaveTimers()
					ctx.Send("取消成功~")
				} else {
					ctx.Send("没有这个定时器哦~")
				}
			}
		})

	// 随机点名
	zero.Default().OnFullMatchGroup([]string{"翻牌"}).SetBlock(true).SetPriority(40).
		Handle(func(ctx *ZeroBot.Ctx) {
			if ctx.Event.GroupID > 0 {
				list := ctx.GetGroupMemberList(ctx.Event.GroupID)
				rand.Seed(time.Now().UnixNano())
				rand_index := fmt.Sprint(rand.Intn(int(list.Get("#").Int())))
				random_card := list.Get(rand_index + ".card").String()
				if random_card == "" {
					random_card = list.Get(rand_index + ".nickname").String()
				}
				ctx.Send(random_card + "，就是你啦!")
			}
		})
	// 入群欢迎
	zero.Default().OnNotice().SetBlock(false).SetPriority(40).
		Handle(func(ctx *ZeroBot.Ctx) {
			if ctx.Event.NoticeType == "group_increase" {
				ctx.SendChain(message.Text("欢迎~，具体用法请参考https://github.com/ssp97/Ka-ineshizuku-Project"))
			}
		})
	// 退群提醒
	zero.Default().OnNotice().SetBlock(false).SetPriority(40).
		Handle(func(ctx *ZeroBot.Ctx) {
			if ctx.Event.NoticeType == "group_decrease" {
				ctx.SendChain(message.Text("有人跑路了~"))
			}
		})
	// 运行 CQ 码
	zero.Default().OnRegex(`^run(.*)$`, ZeroBot.SuperUserPermission).SetBlock(true).SetPriority(0).
		Handle(func(ctx *ZeroBot.Ctx) {
			var cmd = ctx.State["regex_matched"].([]string)[1]
			cmd = strings.ReplaceAll(cmd, "&#91;", "[")
			cmd = strings.ReplaceAll(cmd, "&#93;", "]")
			ctx.Send(cmd)
		})
}

func strToInt(str string) int64 {
	val, _ := strconv.ParseInt(str, 10, 64)
	return val
}
