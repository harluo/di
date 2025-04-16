package builder

import (
	"github.com/heluon/di/internal/internal/internal/container"
	"github.com/heluon/di/internal/internal/internal/param"
	"github.com/heluon/di/internal/internal/runtime"
	"go.uber.org/dig"
)

type Dependency struct {
	*shadow

	container *dig.Container
	params    *param.Dependency
}

func NewDependency(container *dig.Container, parent *param.Container) (dependency *Dependency) {
	dependency = new(Dependency)

	params := new(param.Container)
	*params = *parent
	dependency.shadow = NewContainer(dependency, params)

	dependency.container = container
	dependency.params = param.NewDependency(parent)

	return
}

func (d *Dependency) Put(constructor runtime.Constructor) *Put {
	return NewPut(d, constructor)
}

func (d *Dependency) Puts(required runtime.Constructor, optionals ...runtime.Constructor) (dependency *Dependency) {
	d.params.Puts = append(d.params.Puts, param.NewPut(required))
	for _, getter := range optionals {
		d.params.Puts = append(d.params.Puts, param.NewPut(getter))
	}
	dependency = d

	return
}

func (d *Dependency) Get(getter runtime.Getter) *Get {
	return NewGet(d, getter)
}

func (d *Dependency) Gets(required runtime.Getter, optionals ...runtime.Getter) (dependency *Dependency) {
	d.params.Gets = append(d.params.Gets, param.NewGet(required))
	for _, getter := range optionals {
		d.params.Gets = append(d.params.Gets, param.NewGet(getter))
	}
	dependency = d

	return
}

func (d *Dependency) Build() *container.Dependency {
	return container.NewDependency(d.container, d.params)
}
