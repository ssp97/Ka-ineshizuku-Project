package app

import (
	"github.com/ssp97/ZeroBot-Plugin/app/EEAsst"
	"github.com/ssp97/ZeroBot-Plugin/app/gag"
	"github.com/ssp97/ZeroBot-Plugin/app/manager"
	"github.com/ssp97/ZeroBot-Plugin/app/setutime"
	"github.com/ssp97/ZeroBot-Plugin/app/snare"
	"github.com/ssp97/ZeroBot-Plugin/app/study"
	"github.com/ssp97/ZeroBot-Plugin/app/thunder"
	"github.com/ssp97/ZeroBot-Plugin/conf"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
	"github.com/wdvxdr1123/ZeroBot/message"
	"os"
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
		ctx.SendChain(message.Text("pong"))
	})

}