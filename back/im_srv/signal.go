package im

import (
	"os"
	"os/signal"
	"syscall"


)

func InitSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			return
		case syscall.SIGHUP:
			reload()
		default:
			return
		}
	}
}

func reload() {
	//newConf, err := ReloadConfig()
	//if err != nil {
	//	log.Error("ReloadConfig() error(%v)", err)
	//	return
	//}
	//Conf = newConf
}
