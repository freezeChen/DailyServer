package log

import (
	. "DailyServer/commons/config"
	"DailyServer/commons/util"

	"github.com/aiwuTech/fileLogger"
)

var LogFile *fileLogger.FileLogger

func SetLog() {
	rootDir := util.GetCurrentDirectory()
	name := Cfg.MustValue("log", "log_name", "")
	path := Cfg.MustValue("log", "log_path", "")
	LogFile = fileLogger.NewDailyLogger(rootDir+path, name, "", fileLogger.DEFAULT_LOG_SCAN, fileLogger.DEFAULT_LOG_SEQ)
	LogFile.SetLogConsole(true)
	LogFile.SetLogLevel(fileLogger.TRACE)

}
