package v1

import (
	"Gmicro/conf"
	"Gmicro/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Res struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	TraceId string      `json:"traceId,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, Res{
		Code: 0,
		Msg:  "alive",
	})
}

func Test(c *gin.Context) {
	ctx := getCtxWithTraceId(c)
	v, _ := ctx.Value(conf.KeyTraceId).(string)
	c.String(http.StatusOK, v)
}
func TestAuth(c *gin.Context) {
	c.JSON(http.StatusOK, Res{
		Code: 0,
		Msg:  "success.",
	})
}

// 生成带traceId的ctx
func getCtxWithTraceId(c *gin.Context) context.Context {
	traceId, ok := c.Request.Header[conf.KeyTraceId]
	if ok {
		return context.WithValue(context.Background(), conf.KeyTraceId, traceId[0])
	} else {
		return utils.GenCtxWithTraceId(context.Background())
	}
}
