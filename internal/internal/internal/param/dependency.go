package param

type Dependency struct {
	*Container

	Puts []*Put
	Gets []*Get
}

func NewDependency(container *Container) *Dependency {
	return &Dependency{
		Container: container,
	}
}

func (d *Dependency) Clear() {
	d.Puts = make([]*Put, 0, 1)
	d.Gets = make([]*Get, 0, 1)
}
