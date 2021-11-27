package logger

import (
	"Gmicro/conf"
	"log"
	"github.com/GuoFlight/glog"
	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger
)

func InitLog() {
	path := conf.GlobalConfig.Log.Path
	logLevel := conf.GlobalConfig.Log.Level
	var err error
	Logger,err = glog.NewLogger(path,logLevel,false,10)
	if err!=nil{
		log.Fatal("日志初始化失败:",err)
	}
}
