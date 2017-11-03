package modules

import "github.com/gin-gonic/gin"

type Base struct {
	ResultInfo
}

func (self Base) UploadImage(ctx *gin.Context) {
	defer func() {
		JsonResult("login", ctx, self.ResultInfo)
	}()


}
