package gifApp

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"

	//"github.com/golang/freetype/raster"
	"os"
)

func ttfFontLoad(path string)(*truetype.Font,error){
	f,err := os.ReadFile(path)
	if err != nil{
		return nil,err
	}
	font, err := freetype.ParseFont(f)
	if err != nil {
		return nil,err
	}
	return font,err
}

func otfFontLoad(path string)(*sfnt.Font, error){
	f,err := os.ReadFile(path)
	if err != nil{
		return nil,err
	}
	font, err := opentype.Parse(f)
	if err != nil {
		return nil,err
	}
	return font,err
}

