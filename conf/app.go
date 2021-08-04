package conf

import (
	"github.com/ssp97/ZeroBot-Plugin/app/EEAsst"
	"github.com/ssp97/ZeroBot-Plugin/app/gag"
	"github.com/ssp97/ZeroBot-Plugin/app/manager"
	"github.com/ssp97/ZeroBot-Plugin/app/setutime"
	"github.com/ssp97/ZeroBot-Plugin/app/snare"
	"github.com/ssp97/ZeroBot-Plugin/app/study"
	"github.com/ssp97/ZeroBot-Plugin/app/thunder"
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

