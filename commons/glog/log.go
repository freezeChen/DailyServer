package glog

import (
	"dailyserver2/commons/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

const TIMEFORMAT = "2006-01-02 15:04:05"

var gLogger *zap.Logger

func InitLogger() {

	path := util.GetCurrentDirectory()
	zapcore.AddSync(zap.CombineWriteSyncers())
	writerSyncer := zapcore.AddSync(&Logger{
		Filename:   path + "/log/log.txt",
		MaxSize:    20, // 单文件容量上限(MB)
		MaxBackups: 30, //
		MaxAge:     30, // 文件保存天数
		LocalTime:  true,
	})

	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	core := zapcore.NewCore(encoder, writerSyncer, zap.InfoLevel)
	gLogger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(3), zap.AddStacktrace(zapcore.ErrorLevel))
}

func Info(args ...interface{}) {
	defer gLogger.Sync()
	gLogger.Sugar().Info(args)
}
func APIInfo(path string, params interface{}) {
	defer gLogger.Sync()
	gLogger.Named("api").Info("操作成功", zap.String("path", path), zap.Any("params", params))
}

func APIWarn(path, errMsg string, params interface{}) {
	defer gLogger.Sync()
	gLogger.Named("api").Warn(errMsg, zap.String("path", path), zap.Any("params", params))
}

func Error(args ...interface{}) {
	defer gLogger.Sync()
	gLogger.Sugar().Error(args)
}

func Painc(args ...interface{}) {
	defer gLogger.Sync()
	gLogger.Sugar().Panic(args)

}

//日志时间格式化
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(TIMEFORMAT))
}
