package v1

import (
	"Gmicro/conf"
	"Gmicro/utils"
	"context"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.String(200, "alive")
}

func Test(c *gin.Context) {
	ctx := getCtxWithTraceId(c)
	v, _ := ctx.Value(conf.TraceIdName).(string)
	c.String(200, v)
}

// 生成带traceId的ctx
func getCtxWithTraceId(c *gin.Context) context.Context {
	traceId, ok := c.Request.Header["Traceid"]
	if ok {
		return context.WithValue(context.Background(), conf.TraceIdName, traceId[0])
	} else {
		return utils.GenCtxWithTraceId(context.Background())
	}
}
