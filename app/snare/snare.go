package snare

import (
	"fmt"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/fsUtils"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/zero"
	ZeroBot "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"io"
	"io/ioutil"
	"math/rand"
	"modernc.org/strutil"
	"net/http"
	"os"
	"path"
)

const PATH = "data/snare"

type Config struct{
	Enable 		bool
	RegexAdd 	string
	RegexExec	string
	RegexDel	string
	Priority	int
}

func Init(c Config) {

	if c.Enable == false{
		return
	}

	zero.Default().OnRegex(c.RegexExec).SetBlock(true).SetPriority(c.Priority).
		Handle(func(ctx *ZeroBot.Ctx) {
			groupId := ctx.Event.GroupID
			dir := fmt.Sprintf("%s/%v", PATH,groupId)
			rootDir, _ := os.Getwd()
			rd, err := ioutil.ReadDir(dir)
			if err != nil{ // 没有文件夹陷害不了
				return
			}
			picCount := len(rd)
			if picCount == 0{ // 没有图陷害不了
				return
			}
			picId := rand.Intn(picCount)

			url := fmt.Sprintf("file://%s/%s/%v",rootDir,dir,rd[picId].Name())
			//d, err := pathToBase64(url)
			//if err!= nil{
			//	return
			//}
			ctx.SendChain(message.Image(url))
	})

	zero.Default().OnRegex(c.RegexAdd, ZeroBot.OnlyGroup).SetBlock(true).SetPriority(c.Priority).
		Handle(func(ctx *ZeroBot.Ctx) {
			groupId := ctx.Event.GroupID
			for _, elem := range ctx.Event.Message {
				fmt.Println(elem.Data)
				if elem.Type == "image" {
					dir := fmt.Sprintf("%s/%v", PATH,groupId)
					_, ext := os.Stat(dir)
					if os.IsNotExist(ext){
						err := os.MkdirAll(dir, os.ModePerm)
						if err!=nil{
							fmt.Println(err)
						}
					}

					file := fmt.Sprintf("%s/%v/%s", PATH,groupId,elem.Data["file"])
					_, ext = os.Stat(file)
					if os.IsExist(ext){
						ctx.Send("已经存在这张图了")
						continue
					}
					err := picDownload(ctx,file,elem.Data["url"])
					if err!=nil{
						fmt.Println(err)
					}
					ctx.SendChain(message.Text("安排上了"),
						message.At(ctx.Event.UserID),
						message.Text("的图"),
						message.Image(fmt.Sprintf("file:///%s",path.Join(fsUtils.Getwd(), file))))
				}
			}
	})

	zero.Default().OnRegex(c.RegexDel, ZeroBot.OnlyGroup).SetBlock(true).SetPriority(c.Priority).
		Handle(func(ctx *ZeroBot.Ctx) {
		groupId := ctx.Event.GroupID
		for _, elem := range ctx.Event.Message {
			fmt.Println(elem.Data)
			if elem.Type == "image" {
				file := fmt.Sprintf("%s/%v/%s", PATH,groupId,elem.Data["file"])
				_, ext := os.Stat(file)
				if os.IsNotExist(ext){
					ctx.Send("没有这张图哦")
				} else {
					if os.Remove(file) != nil{
						ctx.Send("删除失败")
					}else{
						ctx.Send("删除成功")
					}
				}
				break
			}
		}
	})
	
}

func picDownload(ctx *ZeroBot.Ctx,file string, url string)(err error){
	f,err := os.Create(file)
	if err!= nil{
		return
	}
	defer f.Close()
	res, err := http.Get(url)
	if err!= nil{
		return
	}
	io.Copy(f, res.Body)
	return
}

func pathToBase64(path string)(b string, err error){
	d,err := ioutil.ReadFile(path)
	if err  != nil{
		//fmt.Println(err)
		return
	}
	b = fmt.Sprintf("base64:/%s==", strutil.Base64Encode(d))
	return
}
