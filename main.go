package main

import (
	"Gmicro/conf"
	"Gmicro/flag"
	"Gmicro/http"
	"Gmicro/logger"
	"fmt"
)

func main() {
	//解析命令行
	flag.ParseConfig()

	//解析配置文件
	conf.ParseConfig(*flag.PathConfFile)

	//初始化日志
	logger.InitLog()
	logger.Logger.Info("日志初始化完成")

	//功能导航
	if *flag.Version {
		fmt.Println(conf.Version)
		return
	}

	//启动http服务
	go http.StartHttpServer()

	//阻塞主进程
	select {}
}

