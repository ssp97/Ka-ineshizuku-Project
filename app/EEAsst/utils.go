package EEAsst

import "fmt"

func autoUnit(value float64)(string){
	if value < 1 {
		return fmt.Sprintf("%.2fm", value*1000)
	}else if value < 1000 {
		return fmt.Sprintf("%.2f", value)
	}else if value < 1000000 {
		return fmt.Sprintf("%.2fk", value/1000)
	}else {
		return fmt.Sprintf("%.2fM", value/1000000)
	}
}