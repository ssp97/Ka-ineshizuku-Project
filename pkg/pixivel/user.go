package pixivel

import (
	"errors"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/tidwall/gjson"
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

func GetUserAllIllust(userId int64){
	var(
		page = 0
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
			return
		}
		data := res.Bytes()
		json := gjson.ParseBytes(data)

		m := json.Get("illusts").Array()
		for i := range m {
			fmt.Println(m[i])
		}

		if json.Get("next_url").String() == ""{
			break
		}
		page ++
	}

}
