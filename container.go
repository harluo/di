package di

import (
	"github.com/harluo/di/internal/builder"
	"github.com/harluo/di/internal/core"
)

// Container 依赖管理容器
type Container = core.Container

func New() *builder.Container {
	return builder.NewContainer()
}
