package study

import (
	"fmt"
	"github.com/ssp97/Ka-ineshizuku-Project/pkg/fsUtils"
	"io/ioutil"
	"path"
)

func initStudyData() {
	data,err := ioutil.ReadFile(path.Join(fsUtils.Getwd(), "static", "sql", "chat_studies.sql"))
	if err != nil{
		fmt.Println(err)
	}
	db.DB.Exec(string(data))
}
