package glog

import (
	"testing"
)

func TestInitLogger(t *testing.T) {
	InitLogger()

	gLogger.Warn("sfsffffffdf")
}
func check(info error) {
	Painc(info)
}
