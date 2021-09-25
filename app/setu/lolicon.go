package setu

import (
	"github.com/lxzan/hasaki"
	"github.com/tidwall/gjson"
	"strings"
)

const (
	API = "https://api.lolicon.app/setu/v2"
)

type loliconStruct struct {
	error string
	data struct{

	}
}

func loliconSearch(keyword string, r18 int)(pid int, url string){
	data, err := hasaki.Get(API).Send(hasaki.Any{
		"keyword": keyword,
		"r18": r18,
	}).GetBody()
	if err != nil{
		return 0, ""
	}
	json := gjson.ParseBytes(data)
	m := json.Get("data").Map()
	if len(m) == 0{
		return 0, ""
	}
	item := json.Get("data.0")
	pid = int(item.Get("pid").Int())
	url = item.Get("urls.original").Str
	url = strings.ReplaceAll(url, "https://i.pixiv.cat/", "")
	return
}
