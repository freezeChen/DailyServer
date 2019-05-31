/*
    @Time : 2018-12-19 11:58 
    @Author : frozenchen
    @File : conf
    @Software: 24on
*/

package zconf

import (
	"fmt"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source"
	"github.com/micro/go-config/source/env"
	"github.com/micro/go-config/source/file"
	"os"
	"path/filepath"
)

type conf struct {
	Conf struct {
		DB struct {
			Driver     string `json:"driver"`
			SourceName string `json:"sourceName"`
		} `json:"db"`
		Redis struct {
			IP   string `json:"ip"`
			Auth string `json:"auth"`
		} `json:"redis"`
		Debug string `json:"debug"`
		Kafka string `json:"kafka"`
	} `json:"conf"`
}

var myConf *conf

func GetEnv() *conf {
	return myConf
}

func InitConfig() error {
	var sources []source.Source

	myConf = new(conf)
	conf := config.NewConfig()

	//加载环境变量

	envSource := env.NewSource(env.WithPrefix("conf"))
	sources = append(sources, envSource)

	//加载文件变量
	if fileSource := loadFileSource(); fileSource != nil {
		sources = append(sources, fileSource)
	}

	err := conf.Load(sources...)
	if err != nil {

		return err
	}
	err = conf.Scan(myConf)
	if err != nil {
		return err
	}
	return nil
}

func loadFileSource() source.Source {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil
	}
	fmt.Println(dir)

	_, err = os.Stat(dir + "/conf/conf.yaml")
	if err != nil {
		fmt.Println("os.stat", err)
		return nil
	}
	fileSource := file.NewSource(file.WithPath(dir + "/conf/conf.yaml"))

	return fileSource
}
