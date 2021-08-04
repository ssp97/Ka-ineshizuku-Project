package study

import (
	"fmt"
	"github.com/ssp97/ZeroBot-Plugin/pkg/jieba"
	log "github.com/sirupsen/logrus"
	"github.com/wdvxdr1123/ZeroBot/message"
	"math/rand"
	"sync"
)

type TobeResult struct { 		// 待选结果
	direct 		*string			// 直接搜索
	vague 		*string			// 模糊搜索
	participle	[]string		// 分词结果
	result 		map[string]int	// 分词搜索
	mutter		*string			// 敷衍了事
	replyPer	int				// 回复概率 千分比
}

func TobeResultAdd(t *TobeResult, d string){
	n,ok := t.result[d]
	if ok{
		n++
		t.result[d] = n
	}else{
		t.result[d] = 1
	}
}

// 选词器
func TobeResultMax(t *TobeResult) string{
	maxList := []string{}
	maxCount := 0

	for k, v := range t.result {
		if v > maxCount{
			maxList = []string{}
			maxList = append(maxList, k)
			log.Infof("%v", maxList)
			maxCount = v
		}else if v == maxCount {
			maxList = append(maxList, k)
		}
	}
	log.Infof("%v",maxList)
	l := len(maxList)
	log.Info(l)
	n := rand.Intn(l)
	return maxList[n]
}

func ToBeResultDo(message message.Message)(TobeResult){
	var wg sync.WaitGroup
	tobe := TobeResult{result: map[string]int{}}

	allStr := message.String()
	extStr := message.ExtractPlainText()

	// 直接搜索
	wg.Add(1)
	go func() {
		defer wg.Done()
		var data ChatStudy
		result := db.DB.Model(&ChatStudy{}).Where("ask = ?", allStr).Order("RANDOM()").First(&data)
		if result.Error == nil{
			//log.Debugf("直接搜索到%v,%v", result.Error,data.Answer)
			tobe.direct = &data.Answer
		}
	}()
	// 模糊搜索
	wg.Add(1)
	go func() {
		defer wg.Done()
		var data ChatStudy
		result := db.DB.Model(&ChatStudy{}).Where("ask like ?", fmt.Sprintf("%%%s%%", allStr)).
			Or("ask like ?", fmt.Sprintf("%s%%", allStr)).
			Or("ask like ?", fmt.Sprintf("%%%s", allStr)).
			Or("ask = ?", fmt.Sprintf("%s", allStr)).
			Order("RANDOM()").First(&data)
		if result.Error == nil{
			tobe.vague = &data.Answer
		}
	}()
	// 分词模糊搜索
	wg.Add(1)
	go func() {
		defer wg.Done()
		d := jieba.Seg.Extract(extStr, 10)
		tobe.participle = d
		for _, s := range d {
			var studyData []ChatStudy
			log.Debugf("%s", s)
			db.DB.Model(&ChatStudy{}).Where("ask like ?", fmt.Sprintf("%%%s%%", s)).
				Or("ask like ?", fmt.Sprintf("%s%%", s)).
				Or("ask like ?", fmt.Sprintf("%%%s", s)).
				Or("ask = ?", fmt.Sprintf("%s", s)).
				Find(&studyData)
			for _, datum := range studyData {
				TobeResultAdd(&tobe, datum.Answer)
			}
		}
	}()
	// 敷衍而已
	wg.Add(1)
	go func() {
		defer wg.Done()
		var mutter ChatMutter
		result := db.DB.Model(&ChatMutter{}).Order("RANDOM()").First(&mutter)
		if result.Error == nil{
			tobe.mutter = &mutter.Mutter
		}
	}()
	wg.Wait()

	// 概率表
	if tobe.direct != nil{
		tobe.replyPer = 80
	}else if tobe.vague != nil{
		tobe.replyPer = 40
	} else if len(tobe.result) >= 0{
		tobe.replyPer = 20
	}else{
		tobe.replyPer = 5
	}

	return tobe
}