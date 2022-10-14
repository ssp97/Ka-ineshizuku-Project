package setuai

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"github.com/levigross/grequests"
	"os"
	"strings"
)

func Decode(data []byte) []byte{
	d := string(data)
	l := strings.Split(d, "\n")
	b64 := strings.Split(l[2], ":")[1]
	img,err := base64.StdEncoding.DecodeString(b64)
	if err!= nil{
		return nil
	}
	return img
}

func Request(url string, prompt, width, height, scale, sampler, steps, seed, uc *string)([]byte,string){
	var n uint32
	binary.Read(rand.Reader, binary.LittleEndian, &n)

	jsonData := map[string]string{
		"prompt":"masterpiece, best quality, rating:explicit,bdsm,loli,guro,photo,aqua_eyes,naughty_face,flat_chest,spread_legs,no_bra,white_thighhighs,open_mouth,rolleyes,bed_background,green_hair,green_hair,10years,anal,femdom,saliva,squirting,torture",
		"width":"512",
		"height":"768",
		"scale":"12",
		"sampler":"k_euler_ancestral",
		"steps":"20",
		"seed":"2664075441",
		"n_samples":"1",
		"ucPreset":"0",
		"uc":"lowres, bad anatomy, bad hands, text, error, missing fingers, extra digit, fewer digits, cropped, worst quality, low quality, normal quality, jpeg artifacts, signature, watermark, username, blurry, lowres, bad anatomy, bad hands, text,error, missing fngers,extra digt ,fewer digits,cropped, wort quality ,low quality,normal quality, jpeg artifacts,signature,watermark, username, blurry, bad feet",
	}

	if prompt!=nil {
		jsonData["prompt"] = *prompt
		print("prompt:", *prompt, "\r\n")
	}
	if width!=nil{
		jsonData["width"] = *width
	}
	if height!=nil{
		jsonData["height"] = *height
	}
	if scale!=nil{
		jsonData["scale"] = *scale
	}
	if sampler!=nil{
		jsonData["sampler"] = *sampler
	}
	if steps!=nil{
		jsonData["steps"] = *steps
	}
	if seed!=nil{
		jsonData["seed"] = *seed
	} else {
		jsonData["seed"] = fmt.Sprintf("%d", n)
		print("seed = ", jsonData["seed"], "\r\n")
	}
	if uc!=nil{
		jsonData["uc"] = *uc
	}

	opt := grequests.RequestOptions{
		Headers: map[string]string{
			"content-type":"application/json",
		},
		JSON: &jsonData,
	}

	r,err := grequests.Post(url, &opt)
	if err!=nil{
		print("post err")
		print(err)
		return nil, ""
	}

	img := Decode(r.Bytes())
	txt := fmt.Sprintf("prompt:%s\r\nseed:%s\r\n","哔~" , jsonData["seed"])//jsonData["prompt"]
	return img, txt
}


func main() {
	url := "http://192.168.123.177:6969/generate-stream"
	opt := grequests.RequestOptions{
		Headers: map[string]string{
			"content-type":"application/json",
		},
		JSON: map[string]string{"prompt":"masterpiece, best quality, rating:explicit,bdsm,loli,guro,photo,aqua_eyes,naughty_face,flat_chest,spread_legs,no_bra,white_thighhighs,open_mouth,rolleyes,bed_background,green_hair,green_hair,10years,anal,femdom,saliva,squirting,torture",
			"width":"512",
			"height":"768",
			"scale":"12",
			"sampler":"k_euler_ancestral",
			"steps":"28",
			"seed":"2664075441",
			"n_samples":"1",
			"ucPreset":"0",
			"uc":"lowres, bad anatomy, bad hands, text, error, missing fingers, extra digit, fewer digits, cropped, worst quality, low quality, normal quality, jpeg artifacts, signature, watermark, username, blurry, lowres, bad anatomy, bad hands, text,error, missing fngers,extra digt ,fewer digits,cropped, wort quality ,low quality,normal quality, jpeg artifacts,signature,watermark, username, blurry, bad feet"},

	}
	r,err := grequests.Post(url, &opt)
	if err!=nil{
		print("post err")
		print(err)
		return
	}

	img := Decode(r.Bytes())

	file, err := os.OpenFile("./test.png", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.Write(img)
}
