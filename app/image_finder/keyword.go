package image_finder

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/ssp97/ZeroBot-Plugin/pkg/avoidExamine"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var PATH = "data/pixiv/cache"

type AutoGenerated struct {
	Illusts []struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Type      string `json:"type"`
		ImageUrls struct {
			SquareMedium string `json:"square_medium"`
			Medium       string `json:"medium"`
			Large        string `json:"large"`
		} `json:"image_urls"`
		Caption  string `json:"caption"`
		Restrict int    `json:"restrict"`
		User     struct {
			ID               int    `json:"id"`
			Name             string `json:"name"`
			Account          string `json:"account"`
			ProfileImageUrls struct {
				Medium string `json:"medium"`
			} `json:"profile_image_urls"`
			IsFollowed bool `json:"is_followed"`
		} `json:"user"`
		Tags []struct {
			Name           string      `json:"name"`
			TranslatedName interface{} `json:"translated_name"`
		} `json:"tags"`
		Tools          []interface{} `json:"tools"`
		PageCount      int           `json:"page_count"`
		Width          int           `json:"width"`
		Height         int           `json:"height"`
		SanityLevel    int           `json:"sanity_level"`
		XRestrict      int           `json:"x_restrict"`
		Series         interface{}   `json:"series"`
		MetaSinglePage struct {
			OriginalImageURL string `json:"original_image_url"`
		} `json:"meta_single_page,omitempty"`
		MetaPages      []interface{} `json:"meta_pages"`
		TotalView      int           `json:"total_view"`
		TotalBookmarks int           `json:"total_bookmarks"`
		IsBookmarked   bool          `json:"is_bookmarked"`
		Visible        bool          `json:"visible"`
		IsMuted        bool          `json:"is_muted"`
	} `json:"illusts"`
	NextURL         string `json:"next_url"`
	SearchSpanLimit int    `json:"search_span_limit"`
}

var SensitiveWords = map[string]bool{
	"R-18":true,
	"R-18G":true,
	"R18G":true,
	"R18":true,
}

func init() {
	zero.OnRegex(`^来张(.*)$`).//, zero.AdminPermission
		Handle(func(ctx *zero.Ctx) {
			keyword := ctx.State["regex_matched"].([]string)[1]
			if len(keyword) > 6*3{
				return
			}
			soutujson := soutuapi(keyword)

			if(len(soutujson.Illusts) <= 0){
				ctx.SendChain(message.Text("没有"))
				return
			}
			pom1 := "http://i.pixiv.cat"
			for i := 0; i < 30; i++ {
				Sensitive := false
				rannum := Suiji(len(soutujson.Illusts))
				tags := soutujson.Illusts[rannum].Tags
				for j := 0; j < len(tags); j++ {
					if SensitiveWords[tags[j].Name]{
						Sensitive = true
						break
					}
				}
				if Sensitive == false{
					rootDir, _ := os.Getwd()
					pom2 := soutujson.Illusts[rannum].ImageUrls.Medium[19:]
					file := fmt.Sprintf("%s/%s/%d",rootDir,PATH,soutujson.Illusts[rannum].ID)
					//ctx.SendChain(message.Image(pom1 + pom2),message.Text(soutujson.Illusts[rannum].ID))

					err := picDownload(ctx, file, pom1 + pom2)
					if err!= nil{
						ctx.SendChain(message.Text(fmt.Sprintf("啊偶，图片下载失败了...因为%s",err)))
						return
					}
					err = avoidExamine.PicFile(file)
					if err!= nil{
						ctx.SendChain(message.Text("啊偶，图片处理出错了，因为%s", err))
					}
					ctx.SendChain(message.Image(fmt.Sprintf("file:///%s/%s/%d",rootDir,PATH,soutujson.Illusts[rannum].ID)),
						message.Text(fmt.Sprintf("%d-%s",soutujson.Illusts[rannum].ID,soutujson.Illusts[rannum].Title)))

					go func(file string) {
						time.Sleep(time.Second * 120)
						os.Remove(file)
					}(file)

					return
				}
			}
			ctx.SendChain(message.Text("没有"))

		})

}

// soutuapi 请求api
func soutuapi(keyword string) *AutoGenerated {

	url := "https://api.pixivel.moe/pixiv?type=search&page=0&mode=partial_match_for_tags&word=" + keyword
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("accept", "application/json, text/plain, */*")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	result := &AutoGenerated{}
	if err := json.NewDecoder(res.Body).Decode(result); err != nil {
		panic(err)
	}
	return result
}

// Suiji 从json里的30条数据中随机获取一条返回
func Suiji(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func picDownload(ctx *zero.Ctx,file string, url string)(err error){
	f,err := os.Create(file)
	if err!= nil{
		return
	}
	defer f.Close()
	res, err := http.Get(url)
	if err!= nil{
		return
	}
	io.Copy(f, res.Body)
	return
}