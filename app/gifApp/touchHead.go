package gifApp

import (
	"embed"
	"fmt"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/OicqUtils"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/TypeUtils"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/softwarebackend"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/webp"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	"path"
)

//go:embed static
var static embed.FS

var PATH = path.Join("data","cache")

type faceOffset struct {
	x,y,w,h float64
}

func isInPalette(p color.Palette, c color.Color) int {
	ret := -1
	for i, v := range p {
		if v == c {
			return i
		}
	}
	return ret
}


func make(face *image.Image,_path string)  {
	Sprite,err := static.Open(path.Join("static", "sprite.png"))
	if err != nil{
		panic(err)
	}
	defer Sprite.Close()
	fmt.Println(Sprite)


	SpritePng,err := png.Decode(Sprite)
	if err != nil{
		panic(err)
	}

	g := gif.GIF{LoopCount: 0}
	offset := []faceOffset{
		{ x: 0, y: 0, w: 0, h: 0 },
		{ x: -4, y: 12, w: 4, h: -12 },
		{ x: -12, y: 18, w: 12, h: -18 },
		{ x: -8, y: 12, w: 4, h: -12 },
		{ x: -4, y: 0, w: 0, h: 0 },
	}

	for i := 0; i < 5; i++ {
		backend := softwarebackend.New(112, 112)
		cv := canvas.New(backend)
		cv.SetFillStyle("#FFFFFF")
		cv.FillRect(0,0, 112, 112)
		cv.DrawImage(*face,offset[i].x+5,offset[i].y+5, offset[i].w+122, offset[i].h+122)
		cv.DrawImage(SpritePng,112*float64(i),0,112,112,0,0,112,112)

		pleated := image.NewPaletted(backend.Image.Bounds(), palette.Plan9)
		draw.FloydSteinberg.Draw(pleated, backend.Image.Bounds(), backend.Image, image.ZP)

		g.Image = append(g.Image, pleated)
		g.Delay = append(g.Delay, 5)
	}

	f, err := os.OpenFile(_path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}

	err = gif.EncodeAll(f, &g)
	if err != nil {
		panic(err)
	}
}

func init() {
	root, _ := os.Getwd()
	ZeroBot.OnRegex("^摸头").SetBlock(true).SetPriority(60).Handle(func(ctx *ZeroBot.Ctx) {
		//str := ctx.State["regex_matched"].([]string)[1]
		userId := ctx.Event.UserID
		for _, segment := range ctx.Event.Message {
			if segment.Type == "at"{
				userId = TypeUtils.StrToInt(segment.Data["qq"])
			}
		}
		file := fmt.Sprintf("%d.jpg",userId)
		_path := path.Join(root, PATH, file)

		faceImg := OicqUtils.GetQQFaceImg(userId)
		if faceImg == nil{
			return
		}
		make(faceImg, _path)
		ctx.SendChain(message.Image(fmt.Sprintf("file:///%s" ,_path)))

	})

}
