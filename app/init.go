package app

import (
	"fmt"
	"github.com/ssp97/Ka-ineshizuku-Project/app/EEAsst"
	"github.com/ssp97/Ka-ineshizuku-Project/app/gag"
	_ "github.com/ssp97/Ka-ineshizuku-Project/app/gifApp"
	_ "github.com/ssp97/Ka-ineshizuku-Project/app/haveAFriend"
	"github.com/ssp97/Ka-ineshizuku-Project/app/manager"
	"github.com/ssp97/Ka-ineshizuku-Project/app/setutime"
	"github.com/ssp97/Ka-ineshizuku-Project/app/snare"
	"github.com/ssp97/Ka-ineshizuku-Project/app/study"
	"github.com/ssp97/Ka-ineshizuku-Project/app/thunder"
	"github.com/ssp97/Ka-ineshizuku-Project/conf"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/zero"
	wsServerDriver "github.com/ssp97/Ka-ineshizuku-Project/pkg/zero/driver"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
	"github.com/wdvxdr1123/ZeroBot/message"
	"math/rand"
	"net/http"
	"os"
	"time"
)


func Init(c *conf.Config){

	manager.Init(c.App.Manager)

	snare.Init(c.App.Snare)
	gag.Init(c.App.Gag)
	setutime.Init(c.App.Setutime)
	thunder.Init(c.App.Thunder)

	study.Init(c.App.Study, c.Zerobot.NickName[0])

	EEAsst.Init(c.App.EEAsst)

	zerobotConfig := &c.Zerobot

	var dri []ZeroBot.Driver
	dri = append(dri, wsServerDriver.NewWebSocketServer())
	for i, _ := range zerobotConfig.Url {
		dri = append(dri, &driver.WSClient{
			Url:	zerobotConfig.Url[i],
			AccessToken: zerobotConfig.Token[i],
		})
	}


	zero.RunDefault(ZeroBot.Config{
		NickName:      zerobotConfig.NickName,
		CommandPrefix: zerobotConfig.Prefix,

		SuperUsers: append(zerobotConfig.SuperUser, os.Args[1:]...),

		Driver: dri,
	})

	ZeroBot.OnCommand("ping").SetBlock(true).SetPriority(999).Handle(func(ctx *ZeroBot.Ctx) {

		var d = float64(rand.Intn(10000*100000))/100000
		t := time.Now()
		for i:= 1.0; i <= 114514.0; i++ {
			d += i + i/10.0
		}
		result := fmt.Sprintf("pong %f %v",d,time.Since(t))
		ctx.SendChain(message.Text(result))
	})

	//ZeroBot.OnRequest().Handle(func(ctx *ZeroBot.Ctx) {
	//	fmt.Println(ctx.Event.RequestType)
	//	if ctx.Event.RequestType=="friend"{
	//		flag := ctx.Event.Flag
	//		ctx.SetFriendAddRequest(flag, true, "")
	//		fmt.Println("处理加好友")
	//	}
	//	if ctx.Event.RequestType=="group"{
	//		flag := ctx.Event.Flag
	//		subType := ctx.Event.SubType
	//		ctx.SetGroupAddRequest(flag, subType, true, "")
	//	}
	//})

	go http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", zerobotConfig.ServerPort), nil)

}