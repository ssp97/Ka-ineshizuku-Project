package conf

import (
	"github.com/FloatTech/ZeroBot-Plugin/app/gag"
	"github.com/FloatTech/ZeroBot-Plugin/app/setutime"
	"github.com/FloatTech/ZeroBot-Plugin/app/snare"
	"github.com/FloatTech/ZeroBot-Plugin/app/thunder"
)

type AppConfig struct {
	Snare 		snare.Config  //随机祸害
	Gag			gag.Config
	Setutime 	setutime.Config
	Thunder		thunder.Config
}

