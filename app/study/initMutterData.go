package study

import (
	"fmt"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/fsUtils"
	"io/ioutil"
	"path"
)

func initMutterData(){

	data,err := ioutil.ReadFile(path.Join(fsUtils.Getwd(), "static", "sql", "chat_mutters.sql"))
	if err != nil{
		fmt.Println(err)
	}
	db.DB.Exec(string(data))

}
