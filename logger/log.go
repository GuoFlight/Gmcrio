package logger

import (
	"Gmicro/conf"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger
)

func InitLog() {
	Logger = logrus.New()

	//设置日志级别
	if conf.GlobalConfig.Log.Level == "TRACE" {
		Logger.SetLevel(logrus.TraceLevel)
	} else if conf.GlobalConfig.Log.Level == "DEBUG" {
		Logger.SetLevel(logrus.DebugLevel)
	} else if conf.GlobalConfig.Log.Level == "WARN" {
		Logger.SetLevel(logrus.WarnLevel)
	} else if conf.GlobalConfig.Log.Level == "ERROR" {
		Logger.SetLevel(logrus.ErrorLevel)
	} else if conf.GlobalConfig.Log.Level == "FATAL" {
		Logger.SetLevel(logrus.FatalLevel)
	} else if conf.GlobalConfig.Log.Level == "PANIC" {
		Logger.SetLevel(logrus.PanicLevel)
	} else {
		Logger.SetLevel(logrus.InfoLevel) //默认设置为INFO级别
	}

	//全部日志输出行号
	//Logger.SetReportCaller(true)
	//指定日志等级输出行号
	var printFileAndNumHook PrintFileAndNumHook
	Logger.AddHook(&printFileAndNumHook)

	//多文件输出+日志切割
	logDir := fmt.Sprintf("%s/server",conf.GlobalConfig.Log.Path)	//http日志子目录
	pathMap := lfshook.WriterMap{
		logrus.TraceLevel: GetWriter(fmt.Sprintf("%s/TRACE.log",logDir)),
		logrus.DebugLevel: GetWriter(fmt.Sprintf("%s/DEBUG.log",logDir)),
		logrus.InfoLevel:  GetWriter(fmt.Sprintf("%s/INFO.log",logDir)),
		logrus.WarnLevel: GetWriter(fmt.Sprintf("%s/WARN.log",logDir)),
		logrus.ErrorLevel: GetWriter(fmt.Sprintf("%s/ERROR.log",logDir)),
		logrus.FatalLevel: GetWriter(fmt.Sprintf("%s/FATAL.log",logDir)),
		logrus.PanicLevel: GetWriter(fmt.Sprintf("%s/PANIC.log",logDir)),
	}
	Logger.AddHook(lfshook.NewHook(
		pathMap,
		&logrus.TextFormatter{},		//普通文本模式
	))
	//输出到单个文件
	//file, err := os.OpenFile("./logs/server.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	//if err == nil {
	//	Logger.SetOutput(file) //将日志输出到此文件
	//} else {
	//	fmt.Println("创建日志文件失败：" + err.Error())
	//}

	//取消日志标准输出(终端输出)
	if conf.GlobalConfig.Log.Level != "DEBUG" {
		nullFile, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err == nil {
			Logger.SetOutput(nullFile) //将日志输出到此文件
		} else {
			log.Panic("打开null文件失败")
		}
	}
}

//得到日志切割的输出对象
func GetWriter(pathLog string) *rotatelogs.RotateLogs{
	writer, err := rotatelogs.New(
		pathLog+".%Y%m%d%H",		        //日志文件后缀：年月日时
		rotatelogs.WithLinkName(pathLog),	//为当前正在输出的日志文件建立软连接
		rotatelogs.WithRotationCount(conf.GlobalConfig.Log.RotationCount),//日志文件保存的个数(包括当前正在输出的日志)
		rotatelogs.WithRotationTime(time.Duration(1)*time.Hour),		//设置日志分割的时间(隔多久分割一次)
	)
	if err!=nil{
		log.Panic("获取日志切割的对象失败：",err.Error())
	}
	return writer
}

//指定日志等级输出文件名+行号的hook
type PrintFileAndNumHook struct {
}
func (hook *PrintFileAndNumHook) Fire(entry *logrus.Entry) error {
	pcs := make([]uintptr,10)			//容量要足够，否则会被截断
	i := runtime.Callers(0,pcs)
	entry.Data["func"] = runtime.FuncForPC(pcs[i-2]).Name()
	file,line := runtime.FuncForPC(pcs[i-2]).FileLine(pcs[i-2])
	entry.Data["filename"] = file
	entry.Data["lineNum"] = line
	return nil
}
func (hook *PrintFileAndNumHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel, logrus.WarnLevel} //在这些Levels上生效
}
