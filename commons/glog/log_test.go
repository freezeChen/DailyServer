package glog

import (
	"testing"
	"fmt"
)

func TestInitLogger(t *testing.T) {
	InitLogger()

	defer func() {
		i := recover()
		fmt.Println(i)
	}()

	Painc("123")
}
func check(info error) {
	Painc(info)
}
