package api

import (
	"Gmicro/conf"
	"Gmicro/myctx"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

// 得到日志切割的输出对象
func getWriter(pathLog string, rotationCount uint) (*rotatelogs.RotateLogs, error) {
	return rotatelogs.New(
		pathLog+".%Y%m%d%H",                                     // 日志文件后缀：年月日时
		rotatelogs.WithLinkName(pathLog),                        // 为当前正在输出的日志文件建立软连接
		rotatelogs.WithRotationCount(rotationCount),             // 日志文件保存的个数(包括当前正在输出的日志)
		rotatelogs.WithRotationTime(time.Duration(1)*time.Hour), // 设置日志分割的时间(隔多久分割一次)
	)
}

// 日志中间件
func logMiddleware() gin.HandlerFunc {
	ginLog := logrus.New()

	// 日志不输出到终端
	nullFile, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal("打开null文件失败:", err)
	}
	ginLog.SetOutput(nullFile)

	// 设置日志切割
	logWriter, err := getWriter(conf.GConf.Http.LogFilePath, conf.GConf.Http.LogFileCount)
	if err != nil {
		log.Fatal(err)
	}
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	ginLog.AddHook(lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))
	// 返回一个函数，每次请求时都会运行此函数
	return func(c *gin.Context) {
		// 获取请求耗时
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// 打印日志
		ginLog.WithFields(logrus.Fields{
			"client_ip":   c.ClientIP(),
			"method":      c.Request.Method,
			"query_args":  c.Request.URL.RawQuery,
			"latency":     latency,
			"url":         c.Request.URL.Path,
			"status_code": c.Writer.Status(),
			"trace_id":    myctx.GCtx.Gin.GetTraceId(c),
			"op_user":     myctx.GCtx.Gin.GetOpUser(c),
		}).Info()
	}
}
