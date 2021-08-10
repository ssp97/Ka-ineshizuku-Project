package EEAsst

import (
	"fmt"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/dbManager"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

type Config struct {
	Enable 		bool
}

type ComponentSilk struct {
	ID 			uint64		`json:"id" form:"id" gorm:"primary_key;auto_increment;"`
	UserId		uint64		`json:"user_id" form:"user_id"`
	SilkName	string	`json:"silk_name" form:"silk_name"`
	PdfUrl		string	`json:"pdf_url" form:"pdf_url"`
}

var db *dbManager.ORM

func Resistance(str string)(string, error){

	return "", nil
}


func Init(c Config){
	if c.Enable == false{
		return
	}

	db = dbManager.GetDb(dbManager.DEFAULT_DB_NAME)
	db.DB.AutoMigrate(ComponentSilk{})

	ZeroBot.OnRegex("^电阻(.*?)$").SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		code := ctx.State["regex_matched"].([]string)[1]
		result, err := getEIA96Res(code)
		if err != nil {
			//ctx.SendChain(message.Text(fmt.Sprintf("%s对应的尺寸是：%s",code,result)))
		}else{
			ctx.SendChain(message.Text(fmt.Sprintf("%s对应的电阻是：%sΩ",code,result)))
		}
	})

	ZeroBot.OnRegex("^尺寸(.*?)$").SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		code := ctx.State["regex_matched"].([]string)[1]
		result, err := getEIASize(code)
		if err != nil {
			//ctx.SendChain(message.Text(fmt.Sprintf("%s对应的尺寸是：%s",code,result)))
		}else{
			ctx.SendChain(message.Text(fmt.Sprintf("%s对应的尺寸是：%s",code,result)))
		}
	})

}
