package logging

import (
	"fmt"
	"github.com/chenwbyx/leafserver/server/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	AppLogger *zap.Logger
	HTTPLogger *zap.Logger
	DefaultLogger *zap.Logger
	LoginLogger *zap.Logger
)

const (
	MAXSIZE = 8
	MAXBACKUPS = 2
	MAXAGE = 7
)

//定制日志
func Setup() {
	//定制日志
	AppLogger = NewLogger(setting.LogSetting.App, zapcore.InfoLevel, MAXSIZE, MAXBACKUPS, MAXAGE, true, setting.LogSetting.ServiceName)
	HTTPLogger = NewLogger(setting.LogSetting.Http, zapcore.InfoLevel, MAXSIZE, MAXBACKUPS, MAXAGE,true, setting.LogSetting.ServiceName)
	DefaultLogger = NewLogger(setting.LogSetting.Default, zapcore.InfoLevel, MAXSIZE, MAXBACKUPS, MAXAGE,true, setting.LogSetting.ServiceName)
	LoginLogger = NewLogger(setting.LogSetting.Login, zapcore.InfoLevel, MAXSIZE, MAXBACKUPS, MAXAGE,true, setting.LogSetting.ServiceName)
}

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}