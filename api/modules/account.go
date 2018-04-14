package modules

import (
	"github.com/gin-gonic/gin"

	"DailyServer/api/models"
	"fmt"
)

type Account struct {
	ResultInfo
}

func (self Account) Login(ctx *gin.Context) {
	defer func() {
		JsonResult("login", ctx, self.ResultInfo)
	}()
	account := ctx.PostForm("FAccount")

	psw := ctx.PostForm("FPassword")

	if len(account) == 0 {
		self.Msg = "empty Account"
		return
	}
	if len(psw) == 0 {
		self.Msg = "empty psw"
		return
	}

	user := &models.User{Faccount: account}

	has, err := models.Engine().Get(user)
	if err != nil {
		self.Msg = fmt.Sprintf("Faild to select user:%s", err)
		return
	}

	if has {
		if user.Fpassword == psw {
			self.Code = 0
		} else {
			self.Code = 10001
		}
		return
	}

	self.Code = 10002

}

func (self Account) Register(c *gin.Context) {
	defer func() {
		JsonResult("register", c, self.ResultInfo)
	}()

	account := c.PostForm("FAccount")
	psw := c.PostForm("FPassword")
	name := c.PostForm("FName")

	if len(account) == 0 {
		self.Msg = "empty Account"
		return
	}
	if len(psw) == 0 {
		self.Msg = "empty psw"
		return
	}
	if len(name) == 0 {
		self.Msg = "empty name"
		return
	}
	user := &models.User{}

	has, err := models.Engine().Where("FAccount=?", account).Get(user)
	if err !=nil {
	self.Msg = fmt.Sprintf("Faild to selec user:%s",err)
	return
	}
	if has {
		self.Code =10003
	}


}
