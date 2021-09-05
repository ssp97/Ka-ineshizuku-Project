package setu

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/dbManager"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/fsUtils"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/zero"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/extension/rate"
	"github.com/wdvxdr1123/ZeroBot/message"
	"io/ioutil"
	"path"
	"strings"
	"time"
)

const PIXIV_IMG_PROXY = "https://i.pixiv.cat"

var db *dbManager.ORM
var limit = rate.NewManager(time.Minute*1, 5)

type Config struct {
	Enable bool
}

type setu struct {
	Id				int
	Pid 			int
	P   			int
	Title 			string
	UserId 			int
	UserAccount 	string
	UserName		string
	Url 			string
	R18				int
	Width			int
	Height			int
	Tag 			string
}

func initSetuData(){


	data,err := ioutil.ReadFile(path.Join(fsUtils.Getwd(), "static", "sql", "setu.sql"))
	if err != nil{
		fmt.Println(err)
	}
	sqlArr:=strings.Split(string(data),";")
	for _,sql:=range sqlArr{
		if sql==""{
			continue
		}
		db.DB.Exec(sql)
	}

}

func Init(c Config) {
	if c.Enable == false{
		return
	}

	var count int64
	db = dbManager.GetDb(dbManager.DEFAULT_DB_NAME)
	db.DB.AutoMigrate(
		setu{},
	)
	db.DB.Model(&setu{}).Count(&count)
	if count < 10000{
		log.Warn("Init setu data")
		initSetuData()
		log.Warn("Init setu ok")
	}

	zero.Default().OnRegex(`^来点(.*)$`).SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		if !limit.Load(ctx.Event.UserID).Acquire() {
			ctx.SendChain(message.Text("服务受限！"))
			return
		}
		var tag = ctx.State["regex_matched"].([]string)[1]
		var data setu
		result := db.DB.Model(&setu{}).Where("title like ?", fmt.Sprintf("%%%s%%", tag)).Where("r18 = ?", "0").Order("RANDOM()").First(&data)
		if result.Error != nil{
			ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text(fmt.Sprintf("未找到%s相关的图片", tag)))
			return
		}
		url := fmt.Sprintf("%s%s",PIXIV_IMG_PROXY, data.Url)
		id := ctx.SendChain(message.Reply(ctx.Event.MessageID),zero.ImageUrlMessage(url))
		if id == 0 {
			ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text("图片发送失败了"))
		}
	})
}