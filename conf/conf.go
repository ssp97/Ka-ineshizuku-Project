package conf

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/FloatTech/ZeroBot-Plugin/pkg/db"
)


type Config struct {
	DB  db.OrmConfig
	Zerobot ZerobotConfig
	App AppConfig
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
