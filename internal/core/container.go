package core

import (
	"github.com/harluo/di/internal/core/internal"
	"github.com/harluo/di/internal/internal/builder"
	"github.com/harluo/di/internal/internal/param"
	"go.uber.org/dig"
)

var shadow *Container

type Container struct {
	container *dig.Container
	params    *param.Container
}

func NewContainer(params *param.Container) (container *Container) {
	internal.Once.Do(func() {
		shadow = new(Container)
		shadow.container = dig.New()
		shadow.params = params
	})
	container = shadow

	return
}

func (c *Container) Dependency() *builder.Dependency {
	return builder.NewDependency(c.container, c.params)
}
