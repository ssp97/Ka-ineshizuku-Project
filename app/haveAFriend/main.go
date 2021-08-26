package haveAFriend

import (
	"encoding/base64"
	"fmt"
	"github.com/golang/freetype"
	"github.com/nfnt/resize"
	"github.com/orcaman/writerseeker"
	log "github.com/sirupsen/logrus"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/OicqUtils"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/avoidExamine"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/zero"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/softwarebackend"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	_ "golang.org/x/image/webp"
	"image"
	"image/color"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"io/ioutil"
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



func haveAFriendMake(str1, str2 string,face *image.Image)(b *string){

	t := time.Now().Format("15:04")
	_t := time.Now()
	mul := 1.5

	fontNoto,err := os.ReadFile(FONT_PATH)
	if err != nil{
		fmt.Println(err)
	}
	font, err := freetype.ParseFont(fontNoto)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Debugln(time.Since(_t)) //70ms
	backend := softwarebackend.New(int(350*2*mul), int(80*2*mul))
	//backend.MSAA = 1
	cv := canvas.New(backend)
	log.Debugln(time.Since(_t)) //74ms
	//cv.SetFillStyle("#FFFFFF")
	//cv.SetStrokeStyle("#000000")
	//cv.FillRect(0,0, float64(350*2*mul), float64(80*2*mul))

	// 在armv7l中，这样可以将总时间从1.1s缩减到600ms内
	for x := 0; x < int(350*2*mul); x++ {
		for y := 0; y < int(80*2*mul); y++ {
			backend.Image.SetRGBA(x,y,color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 0,
			})
		}
	}

	//for i := 0; i < col; i++ {
	//
	//}
	log.Debugln(time.Since(_t)) //703ms
	cv.SetTextAlign(canvas.Left)

	cv.SetFont(font,20*2*float64(mul))
	cv.SetLineWidth(0)
	cv.SetFillStyle("#000000")
	cv.FillText(str1,90.25*2*float64(mul),35.25*2*float64(mul))
	//log.Debugln(time.Since(_t)) 806ms
	cv.SetFont(font,16*2*float64(mul))
	cv.SetFillStyle("#716F81")
	cv.FillText(str2, 90.25*2*float64(mul), 60.25*2*float64(mul))
	//log.Debugln(time.Since(_t)) 838
	cv.SetFont(font, 13*2*float64(mul))
	cv.FillText(t, 300.25*2*float64(mul), 35.25*2*float64(mul))
	//log.Debugln(time.Since(_t)) 869
	cv.BeginPath()
	cv.Arc(80*float64(mul),80*float64(mul),28*2*float64(mul),0,math.Pi*2,true)
	cv.Fill()
	cv.Clip()
	cv.DrawImage(*face, 10*2*float64(mul),10*2*float64(mul),60*2*float64(mul),60*2*float64(mul))
	cv.ClosePath()
	//log.Debugln(time.Since(_t)) 1040
	m := resize.Thumbnail(350*2, 80*2, backend.Image, resize.Lanczos3)
	//log.Debugln(time.Since(_t)) 1119ms
	//f, err := os.OpenFile(_path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	//if err != nil {
	//	panic(err)
	//}
	w := new(writerseeker.WriterSeeker)


	err = jpeg.Encode(w, m, &jpeg.Options{Quality: 80})
	if err != nil {
		panic(err)
	}
	d, err := ioutil.ReadAll(w.Reader())
	if err != nil{
		panic(err)
	}
	d = avoidExamine.PicByte(d)
	bs := base64.StdEncoding.EncodeToString(d)
	b = &bs
	return
}


func init(){
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
		img := OicqUtils.GetQQFaceImg(userId)
		t := time.Now()
		b := haveAFriendMake(name, str,img)
		ctx.SendChain(zero.ImageBase64Message(b), message.Text(fmt.Sprintf("%v",time.Since(t))))
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