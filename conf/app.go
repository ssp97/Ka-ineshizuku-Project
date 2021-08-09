package conf

import (
	"github.com/ssp97/Ka-ineshizuku-Project/app/EEAsst"
	"github.com/ssp97/Ka-ineshizuku-Project/app/gag"
	"github.com/ssp97/Ka-ineshizuku-Project/app/manager"
	"github.com/ssp97/Ka-ineshizuku-Project/app/setutime"
	"github.com/ssp97/Ka-ineshizuku-Project/app/snare"
	"github.com/ssp97/Ka-ineshizuku-Project/app/study"
	"github.com/ssp97/Ka-ineshizuku-Project/app/thunder"
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

