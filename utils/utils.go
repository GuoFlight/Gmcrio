package utils

import (
	"Gmicro/conf"
	"context"
	uuid "github.com/satori/go.uuid"
)

// GenCtxWithTraceId 生成带traceId的ctx
func GenCtxWithTraceId(ctx context.Context) context.Context {
	traceId := uuid.NewV4().String()
	return context.WithValue(ctx, conf.KeyTraceId, traceId)
}
