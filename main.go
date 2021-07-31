package main

import (
	"github.com/FloatTech/ZeroBot-Plugin/app"
	"github.com/FloatTech/ZeroBot-Plugin/conf"
	"net/http"
	_ "net/http/pprof"
	// 注：以下插件均可通过前面加 // 注释，注释后停用并不加载插件
	// 下列插件可与 wdvxdr1123/ZeroBot v1.1.2 以上配合单独使用
	// 词库类
	_ "github.com/FloatTech/ZeroBot-Plugin/app/atri" // ATRI词库
	_ "github.com/FloatTech/ZeroBot-Plugin/app/chat" // 基础词库

	// 实用类
	_ "github.com/FloatTech/ZeroBot-Plugin/app/github"  // 搜索GitHub仓库
	_ "github.com/FloatTech/ZeroBot-Plugin/app/manager" // 群管
	_ "github.com/FloatTech/ZeroBot-Plugin/app/runcode" // 在线运行代码

	// 娱乐类
	_ "github.com/FloatTech/ZeroBot-Plugin/app/ai_false" // 服务器监控
	_ "github.com/FloatTech/ZeroBot-Plugin/app/minecraft"
	_ "github.com/FloatTech/ZeroBot-Plugin/app/music"   // 点歌
	_ "github.com/FloatTech/ZeroBot-Plugin/app/shindan" // 测定

	// b站相关
	_ "github.com/FloatTech/ZeroBot-Plugin/app/bilibili" // 查询b站用户信息
	_ "github.com/FloatTech/ZeroBot-Plugin/app/diana"    // 嘉心糖发病

	// 二次元图片
	_ "github.com/FloatTech/ZeroBot-Plugin/app/image_finder" // 关键字搜图
	_ "github.com/FloatTech/ZeroBot-Plugin/app/lolicon"      // lolicon 随机图片
	_ "github.com/FloatTech/ZeroBot-Plugin/app/rand_image"   // 随机图片与点评
	_ "github.com/FloatTech/ZeroBot-Plugin/app/saucenao"     // 以图搜图
	_ "github.com/FloatTech/ZeroBot-Plugin/app/setutime"     // 来份涩图

	// 迫害等
	_ "github.com/FloatTech/ZeroBot-Plugin/app/gag"   // 禁言套餐
	_ "github.com/FloatTech/ZeroBot-Plugin/app/snare" // 随机陷害

	_ "github.com/FloatTech/ZeroBot-Plugin/app/jieba" // 分词

	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var content = []string{
	"* OneBot + ZeroBot + Golang ",
	"* Version 1.0.4 - 2021-07-14 14:09:58.581489207 +0800 CST",
	"* Copyright © 2020 - 2021  Kanri, DawnNights, Fumiama, Suika",
	"* Project: https://github.com/FloatTech/ZeroBot-app",
}

func init() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[zero][%time%][%lvl%]: %msg% \n",
	})
	log.SetLevel(log.DebugLevel)
	go http.ListenAndServe("0.0.0.0:6060", nil)
}

func main() {

	if err := conf.Init(); err != nil {
		log.Error(" conf init err: %v", err)
		panic(err)
	}

	plugin := app.New(conf.Conf)
	plugin.Init()

	select {}
}
