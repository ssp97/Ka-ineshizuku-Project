package EEAsst

import "errors"

var eiaSizeDict = map[string]string{
	"008004":"0.25 x 0.125 mm",
	"01005":"0.4 x 0.2 mm",
	"015015":"0.38 x 0.38 mm",
	"015008":"0.5 x 0.25 mm",
	"0201":"0.6 x 0.3 mm",
	"0202":"0.5 x 0.5 mm",
	"0303":"0.8 x 0.8 mm",
	"02404":"0.6 x 1.0 mm",
	"0402":"1.0 x 0.5 mm",
	"0603":"1.6 x 0.8 mm",
	"0704":"1.8 x 1.0 mm",
	"0805":"2.0 x 1.25 mm",
	"1111":"2.8 x 2.8 mm",
	"1206":"3.2 x 1.6 mm",
	"1210":"3.2 x 2.5 mm",
	"1808":"4.5 x 2.0 mm",
	"1812":"4.5 x 3.2 mm",
	"2211":"5.7 x 2.8 mm",
	"2220":"5.7 x 5.0 mm",
}

func getEIASize(str string)(string, error){
	value, ok := eiaSizeDict[str]
	if ok{
		return value, nil
	}
	return "", errors.New("Unknown")
}
