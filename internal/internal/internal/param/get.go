package param

import (
	"github.com/heluon/di/internal/internal/runtime"
)

type Get struct {
	Getter runtime.Getter
}

func NewGet(getter runtime.Getter) *Get {
	return &Get{
		Getter: getter,
	}
}
