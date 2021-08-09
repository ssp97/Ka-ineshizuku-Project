package OicqUtils

import (
	"fmt"
	"image"
	"net/http"
)

//

func GetQQFaceImg(qqId int64)(img *image.Image){
	url := fmt.Sprintf("https://api.sumt.cn/api/qq.logo.php?qq=%d", qqId)
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
