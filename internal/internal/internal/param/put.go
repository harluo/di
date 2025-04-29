package param

import (
	"github.com/goexl/gox"
	"github.com/harluo/di/internal/internal/runtime"
)

type Put struct {
	*Container

	Constructors []runtime.Constructor
	Names        map[string]*gox.Empty
	Groups       map[string]*gox.Empty
}

func NewPut(container *Container, constructors []runtime.Constructor) *Put {
	return &Put{
		Container: container,

		Constructors: constructors,
		Names:        make(map[string]*gox.Empty),
		Groups:       make(map[string]*gox.Empty),
	}
}
