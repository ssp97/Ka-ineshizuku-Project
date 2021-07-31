package avoidExamine

import (
	"math/rand"
	"os"
)

func randomBytes(size int)[]byte{
	r := make([]byte, size)
	rand.Read(r)
	return r
}

func PicFile(path string)(err error){
	f,err :=os.OpenFile(path, os.O_WRONLY, 0644)
	if err!= nil{
		return
	}
	defer f.Close()
	n, _ := f.Seek(0, os.SEEK_END)
	_, err = f.WriteAt(randomBytes(rand.Intn(20)),n)
	if err != nil{
		return
	}

	return
}
