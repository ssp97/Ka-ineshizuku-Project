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
	url := fmt.Sprintf("%s/v2/pixiv/user/%d", API, userId)

	res, err := grequests.Get(url, &grequests.RequestOptions{})

	if err != nil{
		return nil, err
	}
	data := res.Bytes()
	json := gjson.ParseBytes(data)

	if json.Get("error").Exists(){
		return nil, errors.New("404")
	}

	user.Name = json.Get("data.name").String()
	user.Id = json.Get("data.id").Int()

	return &user, err
}

func GetUserAllIllust(userId int64)(*[]Illust){
	var(
		page = 0
		list []Illust
	)

	for {
		url := fmt.Sprintf("%s/v2/pixiv/user/%d/illusts", API, userId)

		res, err := grequests.Get(url, &grequests.RequestOptions{
			Params: map[string]string{
				"page" : fmt.Sprintf("%d", page),
			},
			UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 Edg/98.0.1108.56",
		})
		if err != nil{
			//fmt.Println(err)
			return nil
		}
		data := res.Bytes()
		json := gjson.ParseBytes(data)
		fmt.Printf("%s\r\n", data)
		m := json.Get("data.illusts").Array()
		fmt.Println(m)
		for i := range m {
			ill := m[i]
			tags := ill.Get("tags").Array()

			//fmt.Println(m[i])
			data := Illust{
				Pid: int(ill.Get("id").Int()),
				P:	int(ill.Get("pageCount").Int()),
				Title: ill.Get("title").String(),
				//UserId: int(ill.Get("user.id").Int()),
				//UserAccount: ill.Get("user.account").String(),
				//UserName: ill.Get("user.name").String(),
				////Url: m[i].Get(),
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
