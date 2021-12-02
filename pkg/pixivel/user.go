package pixivel

import (
	"errors"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/tidwall/gjson"
	"strings"
)


type PixivelUser struct {
	Name string
	Id int64
}

func GetUserInfo(userId int64)(*PixivelUser , error){
	var user PixivelUser
	url := fmt.Sprintf("%s/pixiv", API)

	res, err := grequests.Get(url, &grequests.RequestOptions{
		Params: map[string]string{
			"type" : "member",
			"id" : fmt.Sprintf("%d", userId),
		},
	})

	if err != nil{
		return nil, err
	}
	data := res.Bytes()
	json := gjson.ParseBytes(data)

	if json.Get("error").Exists(){
		return nil, errors.New("404")
	}

	user.Name = json.Get("user.name").String()
	user.Id = json.Get("user.id").Int()

	return &user, err
}

func GetUserAllIllust(userId int64)(*[]Illust){
	var(
		page = 0
		list []Illust
	)

	for {
		url := fmt.Sprintf("%s/pixiv", API)

		res, err := grequests.Get(url, &grequests.RequestOptions{
			Params: map[string]string{
				"type" : "member_illust",
				"id" : fmt.Sprintf("%d", userId),
				"page" : fmt.Sprintf("%d", page),
			},
		})
		if err != nil{
			return nil
		}
		data := res.Bytes()
		json := gjson.ParseBytes(data)

		m := json.Get("illusts").Array()
		for i := range m {
			ill := m[i]
			tags := ill.Get("tags").Array()

			//fmt.Println(m[i])
			data := Illust{
				Pid: int(ill.Get("id").Int()),
				P:	int(ill.Get("page_count").Int()),
				Title: ill.Get("title").String(),
				UserId: int(ill.Get("user.id").Int()),
				UserAccount: ill.Get("user.account").String(),
				UserName: ill.Get("user.name").String(),
				//Url: m[i].Get(),
				R18: 0,
				Width: int(ill.Get("width").Int()),
				Height: int(ill.Get("height").Int()),
				Caption: ill.Get("caption").String(),
			}
			for j:= range tags{
				tag :=tags[j].Get("name").String()
				if tag == "R-18"{
					data.R18 = 1
				}
				data.Tags = append(data.Tags, tag)
			}
			if data.P > 1{
				somePages := ill.Get("meta_pages").Array()
				url := somePages[0].Get("image_urls.original").String()
				url = strings.ReplaceAll(url, "https://i.pximg.net", "")
				url = strings.ReplaceAll(url, "_p0", "_p{count}")
				data.Url = url
			}else{
				singlePage := ill.Get("meta_single_page")
				url := singlePage.Get("original_image_url").String()
				url = strings.ReplaceAll(url, "https://i.pximg.net", "")
				url = strings.ReplaceAll(url, "_p0", "_p{count}")
				data.Url = url
			}
			//fmt.Println(data)
			list = append(list, data)
		}

		if json.Get("next_url").String() == ""{
			break
		}
		page ++
	}
	return &list
}
