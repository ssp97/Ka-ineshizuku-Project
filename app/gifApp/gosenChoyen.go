package gifApp

import (
	"fmt"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/softwarebackend"
	"image/jpeg"
	"math"
	"os"
	"path"
)

// 5000兆円欲しい！ 風格

//var PATH = path.Join("data", "cache")
//const FONT_PATH = "static/font/NotoSansSC-Regular.ttf"
var FONT_PATH_GOSEN_HEAVY = path.Join("static","font","SourceHanSansSC-Heavy.ttf")
var FONT_PATH_GOSEN_BOLD = path.Join("static","font","SourceHanSerifSC-Bold.ttf")

func init() {
	//gosenChoyenMake("燃起來了","真的啊",0)
}

func gosenChoyenMake(upper,lower string, offsetWidth float64){
	shsans ,err := ttfFontLoad(FONT_PATH_GOSEN_HEAVY)
	if err != nil{
		panic(err)
	}
	shserif ,err := ttfFontLoad(FONT_PATH_GOSEN_BOLD)
	if err != nil{
		panic(err)
	}

	backend := softwarebackend.New(1000, 1000)
	ctx := canvas.New(backend)

	ctx.SetFont(shsans,100)
	upperWidth := ctx.MeasureText(upper).Width
	ctx.SetFont(shserif ,100)
	lowerWidth := ctx.MeasureText(lower).Width

	height := 270
	width := math.Max(upperWidth + 80, lowerWidth + offsetWidth + 90)

	fmt.Printf("%f\r\n",width)

	backend = softwarebackend.New(int(width), height)
	ctx = canvas.New(backend)
	ctx.SetLineJoin(canvas.Round)
	ctx.SetFillStyle("#FFFFFF")

	ctx.FillRect(0, 0, width, float64(height))
	ctx.SetTransform(1, 0, -0.4, 1, 0, 0)

	var posx, posy float64

	ctx.SetFont(shsans, 100)
	posx = 70
	posy = 100

	ctx.SetStrokeStyle("#000")
	ctx.SetLineWidth(18)
	ctx.StrokeText(upper, posx + 4, posy + 3)

	grad := ctx.CreateLinearGradient(0, 24, 0, 122)
	grad.AddColorStop(0,0,"rgb(0,15,36)")
	grad.AddColorStop(0.10, "rgb(255,255,255)")
	grad.AddColorStop(0.18, "rgb(55,58,59)")
	grad.AddColorStop(0.25, "rgb(55,58,59)")
	grad.AddColorStop(0.5, "rgb(200,200,200)")
	grad.AddColorStop(0.75, "rgb(55,58,59)")
	grad.AddColorStop(0.85, "rgb(25,20,31)")
	grad.AddColorStop(0.91, "rgb(240,240,240)")
	grad.AddColorStop(0.95, "rgb(166,175,194)")
	grad.AddColorStop(1, "rgb(50,50,50)")
	ctx.SetStrokeStyle(grad)
	ctx.SetLineWidth(17)
	ctx.StrokeText(upper, posx + 4, posy + 3)

	//m := resize.Thumbnail(uint(width), uint(height), backend.Image, resize.Lanczos3)
	//log.Debugln(time.Since(_t)) 1119ms
	f, err := os.OpenFile("test.jpg", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}

	err = jpeg.Encode(f, backend.Image, &jpeg.Options{Quality: 80})
	if err != nil {
		panic(err)
	}

}