package param

import (
	"github.com/harluo/di/internal/internal/runtime"
)

type Get struct {
	*Container

	Getters []runtime.Getter
}

func NewGet(container *Container, getters []runtime.Getter) *Get {
	return &Get{
		Container: container,

		Getters: getters,
	}
}
