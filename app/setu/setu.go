package setu

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/TypeUtils"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/dbManager"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/fsUtils"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/pixivel"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/zero"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/extension/rate"
	"github.com/wdvxdr1123/ZeroBot/message"
	"io/ioutil"
	"math/rand"
	"net/http"
	"path"
	"strings"
	"sync"
	"time"
)

const PIXIV_IMG_PROXY = "https://i.pixiv.re"

var db *dbManager.ORM
var limit = rate.NewManager(time.Minute*1, 2)

type Config struct {
	Enable bool
	Server string
}

func initSetuData(){

	data,err := ioutil.ReadFile(path.Join(fsUtils.Getwd(), "static", "sql", "setu.sql"))
	if err != nil{
		fmt.Println(err)
	}
	db.DB.Exec(string(data))

	data,err = ioutil.ReadFile(path.Join(fsUtils.Getwd(), "static", "sql", "setu_tag.sql"))
	if err != nil{
		fmt.Println(err)
	}
	db.DB.Exec(string(data))
}

func Init(c Config) {
	var count int64
	db = dbManager.GetDb(dbManager.DEFAULT_DB_NAME)
	db.DB.AutoMigrate(
		setu{},
	)
	db.DB.AutoMigrate(
		setuTag{},
	)
	db.DB.AutoMigrate(
		setuTagTranslated{},
	)
	db.DB.Model(&setu{}).Count(&count)
	if count < 10000{
		log.Warn("Init setu data")
		initSetuData()
		log.Warn("Init setu ok")
	}
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	http.Handle("/app/setu/", r)
	r.GET("/app/setu/api/random", func(ctx *gin.Context) {
		var data setu
		//db.DB.Model(&setu{}).Where("r18 = ?", "0").Order("RANDOM()").First(&data)
		//db.DB.Model(&setu{}).Raw("tablesample bernoulli(0.1) ").Where("r18 = ?", "0").Order("RANDOM()").First(&data)
		db.DB.Raw("select * from setus tablesample bernoulli(0.1) where r18 = 0  limit 1").First(&data)
		data.Url = strings.ReplaceAll(data.Url, `{count}`, fmt.Sprintf("%d",rand.Intn(data.P) ))
		fmt.Print(data.Url)
		ctx.Redirect(http.StatusFound, fmt.Sprintf("%s%s",PIXIV_IMG_PROXY, data.Url))
	})

	r.GET("/app/setu/api/randomR18", func(ctx *gin.Context) {
		var data setu
		//db.DB.Model(&setu{}).Where("r18 = ?", "1").Order("	RANDOM()").First(&data)
		//db.DB.Model(&setu{}).Raw("tablesample bernoulli(0.1) ").Where("r18 = ?", "1").Order("RANDOM()").First(&data)
		db.DB.Raw("select * from setus tablesample bernoulli(0.1) where r18 = 1 limit 1").First(&data)
		data.Url = strings.ReplaceAll(data.Url, `{count}`, fmt.Sprintf("%d",rand.Intn(data.P) ))
		fmt.Print(data.Url)
		ctx.Redirect(http.StatusFound, fmt.Sprintf("%s%s",PIXIV_IMG_PROXY, data.Url))
	})

	if c.Enable == false{
		return
	}

	//setuApi.Login()

	zero.Default().OnRegex(`^来点(.*)$`).SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		var tag = ctx.State["regex_matched"].([]string)[1]
		getAndSendPic(ctx,tag, 0)
	})

	zero.Default().OnRegex(`^来张(.*)$`).SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		var tag = ctx.State["regex_matched"].([]string)[1]
		getAndSendPic(ctx,tag, -1)
	})

	zero.Default().OnRegex(`^来发(.*)$`).SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		var tag = ctx.State["regex_matched"].([]string)[1]
		getAndSendPic(ctx,tag, 1)
	})
	
	zero.Default().OnRegex(`^随机色图`).SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		var data setu
		//db.DB.Model(&setu{}).Where("r18 = ?", "0").Order("RANDOM()").First(&data)
		db.DB.Raw("select * from setus tablesample bernoulli(0.1) where r18 = 0 limit 1").First(&data)
		//j := fmt.Sprintf(`{"app":"com.tencent.miniapp","desc":"","meta":{"Image":"%s"}}`, fmt.Sprintf("%s%s",PIXIV_IMG_PROXY, data.Url))
		data.Url = strings.ReplaceAll(data.Url, `{count}`, fmt.Sprintf("%d",rand.Intn(data.P) ))
		//ctx.SendChain(zero.Cardimage(fmt.Sprintf("%s%s",PIXIV_IMG_PROXY, data.Url)))
		SendPixivBigPic(ctx, data)
	})
	
	zero.Default().OnRegex(`^真随机色图`).SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		url := fmt.Sprintf("https://%s/app/setu/api/random", c.Server)
		ctx.SendChain(zero.Share(url,"随机图",url))
	})

	zero.Default().OnCommand("setu_info").SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		type SetuCount struct {
			SetuR18 int
			Count int
		}
		var (
			picCount int64
			r18Count int64
			pgCount  int64
			tagCount int64
		)

		db.DB.Model(&setu{}).Count(&picCount)
		db.DB.Model(&setu{}).Where("r18 = ?", 1).Count(&r18Count)
		db.DB.Model(&setu{}).Where("r18 = ?", -1).Count(&pgCount)
		db.DB.Model(&setuTag{}).Count(&tagCount)
		ctx.SendChain(message.Text(fmt.Sprintf("setu:\r\n  count:\t%d\r\n  isR18:\t%d\r\n  isPG: \t%d\r\n  tags: \t%d",picCount, r18Count, pgCount, tagCount)))
	})

	zero.Default().OnRegex("^!setu_mapping:(.*)->(.*)$", ZeroBot.SuperUserPermission).FirstPriority().SetBlock(true).Handle(func(ctx *ZeroBot.Ctx) {
		zh := ctx.State["regex_matched"].([]string)[1]
		src := ctx.State["regex_matched"].([]string)[2]
		result := db.DB.Model(&setuTagTranslated{}).Create(&setuTagTranslated{
			Zh: zh,
			Src: src,
		})
		if result.Error != nil{
			ctx.SendChain(message.Text("映射创建失败，因为"),message.Text(result.Error))
			return
		}
		ctx.SendChain(message.Text("创建映射成功"))
		return
	})
	
	zero.Default().OnRegex("!setu_user:(.*)$", ZeroBot.SuperUserPermission).FirstPriority().SetBlock(true).Handle(func(ctx *ZeroBot.Ctx) {
		id := TypeUtils.StrToInt(ctx.State["regex_matched"].([]string)[1])
		data, err := pixivel.GetUserInfo(id)
		if err!= nil{
			ctx.SendChain(message.Text(fmt.Sprintf("发生错误了，因为%s", err)))
			return
		}
		ctx.SendChain(message.Text(*data))
		//pixivel.GetUserAllIllust(id)
	})

	zero.Default().OnRegex("!setu_user_ill:(.*)$", ZeroBot.SuperUserPermission).FirstPriority().SetBlock(true).Handle(func(ctx *ZeroBot.Ctx) {
		id := TypeUtils.StrToInt(ctx.State["regex_matched"].([]string)[1])
		_, err := pixivel.GetUserInfo(id)
		if err!= nil{
			ctx.SendChain(message.Text(fmt.Sprintf("发生错误了，因为%s", err)))
			return
		}
		//ctx.SendChain(message.Text(*data))
		data := pixivel.GetUserAllIllust(id)
		ctx.SendChain(message.Text(data))
	})

	zero.Default().OnRegex("!setu_add_user_all:(.*)$", ZeroBot.SuperUserPermission).FirstPriority().SetBlock(true).Handle(func(ctx *ZeroBot.Ctx) {
		count := 0
		id := TypeUtils.StrToInt(ctx.State["regex_matched"].([]string)[1])
		_, err := pixivel.GetUserInfo(id)
		if err!= nil{
			ctx.SendChain(message.Text(fmt.Sprintf("发生错误了，因为%s", err)))
			return
		}
		//ctx.SendChain(message.Text(*data))
		data := pixivel.GetUserAllIllust(id)
		for i:= range *data{
			ill := (*data)[i]
			//Pid 			int			`gorm:"uniqueIndex"`
			//P   			int
			//Title 			string		`gorm:"index"`
			//UserId 			int
			//UserAccount 	string
			//UserName		string
			//Url 			string
			//R18				int			`gorm:"index"`
			//Width			int
			//Height			int
			//Tags  			string
			//TagsTranslated	string
			//Caption			string
			err = addSetu(setu{
				Pid: ill.Pid,
				P:ill.P,
				Title: ill.Title,
				UserId: ill.UserId,
				UserAccount: ill.UserAccount,
				UserName: ill.UserName,
				Url: ill.Url,
				R18: ill.R18,
				Width: ill.Width,
				Height: ill.Height,
				Tags: strings.Join(ill.Tags, ", "),
			}, ill.Tags)
			if err == nil{
				count++
			}
		}
		ctx.SendChain(message.Text(fmt.Sprintf("增加了%d张图", count)))
	})

}

