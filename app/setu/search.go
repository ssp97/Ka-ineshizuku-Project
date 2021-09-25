package setu

import (
	"fmt"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/gocc"
	"math/rand"
	"strings"
)

func searchRandom(tag string, r18 int)(setu, error){
	var tagS2t = gocc.S2t(tag)
	var tagLike = fmt.Sprintf("%%%s%%", tag)
	var tagS2tLkie = fmt.Sprintf("%%%s%%", tagS2t)
	var data setu

	result := db.DB.Table("setus").Joins("join setu_tags ON setu_tags.pid = setus.pid").
		Where("r18 = ? and (setu_tags.tag = ? or setu_tags.tag = ?)",r18, tag, tagS2t).
		Order("RANDOM()").First(&data)
	if result.Error == nil{
		data.Url = strings.ReplaceAll(data.Url, `{count}`, fmt.Sprintf("%d",rand.Intn(data.P) ))
		return data, nil
	}
	result = db.DB.Model(&setu{}).Where("r18 = ? and (title like ? or title like ?)",r18, tagLike, tagS2tLkie).Order("RANDOM()").First(&data)
	if result.Error == nil{
		data.Url = strings.ReplaceAll(data.Url, `{count}`, fmt.Sprintf("%d",rand.Intn(data.P) ))
		return data, nil
	}
	return data, result.Error
}
