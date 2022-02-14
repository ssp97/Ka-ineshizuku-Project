package avoidExamine

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	_ "image/jpeg"
	_ "image/gif"
	_ "golang.org/x/image/bmp"
	"math/rand"
	"os"
)

func randomBytes(size int)[]byte{
	r := make([]byte, size)
	rand.Read(r)
	return r
}

func PicFile(path string)(err error){
	f,err :=os.OpenFile(path, os.O_WRONLY, 0644)
	if err!= nil{
		return
	}
	defer f.Close()
	n, _ := f.Seek(0, os.SEEK_END)
	_, err = f.WriteAt(randomBytes(rand.Intn(20)),n)
	if err != nil{
		return
	}

	return
}

func PicByte(data ...[]byte)(result []byte){
	r := randomBytes(rand.Intn(20))
	result = bytes.Join(data,r)
	return
}

func PicRandomDot(data ...[]byte)(result []byte){
	reader := bytes.NewReader(data[0])
	buffer := new(bytes.Buffer)
	img, formatName, err := image.Decode(reader)
	if err != nil {
		log.Warn(err, formatName)
		log.Info(string(data[0]))
		return nil
	}
	dx := img.Bounds().Dx()
	dy := img.Bounds().Dy()
	x := rand.Intn(dx)
	y := rand.Intn(dy)
	bounds := image.NewRGBA(img.Bounds())
	draw.Draw(bounds, img.Bounds(), img, img.Bounds().Min, draw.Src)

	bounds.SetRGBA(x,y, color.RGBA{A: 255, R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255))})
	bounds.SetRGBA(x,y, color.RGBA{A: 255, R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255))})
	bounds.SetRGBA(x,y, color.RGBA{A: 255, R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255))})

	err = png.Encode(buffer, bounds)
	result = buffer.Bytes()

	return
}

//func PicBase64(data string)(result string){
//	decodeBytes, err := base64.StdEncoding.DecodeString(data)
//	r := randomBytes(rand.Intn(20))
//	data :=
//	rs := base64.StdEncoding.EncodeToString(r)
//}