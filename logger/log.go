package logger

import (
	"Gmicro/conf"
	"context"
	"github.com/GuoFlight/gerror"
	"github.com/GuoFlight/glog"
	"github.com/sirupsen/logrus"
	"log"
)

var (
	GLogger *logrus.Logger
)

func InitLog() {
	path := conf.GConf.Log.Path
	logLevel := conf.GConf.Log.Level
	var err error
	GLogger, err = glog.NewLogger(path, logLevel, false, 10)
	if err != nil {
		log.Fatal("日志初始化失败:", err)
	}
	GLogger.Info("日志初始化完成")
}

// PrintInfo 输出info日志
func PrintInfo(ctx context.Context, msg ...interface{}) {
	traceId, _ := ctx.Value(conf.TraceIdName).(string)
	GLogger.WithFields(logrus.Fields{conf.TraceIdName: traceId}).Info(msg)
}

// PrintErr 输出错误日志
func PrintErr(gerr *gerror.Gerr, elseInfo map[string]interface{}, msg ...interface{}) *gerror.Gerr {
	if len(elseInfo) == 0 {
		elseInfo = make(map[string]interface{})
	}
	elseInfo["ErrFile"] = gerr.ErrFile
	elseInfo["ErrLine"] = gerr.ErrLine
	elseInfo[conf.TraceIdName] = gerr.TraceID
	GLogger.WithFields(elseInfo).Error(gerr.Error(), msg)
	return gerr
}

// PrintWarn 输出Warn日志
func PrintWarn(ctx context.Context, elseInfo map[string]interface{}, msg ...interface{}) {
	traceId, _ := ctx.Value(conf.TraceIdName).(string)
	if len(elseInfo) == 0 {
		elseInfo = make(map[string]interface{})
	}
	elseInfo[conf.TraceIdName] = traceId
	GLogger.WithFields(elseInfo).Warn(msg)
}
