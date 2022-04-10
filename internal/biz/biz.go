package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
// NewSet 接收的参数是 ...interface{} 什么意思
var ProviderSet = wire.NewSet(
	NewGreeterUsecase,
	NewDemoUsecase,
)
