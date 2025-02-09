package main

import (
	"Gmicro/api"
	"Gmicro/conf"
	"Gmicro/db"
	"Gmicro/flag"
	"Gmicro/logger"
	"Gmicro/myctx"
	"Gmicro/timer"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 解析命令行
	flag.InitFlag()
	if *flag.Version {
		fmt.Println(conf.Version)
		return
	}

	// 解析配置文件
	conf.ParseConfig(*flag.PathConfFile)

	// 初始化日志
	log.Println("正在初始化日志")
	logger.InitLog()

	// 初始化数据库
	log.Println("正在初始化数据库")
	ctx := myctx.GenWithTraceId(context.Background(), myctx.GenTraceId())
	db.Init(ctx)

	// 初始化周期性任务
	log.Println("正在初始化周期性任务")
	go timer.InitTimer()
	<-timer.InitDone

	// 启动http服务
	log.Println("正在初始化web服务")
	go api.StartHttpServer()

	// 优雅退出
	sig := make(chan os.Signal)
	done := make(chan bool)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for {
			s := <-sig
			switch s {
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				logger.GLogger.Info("app收到退出信号：", s)
				<-timer.Done
				<-api.Done
				db.Exit()
				logger.GLogger.Info("app正常退出")
				done <- true
			default:
				logger.GLogger.Warn("app收到即将忽略的信号:", s)
			}
		}
	}()
	<-done
}
