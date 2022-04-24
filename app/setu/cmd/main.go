package main

import "github.com/ssp97/Ka-ineshizuku-Project/app/setu"

func main() {
	c := new(setu.Config)
	c.PixivUser = "xiaonabot"
	c.PixivPassword = "xn123456."

	pixivapi := new(setu.SetuPixivApi)
	pixivapi.Login()
	pixivapi.GetUserAllPic(5598737)

}

