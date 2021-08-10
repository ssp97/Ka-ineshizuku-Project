package haveAFriend

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/nfnt/resize"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/OicqUtils"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/zero"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/softwarebackend"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	_ "golang.org/x/image/webp"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)
var PATH = path.Join("data", "cache")
//const FONT_PATH = "static/font/NotoSansSC-Regular.ttf"
var FONT_PATH = path.Join("static","font","simhei.ttf")

////go:embed font
//var fonts embed.FS



func make(str1, str2 string,face *image.Image,_path string){

	t := time.Now().Format("15:04")
	mul := 2

	fontNoto,err := os.ReadFile(FONT_PATH)
	if err != nil{
		fmt.Println(err)
	}
	font, err := freetype.ParseFont(fontNoto)
	if err != nil {
		fmt.Println(err)
		return
	}

	backend := softwarebackend.New(350*2*mul, 80*2*mul)
	//backend.MSAA = 1
	cv := canvas.New(backend)

	cv.SetFillStyle("#FFFFFF")
	//cv.SetStrokeStyle("#000000")
	cv.FillRect(0,0, float64(350*2*mul), float64(80*2*mul))

	cv.SetTextAlign(canvas.Left)

	cv.SetFont(font,20*2*float64(mul))
	cv.SetLineWidth(0)
	cv.SetFillStyle("#000000")
	cv.FillText(str1,90.25*2*float64(mul),35.25*2*float64(mul))

	cv.SetFont(font,16*2*float64(mul))
	cv.SetFillStyle("#716F81")
	cv.FillText(str2, 90.25*2*float64(mul), 60.25*2*float64(mul))

	cv.SetFont(font, 13*2*float64(mul))
	cv.FillText(t, 300.25*2*float64(mul), 35.25*2*float64(mul))

	cv.BeginPath()
	cv.Arc(80*float64(mul),80*float64(mul),28*2*float64(mul),0,math.Pi*2,true)
	cv.Fill()
	cv.Clip()
	cv.DrawImage(*face, 10*2*float64(mul),10*2*float64(mul),60*2*float64(mul),60*2*float64(mul))
	cv.ClosePath()

	m := resize.Thumbnail(350*2, 80*2, backend.Image, resize.Lanczos3)

	f, err := os.OpenFile(_path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}

	err = jpeg.Encode(f, m, &jpeg.Options{Quality: 80})
	if err != nil {
		panic(err)
	}
}


func init(){
	root, _ := os.Getwd()

	zero.Default().OnRegex("^我有个朋友说(.*?)$").SetBlock(true).SetPriority(20).Handle(func(ctx *ZeroBot.Ctx) {
		str := ctx.State["regex_matched"].([]string)[1]
		userId := ctx.Event.UserID
		for _, segment := range ctx.Event.Message {
			if segment.Type == "at"{
				userId = strToInt(segment.Data["qq"])
				str = strings.ReplaceAll(str, segment.String(),"")
			}
		}
		name := ctx.GetGroupMemberInfo(ctx.Event.GroupID,userId,true).Get("nickname").String()
		file := fmt.Sprintf("%d.jpg",userId)
		p := path.Join(root, PATH, file)
		img := OicqUtils.GetQQFaceImg(userId)
		//make(name, str,img,p)
		//err := makeByGpu(name, str,img,p)
		//if err != nil{
		//	make(name, str,img,p)
		//}
		make(name, str,img,p)
		ctx.SendChain(message.Image(fmt.Sprintf("file:///%s" ,p)))
	})
}

func strToInt(str string) int64 {
	val, _ := strconv.ParseInt(str, 10, 64)
	return val
}

//func main(){
//	root, _ := os.Getwd()
//	img := getQQFaceImg(38263547)
//
//}