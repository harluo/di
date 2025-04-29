package builder

import (
	"github.com/harluo/di/internal/internal/internal/container"
	"github.com/harluo/di/internal/internal/internal/param"
	"github.com/harluo/di/internal/internal/runtime"
	"go.uber.org/dig"
)

type Get struct {
	*shadowGet

	container *dig.Container
	params    *param.Get
}

func NewGet(container *dig.Container, getters []runtime.Getter) (get *Get) {
	get = new(Get)
	params := new(param.Container)
	get.shadowGet = NewContainer(get, params)
	get.container = container
	get.params = param.NewGet(params, getters)

	return
}

func (g *Get) Build() *container.Get {
	return container.NewGet(g.container, g.params)
}
