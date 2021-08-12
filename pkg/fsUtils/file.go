package fsUtils

import "os"

func Getwd() string{
	f,_ := os.Getwd()
	return f
}
