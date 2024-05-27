package main

import (
	"Gmicro/api"
	"Gmicro/conf"
	"Gmicro/flag"
	"Gmicro/logger"
	"Gmicro/timer"
	"fmt"
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
	logger.InitLog()

	// 初始化周期性任务
	go timer.InitTimer()
	<-timer.InitDone

	// 启动http服务
	go api.StartHttpServer()

	// 阻塞主进程
	select {}
}
