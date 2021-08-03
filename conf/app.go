package conf

import (
	"github.com/FloatTech/ZeroBot-Plugin/app/EEAsst"
	"github.com/FloatTech/ZeroBot-Plugin/app/gag"
	"github.com/FloatTech/ZeroBot-Plugin/app/manager"
	"github.com/FloatTech/ZeroBot-Plugin/app/setutime"
	"github.com/FloatTech/ZeroBot-Plugin/app/snare"
	"github.com/FloatTech/ZeroBot-Plugin/app/study"
	"github.com/FloatTech/ZeroBot-Plugin/app/thunder"
)

type AppConfig struct {
	Snare 		snare.Config  //随机祸害
	Gag			gag.Config
	Setutime 	setutime.Config
	Thunder		thunder.Config
	Manager 	manager.Config
	EEAsst		EEAsst.Config
	Study		study.Config
}

