package haveAFriend

import (
	"embed"
	"fmt"
	"github.com/golang/freetype"
	"github.com/nfnt/resize"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/softwarebackend"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"image"
	"image/jpeg"
	_ "image/gif"
	_ "image/png"
	_ "golang.org/x/image/webp"
	"math"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)
const PATH = "data/cache"
//go:embed font
var fonts embed.FS



func getQQFaceImg(qq_id int64)(img *image.Image){
	url := fmt.Sprintf("https://api.sumt.cn/api/qq.logo.php?qq=%d", qq_id)
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	_img, _,err := image.Decode(res.Body)
	if err != nil{
		panic(err)
	}
	img = &_img
	//fmt.Println(_img)
	return
}

//func makeByGpu(str1, str2 string,face *image.Image,_path string)(err error){
//	t := time.Now().Format("15:04")
//
//	fontNoto,err := fonts.ReadFile("font/NotoSansSC-Regular.ttf")
//	if err != nil{
//		fmt.Println(err)
//	}
//	font, err := freetype.ParseFont(fontNoto)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	ctx,err := goglbackend.NewGLContext()
//	if err!=nil{
//		fmt.Println(err)
//		return
//	}
//	backend,err := goglbackend.New(0,0,350*2,80*2, ctx)
//	if err!=nil{
//		fmt.Println(err)
//		return
//	}
//	cv := canvas.New(backend)
//
//	cv.SetFillStyle("#FFFFFF")
//	//cv.SetStrokeStyle("#000000")
//	cv.FillRect(0,0,350*2,80*2)
//
//	cv.SetTextAlign(canvas.Left)
//
//	cv.SetFont(font,20*2)
//	cv.SetLineWidth(0)
//	cv.SetFillStyle("#000000")
//	cv.FillText(str1,90.25*2,35.25*2)
//
//	cv.SetFont(font,16*2)
//	cv.SetFillStyle("#716F81")
//	cv.FillText(str2, 90.25*2, 60.25*2)
//
//	cv.SetFont(font, 13*2)
//	cv.FillText(t, 300.25*2, 35.25*2)
//
//	cv.BeginPath()
//	cv.Arc(80,80,28*2,0,math.Pi*2,true)
//	cv.Fill()
//	cv.Clip()
//	cv.DrawImage(*face, 10*2,10*2,60*2,60*2)
//	cv.ClosePath()
//
//
//
//	f, err := os.OpenFile(_path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
//	if err != nil {
//		panic(err)
//	}
//
//	err = jpeg.Encode(f, backend.GetImageData(0,0,350*2,80*2), &jpeg.Options{Quality: 80})
//	if err != nil {
//		panic(err)
//	}
//	return
//}

func make(str1, str2 string,face *image.Image,_path string){

	t := time.Now().Format("15:04")
	mul := 2

	fontNoto,err := fonts.ReadFile("font/NotoSansSC-Regular.ttf")
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

	zero.OnRegex("^^我有个朋友说(.*?)$").SetBlock(true).SetPriority(20).Handle(func(ctx *zero.Ctx) {
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
		img := getQQFaceImg(userId)
		//make(name, str,img,p)
		//err := makeByGpu(name, str,img,p)
		//if err != nil{
		//	make(name, str,img,p)
		//}
		make(name, str,img,p)
		ctx.SendChain(message.Image(fmt.Sprintf("file://%s" ,p)))
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