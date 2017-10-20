package main

import (
	"github.com/kardianos/service"
	"log"
	"os"
	"DailySever/commons/config"

	log2 "DailySever/commons/log"
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

func (p *program) run() {
	config.InitConfig()
	log2.SetLog()
}