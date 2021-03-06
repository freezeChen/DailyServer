package main

import (
	"github.com/kardianos/service"
	"os"
	. "DailyServer/commons/config"
	. "DailyServer/commons/log"
	"DailyServer/commons/mysql"
	"github.com/gin-gonic/gin"
	"DailyServer/api/modules"
	"log"
	"time"
	"DailyServer/commons/socketmanager"
)

type program struct{}

const (
	SERVICE_NAME        string = "Daily api service"
	SERVICE_DISPLAYNAME string = "Daily api service"
	SERVICE_DESCRIPTION string = "using for daily"
)

func main() {
	cfg := &service.Config{
		Name:        SERVICE_NAME,
		DisplayName: SERVICE_DISPLAYNAME,
		Description: SERVICE_DESCRIPTION,
	}
	prg := &program{}
	ser, err := service.New(prg, cfg)
	if err != nil {
		log.Fatalf("Failed to new: %ser", err)
	}
	if len(os.Args) == 2 {
		method := os.Args[1]
		if err := service.Control(ser, method); err != nil {
			log.Fatalf("Failed to %ser: %ser\n", method, err)
		}
		log.Printf("Service \"%ser\" %ser.\n", SERVICE_DISPLAYNAME, method)
		return
	}

	if err := ser.Run(); err != nil {
		log.Fatalf("Failed to run: %ser\n", err)
	}
}
func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	InitConfig()
	SetLog()

	mysql.Refresh()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(CORSMiddleware())
	g0 := r.Group("/daily/account")
	{
		account := new(modules.Account)
		account.Code = -1
		g0.GET("login", account.Login)
	}

	r.GET("/daily/socket", socketmanager.HandelHttpToWebSocket)

	r.Run(":8088")

}

func (p *program) Stop(s service.Service) error {
	<-time.After(time.Second * 13)
	return nil
}
func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token, X-File-Name")
		ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
