package core

import (
	"github.com/harluo/di/internal/core/internal"
	"github.com/harluo/di/internal/internal/builder"
	"github.com/harluo/di/internal/internal/param"
	"github.com/harluo/di/internal/internal/runtime"
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

func (d *Container) Put(required runtime.Constructor, optionals ...runtime.Constructor) *builder.Put {
	return builder.NewPut(d.container, append([]runtime.Constructor{required}, optionals...))
}

func (d *Container) Get(required runtime.Getter, optionals ...runtime.Getter) *builder.Get {
	return builder.NewGet(d.container, append([]runtime.Getter{required}, optionals...))
}
