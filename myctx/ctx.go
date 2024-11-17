package myctx

import (
	"Gmicro/conf"
	"Gmicro/logger"
	"context"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// GenWithTraceId 生成带traceId的ctx
func GenWithTraceId(ctx context.Context) context.Context {
	traceId := uuid.NewV4().String()
	return context.WithValue(ctx, conf.KeyTraceId, traceId)
}

// GenFromGin 生成带traceId、用户名的ctx
func GenFromGin(ctxGin *gin.Context) context.Context {
	// 未来若有必要，可以从请求头中读取TraceID
	ctx := GenWithTraceId(context.Background())

	// 写入用户名
	username, ok := ctxGin.Get(conf.KeyUsername)
	if ok {
		return SetOpUser(ctx, username)
	} else {
		return SetOpUser(ctx, "")
	}
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
