package vitsTts

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
)

type TtsRsp struct {
	Code                 int          `json:"code"`
	Audio                string       `json:"audio"`
}


func Decode(data []byte)([]byte, string){
	j := TtsRsp{}
	err := json.Unmarshal(data, &j)
	if err!= nil{
		return nil,""
	}

	b64 := j.Audio
	audio,err := base64.StdEncoding.DecodeString(b64)
	if err!= nil{
		return nil,""
	}

	return audio, ""
}


func Request(url, npc, txt string)([]byte,string){
	opt := grequests.RequestOptions{
		Headers: map[string]string{
			"content-type":"application/json",
		},
		Params: map[string]string{
			"npc": npc,
			"txt": txt,
		},
	}

	r,err := grequests.Get(url+"/tts", &opt)
	if err!= nil{
		fmt.Println("err", err)
		return nil, ""
	}
	audio,txt := Decode(r.Bytes())
	fmt.Println(audio)
	return audio, txt
}

func main(){
	fmt.Println("test")
	Request("http://192.168.123.171:8000", "可莉", "我叼你妈的")

}