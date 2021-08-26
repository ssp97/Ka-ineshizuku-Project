package avoidExamine

import (
	"bytes"
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

func PicByte(data ...[]byte)(result []byte){
	r := randomBytes(rand.Intn(20))
	result = bytes.Join(data,r)
	return
}

//func PicBase64(data string)(result string){
//	decodeBytes, err := base64.StdEncoding.DecodeString(data)
//	r := randomBytes(rand.Intn(20))
//	data :=
//	rs := base64.StdEncoding.EncodeToString(r)
//}