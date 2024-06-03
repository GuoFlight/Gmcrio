package v1

import (
	"Gmicro/conf"
	"Gmicro/utils"
	"context"
	"github.com/gin-gonic/gin"
)

type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Health(c *gin.Context) {
	c.String(200, "alive")
}

func Test(c *gin.Context) {
	ctx := getCtxWithTraceId(c)
	v, _ := ctx.Value(conf.TraceIdName).(string)
	c.String(200, v)
}
func TestAuth(c *gin.Context) {
	if hasPerm := GAuth.CheckPermWrite(c); !hasPerm {
		return
	}
	c.JSON(200, Res{
		Code: 0,
		Msg:  "success.",
	})
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
