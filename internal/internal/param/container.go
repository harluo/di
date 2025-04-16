package param

type Container struct {
	// 合法性验证，包括
	// 启动器构造方法合法性验证
	// 构造方法合法性验证
	Validate bool
}

func NewContainer() *Container {
	return &Container{
		Validate: true, // 验证构造方法的合法性
	}
}
