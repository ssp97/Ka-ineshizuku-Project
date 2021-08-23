package main

import (
	"github.com/ssp97/Ka-ineshizuku-Project/app"
	"github.com/ssp97/Ka-ineshizuku-Project/conf"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/dbManager"
	_ "net/http/pprof"
	// 注：以下插件均可通过前面加 // 注释，注释后停用并不加载插件
	// 下列插件可与 wdvxdr1123/ZeroBot v1.1.2 以上配合单独使用
	// 词库类
	_ "github.com/ssp97/Ka-ineshizuku-Project/app/atri" // ATRI词库
	_ "github.com/ssp97/Ka-ineshizuku-Project/app/chat" // 基础词库

	// 实用类
	_ "github.com/ssp97/Ka-ineshizuku-Project/app/runcode" // 在线运行代码

	// 娱乐类
	_ "github.com/ssp97/Ka-ineshizuku-Project/app/ai_false" // 服务器监控

	// 二次元图片
	_ "github.com/ssp97/Ka-ineshizuku-Project/app/image_finder" // 关键字搜图
	_ "github.com/ssp97/Ka-ineshizuku-Project/app/lolicon"      // lolicon 随机图片
	_ "github.com/ssp97/Ka-ineshizuku-Project/app/saucenao"     // 以图搜图

	_ "github.com/ssp97/Ka-ineshizuku-Project/app/jieba" // 分词

	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var content = []string{
	"* ChatBot + ZeroBot + Golang ",
	"* Project: https://github.com/ssp97/ZeroBot-app",
}

func init() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[zero][%time%][%lvl%]: %msg% \n",
	})
	log.SetLevel(log.DebugLevel)
	//
}

func main() {

	if err := conf.Init(); err != nil {
		log.Error(" conf init err: %v", err)
		panic(err)
	}

	db := dbManager.New(&conf.Conf.DB)
	dbManager.SetDb(dbManager.DEFAULT_DB_NAME, db)
	app.Init(conf.Conf)

	//pgweb.Init()

	select {}
}
