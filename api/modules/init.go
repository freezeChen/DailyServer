package modules

import (
	"github.com/gin-gonic/gin"
	"DailyServer/commons/util"
	"DailyServer/commons/config"
	. "DailyServer/commons/log"
	"fmt"
)

type ResultInfo struct {
	Code int64	`json:"code"`
	Msg  string	`json:"msg"`
	Data interface{} `json:"data"`
}

func JsonResult(method string, ctx *gin.Context, ref ResultInfo) {

	errCode := util.ToString(ref.Code)
	errMsg := config.Cfg.MustValue("error", "msg_"+errCode, "")
	h := gin.H{"errcode": errCode, "errmsg": errMsg}
	if ref.Data != nil {
		h["Data"] = ref.Data
	}
	var getParam string
	req := ctx.Request
	for k, v := range req.URL.Query() {
		getParam = getParam + fmt.Sprintf("%s:%s,", k, v[0])
	}
	if len(getParam) > 0 {
		LogFile.I("传入参数(get)", fmt.Sprintf("[%s] "+getParam[0:len(getParam)-1], method))
	} else {
		var postParam string
		if req.Form == nil {
			req.ParseMultipartForm(32 << 20)
		}
		for k1, v1 := range req.Form {
			postParam = postParam + fmt.Sprintf("%s:%s,", k1, v1[0])
		}
		if len(postParam) > 0 {
			LogFile.I("传入参数(post)", fmt.Sprintf("[%s] "+postParam[0:len(postParam)-1], method))
		}
	}
	if len(ref.Msg) > 0 {
		LogFile.I("错误提示", fmt.Sprintf("[%s] "+ref.Msg, method))
	}
	ctx.JSON(200, h)

	if ref.Code != 0 {
		ctx.Abort()
	}
}
