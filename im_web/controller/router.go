/*
    @Time : 2019-03-25 10:25 
    @Author : frozenchen
    @File : router
    @Software: DailyServer
*/
package controller

import (
	"DailyServer/commons/middleware"
	"github.com/gin-gonic/gin"
)

func Start() {
	engine := gin.Default()
	engine.Use(middleware.CORSMiddleware())
	NewUserController().Router(engine)

	engine.Run(":8021")
}
