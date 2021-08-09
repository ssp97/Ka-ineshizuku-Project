package TypeUtils

import "strconv"

func StrToInt(str string) int64 {
	val, _ := strconv.ParseInt(str, 10, 64)
	return val
}
