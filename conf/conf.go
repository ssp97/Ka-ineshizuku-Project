package conf

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/ssp97/ZeroBot-Plugin/pkg/dbManager"
)


type Config struct {
	DB      dbManager.OrmConfig
	Zerobot ZerobotConfig
	App     AppConfig

}

var (
	confPath string
	// Conf config struct
	Conf = &Config{}
)

func init() {
	flag.StringVar(&confPath, "conf", "config.toml", "default config path")
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	fmt.Println(Conf)
	//fmt.Println(Conf.App)
	return
}

// Init init conf
func Init() error {
	if confPath != "" {
		return local()
	}
	return nil
}
