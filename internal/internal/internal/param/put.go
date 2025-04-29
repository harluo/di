package param

import (
	"github.com/harluo/di/internal/internal/runtime"
)

type Put struct {
	*Container

	Constructors []runtime.Constructor
	Names        []string
	Groups       []string
}

func NewPut(container *Container, constructors []runtime.Constructor) *Put {
	return &Put{
		Container: container,

		Constructors: constructors,
		Names:        make([]string, 0),
		Groups:       make([]string, 0),
	}
}
