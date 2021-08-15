package jieba

import (
	"github.com/yanyiwu/gojieba"
	"os"
	"path"
	"path/filepath"
)

var dictDir = path.Join(filepath.Dir(os.Args[0]), path.Join("static", "dict"))
var jiebaPath = path.Join(dictDir, "jieba.dict.utf8")
var hmmPath = path.Join(dictDir, "hmm_model.utf8")
var userPath = path.Join(dictDir, "user.dict.utf8")
var idfPath = path.Join(dictDir, "idf.utf8")
var stopPath = path.Join(dictDir, "stop_words.utf8")
var Seg = gojieba.NewJieba(jiebaPath, hmmPath, userPath, idfPath, stopPath)

