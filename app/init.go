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
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
	"github.com/wdvxdr1123/ZeroBot/message"
	"math/rand"
	"os"
	"time"
)


func Init(c *conf.Config){

	manager.Init(c.App.Manager)

	snare.Init(c.App.Snare)
	gag.Init(c.App.Gag)
	setutime.Init(c.App.Setutime)
	thunder.Init(c.App.Thunder)

	study.Init(c.App.Study)

	EEAsst.Init(c.App.EEAsst)

	zerobotConfig := &c.Zerobot
	zero.Run(zero.Config{
		NickName:      zerobotConfig.NickName,
		CommandPrefix: zerobotConfig.Prefix,

		SuperUsers: append(zerobotConfig.SuperUser, os.Args[1:]...),

		Driver: []zero.Driver{
			&driver.WSClient{
				// OneBot 正向WS 默认使用 6700 端口
				Url:         zerobotConfig.Url,
				AccessToken: zerobotConfig.Token,
			},
		},
	})

	zero.OnCommand("ping").SetBlock(true).SetPriority(999).Handle(func(ctx *zero.Ctx) {

		var d = float64(rand.Intn(10000))
		t := time.Now()
		//log.Println("ping-pong")
		for i:= 1.0; i <= 114514.0; i++ {
			d += i + i/10.0
		}
		result := fmt.Sprintf("pong %f %v",d,time.Since(t))
		ctx.SendChain(message.Text(result))
	})

}