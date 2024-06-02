package api

import (
	v1 "Gmicro/api/v1"
	"Gmicro/conf"
	"Gmicro/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartHttpServer() {
	router := gin.New()
	router.Use(gin.Recovery())

	// 子路由v1
	subRouterV1 := router.Group("/v1")
	subRouterV1.GET("/health", v1.Health)
	subRouterV1.GET("/test", v1.Test)

	// 启动http服务
	logger.GLogger.Info("开始启动http服务")
	router.Run(fmt.Sprintf(":%d", conf.GConf.Http.Port))
}
