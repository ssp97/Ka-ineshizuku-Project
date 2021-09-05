package gocc

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

var goccS2t *OpenCC

func init() {
	_goccs2t, err := New("s2t")
	if err != nil{
		log.Errorf("gocc err %s", err)
	}else{
		goccS2t = _goccs2t
	}
	testStr := "今天是个好日子"
	fmt.Printf("gocc test %s -> %s\r\n",testStr, S2t(testStr))
}

func S2t(in string)(out string){
	out, err := goccS2t.Convert(in)
	if err!=nil{
		log.Errorf("gocc err %s", err)
		return ""
	}
	return out
}
