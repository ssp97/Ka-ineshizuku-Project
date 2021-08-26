package fsUtils

import (
	"encoding/base64"
	"io/ioutil"
	"os"
)

func Getwd() string{
	f,_ := os.Getwd()
	return f
}

func ReadBase64(path string)(string, error){
	d, err := ioutil.ReadFile(path)
	if err != nil{
		return "", err
	}
	base64str := base64.StdEncoding.EncodeToString(d)
	return base64str,nil
}

func ReadFile(path string)(result []byte){
	d, err := ioutil.ReadFile(path)
	if err!= nil{
		return
	}
	result = d
	return
}