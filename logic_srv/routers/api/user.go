package api

import (
	"DailyServer/commons/util"
	"DailyServer/logic_srv/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	models.JsonResult
}

// @Summary 获取用户
// @Produce  json
// @Param id query int true "ID"
// @Success 200 {string} json "{"code":200,"data":{"id":3,"created_on":1516937037,"modified_on":0,"tag_id":11,"tag":{"id":11,"created_on":1516851591,"modified_on":0,"name":"312321","created_by":"4555","modified_by":"","state":1},"content":"5555","created_by":"2412","modified_by":"","state":1},"msg":"ok"}"
// @Router /user/user [get]
func (self UserController) GetUser(c *gin.Context) {
	defer func() {
		models.Response(&self.JsonResult, c)
	}()

	id := c.Query("id")
	user, err := models.GetUserByID(util.ToInt64(id))
	if err != nil {
		self.Msg = fmt.Sprintf("Faild to:%s", err)
		return
	}
	self.Data = user
	self.Code = 0
}

// @Summary 获取用户列表
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"id":3,"created_on":1516937037,"modified_on":0,"tag_id":11,"tag":{"id":11,"created_on":1516851591,"modified_on":0,"name":"312321","created_by":"4555","modified_by":"","state":1},"content":"5555","created_by":"2412","modified_by":"","state":1},"msg":"ok"}"
// @Router /user/list [get]
func (self UserController) GetUserList(c *gin.Context) {
	defer func() {
		models.Response(&self.JsonResult, c)
	}()
	self.Code = 500
	list, err := models.GetUserList()
	if err != nil {
		self.Msg = fmt.Sprintf("Faild to:%s", err)
		return
	}
	self.Code = 0
	self.Data = list
}
