package setu

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/gocc"
	"math/rand"
	"strings"
)

func searchRandom(tag string, r18 int)(setu, error){
	var tagS2t = gocc.S2t(tag)
	var tagLike = fmt.Sprintf("%%%s%%", tag)
	var tagS2tLkie = fmt.Sprintf("%%%s%%", tagS2t)
	var data setu
	var tagTranslated setuTagTranslated

	main := db.DB.Table("setus").Joins("join setu_tags ON setu_tags.pid = setus.pid")

	result := db.DB.Model(&setuTagTranslated{}).Where("zh like ?", fmt.Sprintf("%%%s%%", tag)).First(&tagTranslated)
	if result.Error == nil{
		main.Where("r18 = ? and (setu_tags.tag = ? or setu_tags.tag = ? or setu_tags.tag = ?)",r18, tag, tagS2t, tagTranslated.Src)
	}else{
		main.Where("r18 = ? and (setu_tags.tag = ? or setu_tags.tag = ?)",r18, tag, tagS2t)
	}

	result = main.Order("RANDOM()").First(&data)
	if result.Error == nil{
		//data.Url = strings.ReplaceAll(data.Url, "/img-original", "/c/600x1200_90/img-master")
		//data.Url = strings.ReplaceAll(data.Url, "{count}.png", "{count}_master1200.jpg")
		//data.Url = strings.ReplaceAll(data.Url, "{count}.jpg", "{count}_master1200.jpg")
		data.Url = strings.ReplaceAll(data.Url, `{count}`, fmt.Sprintf("%d",rand.Intn(data.P) ))
		return data, nil
	}
	result = db.DB.Model(&setu{}).Where("r18 = ? and (title like ? or title like ?)",r18, tagLike, tagS2tLkie).Order("RANDOM()").First(&data)

	log.Warn(data.Url)
	if result.Error == nil{
		//data.Url = strings.ReplaceAll(data.Url, "/img-original", "/c/600x1200_90/img-master")
		//data.Url = strings.ReplaceAll(data.Url, "{count}.png", "{count}_master1200.jpg")
		//data.Url = strings.ReplaceAll(data.Url, "{count}.jpg", "{count}_master1200.jpg")
		//data.Url = strings.ReplaceAll(data.Url, `{count}`, fmt.Sprintf("%d",rand.Intn(data.P) ))
		data.Url = fmt.Sprintf("%d-%d.jpg", data.Pid, rand.Intn(data.P))
		return data, nil
	}
	return data, result.Error
}

func addSetu(setu2 setu, tags []string)(error){
	for i := range tags{
		tag := setuTag{
			Pid: setu2.Pid,
			Tag: tags[i],
		}
		db.DB.Model(&setuTag{}).Create(&tag)
	}
	result := db.DB.Table("setus").Create(&setu2)
	return result.Error
}
