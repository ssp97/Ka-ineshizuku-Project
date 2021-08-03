package thunder

import (
	"fmt"
	"github.com/tidwall/gjson"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

type thunder struct {
	onlineList  []int64
	question	string
	answer		string
	victim		int64
	lastVictim  int64
}
var thunderList = map[int64]thunder{}

type Config struct {
	Enable bool
}

func Init(c Config) {

	if c.Enable == false {
		return
	}

	zero.OnRegex("^手捧雷$",zero.OnlyGroup).SetPriority(1).Handle(func(ctx *zero.Ctx) {
		group := ctx.Event.GroupID
		_,ok := thunderList[group]
		gameTime := 60 + rand.Intn(160)

		if ok==true{
			ctx.SendChain(message.Text("场上已经有雷了"))
			return
		}

		ctx.SendChain(message.Text(fmt.Sprintf("手捧雷游戏现在开始，游戏一共%d秒，回答正确，即可将雷传到其他人手中，准备好了吗？游戏即将开始,预备！",gameTime)))
		time.Sleep(5 * time.Second)

		q,a := questionMake()
		_onlineList := ctx.GetGroupMemberList(group)
		var onlineList []int64
		_onlineList.ForEach(func(key, value gjson.Result) bool {
			onlineList = append(onlineList, value.Get("user_id").Int())
			return true
		})

		t := thunder{
			question: q,
			answer: a,
			victim: ctx.Event.UserID,
			onlineList: onlineList,
		}
		thunderList[group] = t
		fmt.Println(t)

		ctx.SendChain(message.At(t.victim),message.Text(t.question))

		go func(ctx *zero.Ctx, group int64) {

			time.Sleep(1*time.Second)

			t := thunderList[group]
			startTime := time.Now().Unix()
			stopTime := startTime + int64(gameTime)
			for true {
				next := zero.NewFutureEvent("message", 1, false, zero.CheckUser(t.victim), func(ctx *zero.Ctx) bool {
					if ctx.Event.GroupID == group{
						return true
					}
					return false
				})

				recv, cancel := next.Repeat()
				WaitAnswer:
				for {
					select {
					case <- time.After(time.Second * time.Duration(stopTime - time.Now().Unix())):
						ctx.SendChain(message.Text("手捧雷BOOM，"),
							message.At(t.victim),
							message.Text(fmt.Sprintf("菊花残，满地伤，躺下%d秒捂菊花",gameTime)))
						ctx.SetGroupBan(
							group,
							t.victim, // 要禁言的人的qq
							int64(gameTime),
						)
						cancel()
						delete(thunderList, group)
						return
					case e := <-recv:
						//cancel()
						newCtx := &zero.Ctx{Event: e, State: zero.State{}}
						reg := regexp.MustCompile(t.answer)
						if reg.Match([]byte(newCtx.Event.Message.String())){
							ctx.SendChain(message.At(t.victim),
							message.Text("回答正确，来。你要把雷丢给谁？"))
							break WaitAnswer
						}else{
							ctx.SendChain(
								//message.At(t.victim),
								message.Text(fmt.Sprintf("回答错误，听清楚了，%s",t.question)))
						}
					}
				}
				WaitNextVictim:
				for  {
					select {
					case <- time.After(time.Second * time.Duration(stopTime - time.Now().Unix())):
						//ctx.SendChain(message.Text("奇怪的事情发生了，雷坏掉了"))
						ctx.SendChain(message.Text("啊偶，"),
							message.At(t.victim),
							message.Text(fmt.Sprintf("没有及时把雷传出去，手捧雷BOOM，菊花残，满地伤，躺下%d秒捂菊花",gameTime)))
						ctx.SetGroupBan(
							group,
							t.victim, // 要禁言的人的qq
							int64(gameTime),
						)
						cancel()
						delete(thunderList, group)
						return
					case e := <-recv:
						//
						newCtx := &zero.Ctx{Event: e, State: zero.State{}}
						reg := regexp.MustCompile("\\[CQ:at,qq=(\\d+)")
						result := reg.FindAllStringSubmatch(newCtx.Event.Message.String(),-1)
						if len(result) <= 0{
							ctx.SendChain(message.At(t.victim),
								message.Text("给谁给谁，我听不清"))
						} else {
							println(result[0][1])
							t.lastVictim = t.victim
							t.victim = strToInt(result[0][1])
							q,a := questionMake()
							t.question = q
							t.answer = a
							break WaitNextVictim
						}
					}
				}
				cancel()

				if t.victim == 1648468212{ // 小夜不会受伤
					ctx.SendChain(message.Text(fmt.Sprintf("问：%s 答：%s",t.question,t.answer)))
					time.Sleep(5*time.Second)
					ctx.SendChain(message.Text(fmt.Sprintf("问：%s 答:","回答正确，来。你要把雷丢给谁？")),
						message.At(t.lastVictim))
					time.Sleep(5*time.Second)
					stopTime += 12
				}

				ctx.SendChain(message.At(t.victim),message.Text(t.question))
			}
		}(ctx, group)

	})
}

func questionMake()(q , a string){
	rand.Seed(time.Now().Unix())
	x := rand.Intn(1000)
	y := rand.Intn(1000)
	z := x + y
	q = fmt.Sprintf("小学数学题： %d + %d = ?", x, y)
	a = fmt.Sprintf("%d",z)
	return
}

func strToInt(str string) int64 {
	val, _ := strconv.ParseInt(str, 10, 64)
	return val
}