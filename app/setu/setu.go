package setu

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/dbManager"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/fsUtils"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/gocc"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/zero"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/extension/rate"
	"github.com/wdvxdr1123/ZeroBot/message"
	"io/ioutil"
	"math/rand"
	"net/http"
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
	Tags  			string
	TagsTranslated	string
	Caption			string
}

type setuTag struct {
	Pid			  	int 		`gorm:"index"`
	Tag 			string		`gorm:"index"`
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
		setuTag{},
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
	if c.Enable == false{
		return
	}

	zero.Default().OnRegex(`^来点(.*)$`).SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		if !limit.Load(ctx.Event.UserID).Acquire() {
			ctx.SendChain(message.Text("服务受限！"))
			return
		}
		var tag = ctx.State["regex_matched"].([]string)[1]
		var tagS2t = gocc.S2t(tag)
		var data setu
		result := db.DB.Model(&setu{}).Where("title like ?", fmt.Sprintf("%%%s%%", tag)).Where("r18 = ?", "0").Order("RANDOM()").First(&data)
		if result.Error == nil{
			SendPixivPic(ctx, data)
			return
		}
		result = db.DB.Model(&setu{}).Where("title like ?", fmt.Sprintf("%%%s%%", tagS2t)).Where("r18 = ?", "0").Order("RANDOM()").First(&data)
		if result.Error == nil{
			SendPixivPic(ctx, data)
			return
		}
		ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text(fmt.Sprintf("未找到%s相关的图片", tag)))
	})

	zero.Default().OnRegex(`^来张(.*)$`).SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {


	})
	
	zero.Default().OnRegex(`^随机色图`).SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		var data setu
		db.DB.Model(&setu{}).Where("r18 = ?", "0").Order("RANDOM()").First(&data)
		//j := fmt.Sprintf(`{"app":"com.tencent.miniapp","desc":"","meta":{"Image":"%s"}}`, fmt.Sprintf("%s%s",PIXIV_IMG_PROXY, data.Url))
		data.Url = strings.ReplaceAll(data.Url, `{count}`, fmt.Sprintf("%d",rand.Intn(data.P) ))
		ctx.SendChain(zero.Cardimage(fmt.Sprintf("%s%s",PIXIV_IMG_PROXY, data.Url)))
	})
	
	zero.Default().OnRegex(`^真随机色图`).SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		data := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
				 <msg serviceID="1">
				 <item><title>客官，这是你要的色图</title></item>
				 <source name="setu" icon="http://222.186.160.172:20018/app/setu/api/random" action="" appid="-1" />
				 </msg>`
		ctx.SendChain(message.XML(data))
	})

	zero.Default().OnCommand("setu").SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		type SetuCount struct {
			SetuR18 int
			Count int
		}
		var (
			picCount int64
			r18Count int64
			tagCount int64
		)

		db.DB.Model(&setu{}).Count(&picCount)
		db.DB.Model(&setu{}).Where("r18 = ?", 1).Count(&r18Count)
		db.DB.Model(&setuTag{}).Count(&tagCount)
		ctx.SendChain(message.Text(fmt.Sprintf("setu:\r\n  count:\t%d\r\n  isR18:\t%d\r\n  tags: \t%d",picCount, r18Count, tagCount)))
	})

	r.GET("/app/setu/api/random", func(ctx *gin.Context) {
		var data setu
		db.DB.Model(&setu{}).Where("r18 = ?", "0").Order("RANDOM()").First(&data)
		data.Url = strings.ReplaceAll(data.Url, `{count}`, fmt.Sprintf("%d",rand.Intn(data.P) ))
		fmt.Print(data.Url)
		ctx.Redirect(http.StatusFound, fmt.Sprintf("%s%s",PIXIV_IMG_PROXY, data.Url))
	})

	r.GET("/app/setu/api/randomR18", func(ctx *gin.Context) {
		var data setu
		db.DB.Model(&setu{}).Where("r18 = ?", "1").Order("RANDOM()").First(&data)
		data.Url = strings.ReplaceAll(data.Url, `{count}`, fmt.Sprintf("%d",rand.Intn(data.P) ))
		fmt.Print(data.Url)
		ctx.Redirect(http.StatusFound, fmt.Sprintf("%s%s",PIXIV_IMG_PROXY, data.Url))
	})

}


func SendPixivPic(ctx *ZeroBot.Ctx, data setu){
	url := fmt.Sprintf("%s%s",PIXIV_IMG_PROXY, data.Url)
	id := ctx.SendChain(message.Reply(ctx.Event.MessageID),zero.ImageUrlMessage(url))
	if id == 0 {
		ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text("图片发送失败了"))
	}
}