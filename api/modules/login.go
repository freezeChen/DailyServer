package modules

import "github.com/gin-gonic/gin"

type Login struct {
	ResultInfo
}

func (self Login) Login(ctx *gin.Context) {
	defer func() {
		JsonResult("login", ctx, self.ResultInfo)
	}()
	account := ctx.Query("account")

	self.code = 0
	self.data = gin.H{
		"account": account,
	}
}
