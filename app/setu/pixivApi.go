package setu

import (
	"fmt"
	"github.com/everpcpc/pixiv"
	"github.com/ssp97/Ka-ineshizuku-Project/app/publicModels"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	//"strings"
	"time"
)

const PIXAPI_TOKEN_KEY = "setu_pixivapi_token"
const PIXAPI_RE_TOKEN_KEY = "setu_pixivapi_refresh_token"
const PIXAPI_TOKEN_TIME_KEY = "setu_pixivapi_time"

type SetuPixivApi struct {
	App *pixiv.AppPixivAPI
}

func (pixivapi *SetuPixivApi) Init() {
	
	pixiv.HookAuth(func(s string, s2 string, t time.Time) error {
		publicModels.SetSetting(PIXAPI_TOKEN_KEY, s)
		publicModels.SetSetting(PIXAPI_RE_TOKEN_KEY, s2)
		publicModels.SetSetting(PIXAPI_TOKEN_TIME_KEY, strconv.FormatInt(t.Unix(),10))
		return nil
	})

	go func() {
		for{
			time.Sleep(time.Second * 1800)

			pixivapi.reLogin()

		}
	}()
}

func (pixivapi *SetuPixivApi) reLogin(){
	_, token := publicModels.GetSetting(PIXAPI_TOKEN_KEY)
	_, reToken := publicModels.GetSetting(PIXAPI_RE_TOKEN_KEY)
	_, tokenTime := publicModels.GetSetting(PIXAPI_TOKEN_TIME_KEY)
	inttokenTime, _  := strconv.ParseInt(tokenTime, 10, 64)
	if token == ""{
		return
	}
	if reToken == ""{
		return
	}
	if tokenTime == ""{
		pixivapi.Login(token, reToken, time.Now())
	}else{
		pixivapi.Login(token, reToken, time.Unix(inttokenTime,0))
	}
}

func (pixivapi *SetuPixivApi) SetTokenAndLogin(token, retoken string)error{
	return pixivapi.Login(token, retoken, time.Now())
}

func (pixivapi *SetuPixivApi) Login(token, retoken string, time time.Time) error {

	account, err := pixiv.LoadAuth(token, retoken, time)
	if err != nil{
		log.Warn(err)
		return err
	}
	log.Info(account)
	if pixivapi.App == nil{
		pixivapi.App = pixiv.NewApp()
	}
	return nil
}

func (pixivapi *SetuPixivApi)GetUserAllPic(uid uint64)int{
	lastNext := 0
	next := 0
	count := 0
	for{
		lastNext = next
		illusts, next, err := pixivapi.App.UserIllusts(uid, "illust", next)
		if err != nil{
			log.Warn(err)
			return 0
		}
		for _, illust := range illusts {
			err := pixivapi.AddPicToDB(illust)
			if err != nil{
				log.Warn(err)
			}else{
				count++
			}
		}
		if next == lastNext{
			break
		}
	}

	return count
}

func (pixivapi *SetuPixivApi)AddPicToDB(illust pixiv.Illust) error{
	var tags []string
	isR18 := 0
	for _, t := range illust.Tags {
		if t.Name == ""{
			continue
		}
		tags = append(tags, t.Name)
		if t.Name != t.TranslatedName && t.TranslatedName != ""{
			tags = append(tags, t.TranslatedName)
		}

		if t.Name == "R18"{
			isR18 = 1
		}
	}
	fmt.Println(illust.Images)

	url := strings.ReplaceAll(illust.Images.Medium, "https://i.pximg.net/c/540x540_70/img-master/", "/img-original/")
	url = strings.ReplaceAll(url, "p0_master1200", "p{count}")

	err := addSetu(setu{
		Pid: int(illust.ID),
		P:illust.PageCount,
		Title: illust.Title,
		UserId: int(illust.User.ID),
		UserAccount: illust.User.Account,
		UserName: illust.User.Name,
		Url: url,
		R18: isR18,
		Width: illust.Width,
		Height: illust.Height,
		Tags: strings.Join(tags, ", "),
	}, tags)
	return err
}