package main

import (
	"github.com/ssp97/Ka-ineshizuku-Project/app"
	"github.com/ssp97/Ka-ineshizuku-Project/conf"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/dbManager"
	"io"
	_ "net/http/pprof"
	"os"
	"time"

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
	_ "github.com/ssp97/Ka-ineshizuku-Project/app/lolicon"      // lolicon 随机图片
	_ "github.com/ssp97/Ka-ineshizuku-Project/app/saucenao"     // 以图搜图

	_ "github.com/ssp97/Ka-ineshizuku-Project/app/jieba" // 分词

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	easy "github.com/ssp97/Ka-ineshizuku-Project/pkg/logFormatter"
)

func init() {
	logFile, err := rotatelogs.New("logs/%Y%m%d"+".log",
		rotatelogs.WithLinkName("logs/now.log"),
		rotatelogs.WithMaxAge(time.Duration(24 * 30)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour))
	if err != nil {
		log.Error("\t->Failed to log to file, using default stderr")
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[zero][%time%][%lvl%]: %msg% \n",
		LogMsgMaxLen: 255,
	})
	log.SetLevel(log.InfoLevel)
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
