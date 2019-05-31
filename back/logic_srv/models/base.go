package models

import (
	"DailyServer/back/commons/db"
	"DailyServer/back/commons/glog"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
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
	Data interface{} `json:"data,omitempty"`
}

func (self *JsonResult) Response(ctx *gin.Context) {
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
	if self.Code == 0 {
		self.Msg = "success"
		glog.APIInfo(ctx.Request.URL.Path, params)
	} else {
		self.Msg = "操作失败"
		glog.APIWarn(ctx.Request.URL.Path, self.Msg, params)
	}

	json := jsoniter.ConfigCompatibleWithStandardLibrary

	str, _ := json.Marshal(self)

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", str)
}

func Engine() *xorm.Engine {
	group, _ := db.NewEngine()
	return group
}
