/*
    @Time : 2018-12-19 13:35 
    @Author : frozenchen
    @File : config_test.go
    @Software: 24on
*/
package zconf

import (
	"fmt"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/env"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	var (
		myconf = new(conf)
	)
	os.Setenv("conf_kafka","1123:123")
	conf := config.NewConfig()
	source := env.NewSource(env.WithPrefix("conf"))

	err := conf.Load(source)
	if err != nil {
		t.Error(err)
	}

	err = conf.Scan(myconf)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v", myconf)
}
