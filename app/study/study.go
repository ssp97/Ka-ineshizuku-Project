/*
语料与参考代码来自：https://github.com/Giftia/ChatDACS
*/

package study

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/dbManager"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/zero"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"gorm.io/gorm"
	"math/rand"
)

type Config struct {
	Enable		bool
}

type ChatStudy struct {
	gorm.Model
	UserId  uint64 `json:"user_id" form:"user_id"`
	Ask     string `json:"ask" form:"ask" gorm:"uniqueIndex:idx_qa"`
	Answer	string `json:"answer" form:"answer" gorm:"uniqueIndex:idx_qa"`
}

type ChatMutter struct {
	gorm.Model
	Mutter  string
}

var db *dbManager.ORM
var BotName = ""


func ChatStudyBot(ctx *ZeroBot.Ctx,mustSend bool)bool{
	var replyStr string
	n := rand.Intn(1000)

	tobe := ToBeResultDo(ctx) // 提取纯文字算了 //ctx.Event.Message.String()
	//fmt.Printf("Study func data : %v", tobe)
	if tobe.direct != nil{
		replyStr = *tobe.direct
	}else if tobe.vague != nil{
		replyStr = *tobe.vague
	} else if len(tobe.result) > 0{
		replyStr = TobeResultMax(&tobe)
	}else{
		replyStr = *tobe.mutter
	}

	if n < tobe.replyPer || mustSend {
		ctx.Send(replaceAnswerSpecialStr(ctx, replyStr))
		return true
	}else {
		log.Infof("%d > %d 不发送结果",n,tobe.replyPer )
	}
	return false
}

func StudyDataAdd(q,a string,user uint64)error{
	data := ChatStudy{
		Ask: q,
		Answer: a,
		UserId: user,
	}
	result := db.DB.Create(&data)
	if result.Error != nil{
		log.Error("add study data err: %v",result.Error)
		return result.Error
	}
	return nil
}


func Init(c Config, botName string){

	if c.Enable == false{
		return
	}

	BotName = botName

	var count int64
	db = dbManager.GetDb(dbManager.DEFAULT_DB_NAME)
	db.DB.AutoMigrate(
		ChatStudy{},
		ChatMutter{},
		)

	db.DB.Model(&ChatStudy{}).Count(&count)
	if count < 10000{
		log.Debugf("Init chat study data")
		initStudyData()
	}
	db.DB.Model(&ChatMutter{}).Count(&count)
	if count < 50{
		log.Debugf("Init chat mutter data")
		initMutterData()
	}

	zero.Default().On("message", ZeroBot.OnlyToMe).SetBlock(true).SetPriority(9998).Handle(func(ctx *ZeroBot.Ctx) {
		if ChatStudyBot(ctx,true) == true{
			return
		}
	})

	zero.Default().On("message").SetBlock(true).SetPriority(9999).Handle(func(ctx *ZeroBot.Ctx) {
		if ChatStudyBot(ctx,false) == true{
			return
		}
	})

	zero.Default().OnRegex("^如果有人跟你说(.*) 你要回答(.*)$", ZeroBot.OnlyToMe).FirstPriority().SetBlock(true).Handle(func(ctx *ZeroBot.Ctx) {
		q := ctx.State["regex_matched"].([]string)[1]
		a := ctx.State["regex_matched"].([]string)[2]
		log.Debugf("%s -> %s", q,a)
		err := StudyDataAdd(q,a, uint64(ctx.Event.UserID))
		if err != nil{
			return
		}
		ctx.Send(fmt.Sprintf("心中默念%s，%s",q,a))
		return
	})

	zero.Default().OnRegex("^问：(.*) 答：(.*)$").FirstPriority().SetBlock(true).Handle(func(ctx *ZeroBot.Ctx) {
		q := ctx.State["regex_matched"].([]string)[1]
		a := ctx.State["regex_matched"].([]string)[2]
		log.Debugf("%s -> %s", q,a)
		err := StudyDataAdd(q,a, uint64(ctx.Event.UserID))
		if err != nil{
			return
		}
		ctx.Send(fmt.Sprintf("心中默念%s，%s",q,a))
		return
	})

	// 下次听到有人说：xx的话你要回答：xx

}
