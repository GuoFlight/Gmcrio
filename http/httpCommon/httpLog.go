package httpCommon

import (
	"Gmicro/conf"
	"Gmicro/logger"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"os"
)

var HttpLogger *logrus.Logger

//日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	//初始化http日志
	HttpLogger = logrus.New()

	//设置日志级别
	if conf.GlobalConfig.Log.Level == "TRACE" {
		HttpLogger.SetLevel(logrus.TraceLevel)
	} else if conf.GlobalConfig.Log.Level == "DEBUG" {
		HttpLogger.SetLevel(logrus.DebugLevel)
	} else if conf.GlobalConfig.Log.Level == "WARN" {
		HttpLogger.SetLevel(logrus.WarnLevel)
	} else if conf.GlobalConfig.Log.Level == "ERROR" {
		HttpLogger.SetLevel(logrus.ErrorLevel)
	} else if conf.GlobalConfig.Log.Level == "FATAL" {
		HttpLogger.SetLevel(logrus.FatalLevel)
	} else if conf.GlobalConfig.Log.Level == "PANIC" {
		HttpLogger.SetLevel(logrus.PanicLevel)
	} else {
		HttpLogger.SetLevel(logrus.InfoLevel) //默认设置为INFO级别
	}

	//全部日志输出行号
	//Logger.SetReportCaller(true)
	//指定日志等级输出行号
	var printFileAndNumHook logger.PrintFileAndNumHook
	HttpLogger.AddHook(&printFileAndNumHook)

	//日志不输出到终端
	if conf.GlobalConfig.Log.Level!="DEBUG"{
		nullFile, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err == nil{
			HttpLogger.SetOutput(nullFile) //将日志输出到此文件
		}else{
			log.Panic("打开null文件失败")
		}
	}

	//多文件输出+日志切割
	httpLogDir := fmt.Sprintf("%s/http",conf.GlobalConfig.Log.Path)	//http日志子目录
	pathMap := lfshook.WriterMap{
		logrus.TraceLevel: logger.GetWriter(fmt.Sprintf("%s/TRACE.log",httpLogDir)),
		logrus.DebugLevel: logger.GetWriter(fmt.Sprintf("%s/DEBUG.log",httpLogDir)),
		logrus.InfoLevel:  logger.GetWriter(fmt.Sprintf("%s/INFO.log",httpLogDir)),
		logrus.WarnLevel: logger.GetWriter(fmt.Sprintf("%s/WARN.log",httpLogDir)),
		logrus.ErrorLevel: logger.GetWriter(fmt.Sprintf("%s/ERROR.log",httpLogDir)),
		logrus.FatalLevel: logger.GetWriter(fmt.Sprintf("%s/FATAL.log",httpLogDir)),
		logrus.PanicLevel: logger.GetWriter(fmt.Sprintf("%s/PANIC.log",httpLogDir)),
	}
	HttpLogger.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))

	//每次请求时都会运行此函数，用于记录请求日志
	return func(c *gin.Context) {
		//获取body
		reqBody,err:=ioutil.ReadAll(c.Request.Body)
		if err!=nil{
			HttpLogger.WithFields(logrus.Fields{
				"Method":c.Request.Method,
				"Url":c.Request.RequestURI,
				"clientAddress":c.Request.RemoteAddr,
			}).Error("获取body失败")
			return
		}
		if len(reqBody)==0{
			HttpLogger.WithFields(logrus.Fields{
				"Method":c.Request.Method,
				"Url":c.Request.RequestURI,
				"clientAddress":c.Request.RemoteAddr,
			}).Info()
		}else{
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
			if err==nil{	//未加密的请求
				HttpLogger.WithFields(logrus.Fields{
					"Method":c.Request.Method,
					"Url":c.Request.RequestURI,
					"clientAddress":c.Request.RemoteAddr,
					"Body":string(reqBody),
				}).Info()
			}else{			//加密请求
				HttpLogger.WithFields(logrus.Fields{
					"Method":c.Request.Method,
					"Url":c.Request.RequestURI,
					"clientAddress":c.Request.RemoteAddr,
					"Body":string(reqBody),
				}).Info()
				return
			}
		}
	}
}

