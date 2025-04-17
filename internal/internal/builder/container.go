package builder

import (
	"github.com/harluo/di/internal/internal/builder/internal"
	"github.com/harluo/di/internal/internal/param"
)

type Container[F any] struct {
	from   *F
	params *param.Container
}

func NewContainer[F any](from *F, params *param.Container) *Container[F] {
	return &Container[F]{
		from:   from,
		params: params,
	}
}

func (c *Container[F]) Validate() *F {
	return c.set(func() {
		c.params.Validate = true
	})
}

func (c *Container[F]) Invalidate() *F {
	return c.set(func() {
		c.params.Validate = false
	})
}

func (c *Container[F]) set(set internal.Set) (from *F) {
	set()
	from = c.from

	return
}
