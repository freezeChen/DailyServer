package routers

import (
	"DailyServer/commons/middleware"
	"DailyServer/module_api/routers/api"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	_ "DailyServer/module_api/routers/api/docs"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	group := r.Group("user")
	{
		userController := api.UserController{}
		group.GET("/user", userController.GetUser)
		group.GET("/list", userController.GetUserList)
	}


	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.GET("swagger", func(context *gin.Context) {
	//
	//	context.HTML(http.StatusOK, "index.html", nil)
	//})
	return r
}
