package builder

import (
	"github.com/harluo/di/internal/core"
	"github.com/harluo/di/internal/internal/builder"
	"github.com/harluo/di/internal/internal/param"
)

type Container struct {
	*shadow

	params *param.Container
}

func NewContainer() (container *Container) {
	container = new(Container)

	params := param.NewContainer()
	container.shadow = builder.NewContainer[Container](container, params)
	container.params = params

	return
}

func (c *Container) Instance() *core.Container {
	return core.NewContainer(c.params)
}
