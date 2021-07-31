package app

import (
	"github.com/FloatTech/ZeroBot-Plugin/app/gag"
	"github.com/FloatTech/ZeroBot-Plugin/app/setutime"
	"github.com/FloatTech/ZeroBot-Plugin/app/snare"
	"github.com/FloatTech/ZeroBot-Plugin/app/thunder"
	"github.com/FloatTech/ZeroBot-Plugin/conf"
	"github.com/FloatTech/ZeroBot-Plugin/pkg/db"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
	"github.com/wdvxdr1123/ZeroBot/message"
	"os"
)

type App struct {
	c     	*conf.Config
	db   	*db.ORM
	//App  *AppConfig
	//minio *minio.Client
	//mqtt  *MQTTOrder
}

func New(c *conf.Config) (app *App) {
	app = &App{
		c:     c,
		db:   db.New(&c.DB),
	}
	return app
}

func (p *App)Init(){
	snare.Init(p.c.App.Snare)
	gag.Init(p.c.App.Gag)
	setutime.Init(p.c.App.Setutime)
	thunder.Init(p.c.App.Thunder)

	zerobotConfig := &p.c.Zerobot
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