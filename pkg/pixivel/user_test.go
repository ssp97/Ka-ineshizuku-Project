package pixivel

import (
	"fmt"
	"testing"
)

func TestGetUserAllIllust(t *testing.T) {
	data := GetUserAllIllust(1226647)
	fmt.Println(data)
}