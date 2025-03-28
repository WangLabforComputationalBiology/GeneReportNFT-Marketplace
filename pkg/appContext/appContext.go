package appContext

import (
	"GeneReport_platform/configs"
	"context"
	"time"
)

// NewTimeoutContext 创建一个带超时的上下文
func NewTimeoutContext() context.Context {
	timeout := time.Duration(configs.GlobalConfig.CtxConfig.Timeout) * time.Second
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	return ctx
}
func NewTimeoutContextWithCancel() (context.Context, context.CancelFunc) {
	timeout := time.Duration(configs.GlobalConfig.CtxConfig.Timeout) * time.Second
	return context.WithTimeout(context.Background(), timeout)
}

// NewTimeoutContextByParent 创建基于父上下文的带超时上下文
func NewTimeoutContextByParent(parent context.Context) context.Context {
	timeout := time.Duration(configs.GlobalConfig.CtxConfig.Timeout) * time.Second
	ctx, _ := context.WithTimeout(parent, timeout)
	return ctx
}
func NewTimeoutContextWithCancelByParent(parent context.Context) (context.Context, context.CancelFunc) {
	timeout := time.Duration(configs.GlobalConfig.CtxConfig.Timeout) * time.Second
	return context.WithTimeout(parent, timeout)
}
