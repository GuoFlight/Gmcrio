package http

import (
	"Gmicro/conf"
	"Gmicro/http/httpCommon"
	v1 "Gmicro/http/v1"
	"Gmicro/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartHttpServer(){
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(httpCommon.LoggerMiddleware()) //使用日志中间件

	//子路由v1
	subRouterV1 := router.Group("/v1")
	subRouterV1.GET("/health", v1.Health)

	//启动http服务
	logger.Logger.Info("开始启动http服务")
	router.Run(fmt.Sprintf(":%d",conf.GlobalConfig.Http.Port))
}
