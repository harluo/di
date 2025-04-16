package di

import (
	"github.com/heluon/di/internal/builder"
	"github.com/heluon/di/internal/core"
)

// Container 依赖管理容器
type Container = core.Container

func New() *builder.Container {
	return builder.NewContainer()
}
