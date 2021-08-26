package gifApp

import (
	"encoding/base64"
	"fmt"
	"github.com/orcaman/writerseeker"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/OicqUtils"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/TypeUtils"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/avoidExamine"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/fsUtils"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/zero"
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
	"io/ioutil"
	"os"
	"path"
	"sync"
	"time"
)

////go:embed static
//var static embed.FS

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


func touchHeadMake(face *image.Image)(b *string)  {
	Sprite,err := os.Open(path.Join(fsUtils.Getwd(), "static", "img", "sprite.png"))
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

	g.Image = make([]*image.Paletted, 5)
	g.Delay = make([]int ,5)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			backend := softwarebackend.New(112, 112)
			cv := canvas.New(backend)
			cv.SetFillStyle("#FFFFFF")
			cv.FillRect(0,0, 112, 112)
			cv.DrawImage(*face,offset[i].x+5,offset[i].y+5, offset[i].w+122, offset[i].h+122)
			cv.DrawImage(SpritePng,112*float64(i),0,112,112,0,0,112,112)

			pleated := image.NewPaletted(backend.Image.Bounds(), palette.Plan9)
			draw.FloydSteinberg.Draw(pleated, backend.Image.Bounds(), backend.Image, image.ZP)

			g.Image[i] = pleated
			g.Delay[i] = 5

			//g.Image = append(g.Image, pleated)
			//g.Delay = append(g.Delay, 5)
			wg.Done()
		}(i)
	}
	wg.Wait()

	//f, err := os.OpenFile(_path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	//if err != nil {
	//	panic(err)
	//}
	w := new(writerseeker.WriterSeeker)

	err = gif.EncodeAll(w, &g)
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

func init() {
	//root, _ := os.Getwd()
	zero.Default().OnRegex("^摸头").SetBlock(true).SetPriority(60).Handle(func(ctx *ZeroBot.Ctx) {
		//str := ctx.State["regex_matched"].([]string)[1]
		userId := ctx.Event.UserID
		for _, segment := range ctx.Event.Message {
			if segment.Type == "at"{
				userId = TypeUtils.StrToInt(segment.Data["qq"])
			}
		}
		//file := fmt.Sprintf("%d.jpg",userId)
		//_path := path.Join(root, PATH, file)

		faceImg := OicqUtils.GetQQFaceImg(userId)
		if faceImg == nil{
			return
		}
		t := time.Now()
		bs := touchHeadMake(faceImg)
		ctx.SendChain(zero.ImageBase64Message(bs), message.Text(fmt.Sprintf("%v",time.Since(t))))

	})

}
