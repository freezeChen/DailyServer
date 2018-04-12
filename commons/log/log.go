package log

import (
	. "DailyServer/commons/config"
	"DailyServer/commons/util"

	"github.com/aiwuTech/fileLogger"
	"fmt"
)

var LogFile *fileLogger.FileLogger

func SetLog() {
	rootDir := util.GetCurrentDirectory()
	fmt.Println(rootDir)
	name := Cfg.MustValue("log", "log_name", "")
	//name:="log.txt"
	path := Cfg.MustValue("log", "log_path", "")
fmt.Println(rootDir+path, name)
	LogFile = fileLogger.NewDailyLogger(rootDir+path, name, "", fileLogger.DEFAULT_LOG_SCAN, fileLogger.DEFAULT_LOG_SEQ)
	LogFile.SetLogConsole(true)
	LogFile.SetLogLevel(fileLogger.TRACE)
	fmt.Println("log~~~")
}
