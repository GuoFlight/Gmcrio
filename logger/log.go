package logger

import (
	"Gmicro/conf"
	"context"
	"github.com/GuoFlight/gerror/v2"
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

func getTraceId(ctx context.Context) string {
	v, ok := ctx.Value(conf.KeyTraceId).(string)
	if !ok {
		return ""
	}
	return v
}
func WithFieldTraceIdFromCtx(ctx context.Context) *logrus.Entry {
	traceId := getTraceId(ctx)
	return GLogger.WithField(conf.KeyTraceId, traceId)
}
func WithFieldTraceIdFromGerr(gerr *gerror.Gerr) *logrus.Entry {
	return GLogger.WithField(conf.KeyTraceId, gerr.TraceID)
}

func Info(ctx context.Context, arg ...any) {
	WithFieldTraceIdFromCtx(ctx).Info(arg)
}
func ErrWithCtx(ctx context.Context, msg string) *gerror.Gerr {
	return HandleGerr(gerror.SetSkip(2).SetTraceIdByCtx(ctx).NewErr(msg), nil)
}

func parseGerrExtInfo(gerr *gerror.Gerr, extInfo []map[string]interface{}) map[string]interface{} {
	var info map[string]interface{}
	if len(extInfo) > 0 && extInfo[0] != nil {
		info = extInfo[0]
	} else {
		info = make(map[string]interface{})
	}
	info["ErrFile"] = gerr.ErrFile
	info["ErrLine"] = gerr.ErrLine
	info["ErrFunc"] = gerr.ErrFunc
	return info
}

// HandleGerr 输出错误日志
func HandleGerr(gerr *gerror.Gerr, extInfo ...map[string]interface{}) *gerror.Gerr {
	info := parseGerrExtInfo(gerr, extInfo)
	WithFieldTraceIdFromGerr(gerr).WithFields(info).Error(gerr.Error())
	return gerr
}

// HandleGerrWarn 输出Warn日志
func HandleGerrWarn(gerr *gerror.Gerr, extInfo ...map[string]interface{}) *gerror.Gerr {
	info := parseGerrExtInfo(gerr, extInfo)
	WithFieldTraceIdFromGerr(gerr).WithFields(info).Warn(gerr.Error())
	return gerr
}
