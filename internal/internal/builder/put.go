package builder

import (
	"github.com/harluo/di/internal/internal/internal/container"
	"github.com/harluo/di/internal/internal/internal/param"
	"github.com/harluo/di/internal/internal/runtime"
	"go.uber.org/dig"
)

type Put struct {
	*shadowPut

	container *dig.Container
	params    *param.Put
}

func NewPut(container *dig.Container, constructors []runtime.Constructor) (put *Put) {
	put = new(Put)
	params := new(param.Container)
	put.shadowPut = NewContainer(put, params)
	put.container = container
	put.params = param.NewPut(params, constructors)

	return
}

func (p *Put) Name(required string, optionals ...string) (put *Put) {
	p.params.Names = append(p.params.Names, append([]string{required}, optionals...)...)
	put = p

	return
}

func (p *Put) Group(required string, optionals ...string) (put *Put) {
	p.params.Groups = append(p.params.Groups, append([]string{required}, optionals...)...)
	put = p

	return
}

func (p *Put) Build() *container.Put {
	return container.NewPut(p.container, p.params)
}
