package study

import (
	"fmt"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/fsUtils"
	"io/ioutil"
	"path"
	"strings"
)

func initMutterData(){

	data,err := ioutil.ReadFile(path.Join(fsUtils.Getwd(), "static", "sql", "chat_mutters.sql"))
	if err != nil{
		fmt.Println(err)
	}
	sqlArr:=strings.Split(string(data),";")
	for _,sql:=range sqlArr{
		if sql==""{
			continue
		}
		db.DB.Exec(sql)
	}

}
