package api

import (
	v1 "Gmicro/api/v1"
	"Gmicro/conf"
	"Gmicro/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tylerb/graceful"
	"net/http"
)

var Done = make(chan bool, 1)

func StartHttpServer() {
	router := gin.New()
	router.Use(traceIdMiddleware(), logMiddleware(), gin.Recovery())

	// 子路由api
	rApi := router.Group("/api")

	// 子路由v1
	rApiV1 := rApi.Group("/v1")
	rApiV1.GET("/health", v1.Health)

	// 子路由/v1/jwt
	rApiV1Jwt := rApiV1.Group("/jwt")
	rApiV1Jwt.POST("/login", v1.GJwt.Login)

	rApiV1Admin := rApiV1.Group("/admin")
	rApiV1Admin.Use(v1.GAuth.Auth())
	rApiV1Admin.GET("/testAuth", v1.TestAuth) // 测试

	// 启动http服务
	logger.GLogger.Info("开始启动http服务")
	server := graceful.Server{
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%d", conf.GConf.Http.Port),
			Handler: router,
		},
		BeforeShutdown: func() bool {
			logger.GLogger.Info("即将关闭http服务")
			return true
		},
	}
	err := server.ListenAndServe()
	if err != nil {
		logger.GLogger.Fatal(err)
	}
	logger.GLogger.Info("http服务已退出")
	Done <- true
}
