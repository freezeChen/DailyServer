/*
    @Time : 2019-03-25 09:51 
    @Author : frozenchen
    @File : accountController
    @Software: DailyServer
*/
package controller

import (
	"DailyServer/back/commons/db"
	models2 "DailyServer/back/logic_srv/models"
	"github.com/gin-gonic/gin"
)

type userController struct {
	models2.JsonResult
}

func NewUserController() *userController {
	var controller = &userController{}
	controller.Code = -1
	return controller
}

func (self userController) Router(router *gin.Engine) {
	group := router.Group("/web/account/")
	group.GET("signin", func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", gin.H{})
	})

	self.api(router)
}

func (self userController) api(router *gin.Engine) {
	group := router.Group("web/api/account/")

	group.POST("login", self.login)
}

func (self userController) login(ctx *gin.Context) {
	defer func() {
		self.Response(ctx)
	}()

	var param = struct {
		Account  string `binding:"required"`
		Password string `binding:"required"`
	}{}

	if err := ctx.ShouldBind(&param); err != nil {
		self.Msg = "参数校验失败"
		return
	}

	var user models2.User
	user.Account = param.Account

	has, err := db.Engine().Get(&user)
	if err != nil {
		self.Msg = err.Error()
		return
	}

	if !has {
		self.Msg = "账号不存在"
		return
	}

	if user.Password != param.Password {
		self.Msg = "密码错误"
		return
	}

	self.Code = 0
	self.Data = user

}
