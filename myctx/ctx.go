package myctx

import (
	"Gmicro/conf"
	"Gmicro/logger"
	"context"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var GCtx Ctx

type Ctx struct {
	Gin Gin
}
type Gin struct{}

func GenTraceId() string {
	return uuid.NewV4().String()
}

// GenWithTraceId 生成带traceId的ctx
func GenWithTraceId(ctx context.Context, traceId string) context.Context {
	return context.WithValue(ctx, conf.KeyTraceId, traceId)
}

func GetTraceId(ctx context.Context) string {
	v, ok := ctx.Value(conf.KeyTraceId).(string)
	if !ok {
		return ""
	}
	return v
}

func SetOpUser(ctx context.Context, user any) context.Context {
	return context.WithValue(ctx, conf.KeyUsername, user)
}

// GetOpUser 从Ctx中获取操作用户
func GetOpUser(ctx context.Context) string {
	v, ok := ctx.Value(conf.KeyUsername).(string)
	if ok {
		return v
	} else {
		logger.ErrWithCtx(ctx, "从context中获取操作用户失败")
		return ""
	}
}

// GenCtx 生成带traceId、用户名的ctx
func (g Gin) GenCtx(c *gin.Context) context.Context {
	// 未来若有必要，可以从请求头中读取TraceID
	traceId := g.SetTraceId(c)
	ctx := GenWithTraceId(context.Background(), traceId)

	// 写入用户名
	username, ok := c.Get(conf.KeyUsername)
	if ok {
		return SetOpUser(ctx, username)
	} else {
		return SetOpUser(ctx, "")
	}
}
func (g Gin) SetTraceId(c *gin.Context) string {
	// 从req中得到TraceId
	traceId := g.GetTraceId(c)
	if traceId == "" {
		traceId = GenTraceId()
		c.Set(conf.KeyTraceId, traceId)
	}
	// 将traceId写入到Response的Header中
	c.Writer.Header()[conf.KeyTraceId] = []string{traceId}
	return traceId
}
func (g Gin) GetTraceId(c *gin.Context) string {
	traceIdAny, ok := c.Get(conf.KeyTraceId)
	if !ok {
		return ""
	}
	traceId, ok := traceIdAny.(string)
	if !ok {
		logger.GLogger.Error("从gin.Context中获取traceId失败")
		return ""
	}
	return traceId
}
func (g Gin) GetOpUser(c *gin.Context) string {
	usernameAny, ok := c.Get(conf.KeyUsername)
	if !ok {
		return ""
	}
	username, ok := usernameAny.(string)
	if !ok {
		logger.GLogger.Error("从gin.Context中获取username失败")
		return ""
	}
	return username
}
