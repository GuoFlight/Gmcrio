package api

import (
	"Gmicro/myctx"
	"github.com/gin-gonic/gin"
)

func traceIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		myctx.GCtx.Gin.SetTraceId(c)
	}
}