func getAndSendPic(ctx *ZeroBot.Ctx, tag string, isR18 int){
	if !limit.Load(ctx.Event.UserID).Acquire() && !ZeroBot.AdminPermission(ctx){
		ctx.SendChain(message.Reply(ctx.Event.MessageID),
			message.Text("服务受限！接口流量限制！"))
		return
	}
	data, err := searchRandom(tag, isR18)
	if err == nil {
		SendPixivPic(ctx, data)
		return
	}
	pid, url := loliconSearch(tag, isR18)
	if pid > 0{
		SendPixivPic(ctx, setu{Pid:pid, Url: url})
	}
	ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text(fmt.Sprintf("未找到%s相关的图片", tag)))
	return
}

func getSomePic(ctx *ZeroBot.Ctx, tag string, isR18 int, maxCount int){
	if !limit.Load(ctx.Event.UserID).Acquire() && !ZeroBot.AdminPermission(ctx){
		ctx.SendChain(message.Reply(ctx.Event.MessageID),
			message.Text("服务受限！接口流量限制！"))
		return
	}

	//SendSomePic
}



func SendPixivPic(ctx *ZeroBot.Ctx, data setu){
	url := fmt.Sprintf("%s/%s",PIXIV_IMG_PROXY, data.Url)
	id := ctx.SendChain(message.Reply(ctx.Event.MessageID),zero.ImageUrlMessage(url))
	if id == 0 {
		ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text("图片发送失败了"))
	}else{
		// 一分钟后撤回
		time.Sleep(60 * time.Second)
		ctx.DeleteMessage(id)
	}
}

func SendSomePic(ctx *ZeroBot.Ctx, data []setu){
	msg := []message.MessageSegment{message.Reply(ctx.Event.MessageID)}
	var lock sync.Locker
	var wait sync.WaitGroup
	for _, s := range data {
		wait.Add(1)
		go func(s setu) {
			url := fmt.Sprintf("%s/%s",PIXIV_IMG_PROXY, s.Url)
			m := zero.ImageUrlMessage(url)
			lock.Lock()
			msg = append(msg, m)
			lock.Unlock()
			wait.Done()
		}(s)
	}
	wait.Wait()
	id := ctx.SendChain(msg...)
	if id == 0 {
		ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text("图片发送失败了"))
	}
}

func SendPixivBigPic(ctx *ZeroBot.Ctx, data setu){
	url := fmt.Sprintf("%s/%s",PIXIV_IMG_PROXY, data.Url)
	id := ctx.SendChain(zero.Cardimage(zero.ImageUrlMessage(url).Data["file"]))
	if id == 0 {
		ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text("图片发送失败了"))
	}
}