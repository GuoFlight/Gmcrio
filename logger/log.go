package logger

import (
	"Gmicro/conf"
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
