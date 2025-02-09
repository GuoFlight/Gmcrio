package v1

import (
	"Gmicro/models"
	"Gmicro/myctx"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Res models.Res

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, Res{
		Code: 0,
		Msg:  "alive",
	})
}

func TestAuth(c *gin.Context) {
	ctx := myctx.GCtx.Gin.GenCtx(c)
	opUser := myctx.GetOpUser(ctx)
	traceId := myctx.GetTraceId(ctx)
	c.JSON(http.StatusOK, Res{
		Code: 0,
		Msg:  "success.",
		Data: fmt.Sprintf("用户名:%s traceId: %s", opUser, traceId),
	})
}
