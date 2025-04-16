package builder

import (
	"github.com/heluon/di/internal/core"
	"github.com/heluon/di/internal/internal/builder"
	"github.com/heluon/di/internal/internal/param"
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

func (c *Container) Get() *core.Container {
	return core.NewContainer(c.params)
}
