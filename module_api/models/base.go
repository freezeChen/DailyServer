package models

import (
	"dailyserver2/commons/config"
	"dailyserver2/commons/db"
	"dailyserver2/commons/glog"
	"dailyserver2/commons/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/json-iterator/go"
	"net/http"
)

var (
	NotExistError = errors.New("don't exist")
)

const (
	TIMEFORMAT   string = "2006-01-02 15:04:05"
	TIMEFORMAT_2 string = "2006-01-02"
	TIMEFORMAT_3 string = "2006-01-02 15:04"
	TIMEFORMAT_4 string = "2006-01"
	TIMEFORMAT_5 string = "200601"
	TIMEFORMAT_6 string = "20060102"
)

type JsonResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}


func Response(self *JsonResult, ctx *gin.Context) {
	var (
		params interface{}
	)

	//获取传入参数
	if ctx.Request.Method == "GET" {
		params = ctx.Request.URL.Query()
	} else {
		ctx.Request.ParseMultipartForm(32 << 20)
		params = ctx.Request.Form
	}

	//获取错误码对应文字
	msg := config.Cfg.MustValue("error", "msg_"+util.ToString(self.Code), "")
	if self.Code == 0 {
		glog.APIInfo(ctx.Request.URL.Path, params)
	} else {
		glog.APIWarn(ctx.Request.URL.Path, self.Msg, params)
		fmt.Println("错误:", ctx.Request.URL.Path, self.Msg, params)
	}

	self.Msg = msg

	json := jsoniter.ConfigCompatibleWithStandardLibrary

	str, _ := json.Marshal(self)

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", str)
	//ctx.JSON(http.StatusOK, str)
}

func engine() *xorm.EngineGroup {
	group, _ := db.NewEngine()
	return group
}
