package container

import (
	"reflect"
	"runtime"

	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/harluo/di/internal/internal/internal/message"
	"github.com/harluo/di/internal/internal/internal/param"
	"go.uber.org/dig"
)

type Get struct {
	container *dig.Container
	params    *param.Get
}

func NewGet(container *dig.Container, params *param.Get) *Get {
	return &Get{
		container: container,
		params:    params,
	}
}

func (g *Get) Apply() {
	if err := g.Inject(); nil != err {
		panic(err)
	}
}

func (g *Get) Inject() (err error) {
	for _, get := range g.params.Getters {
		err = g.invoke(get)
		if nil != err {
			break
		}
	}

	return
}

func (g *Get) invoke(getter any) error {
	return g.container.Invoke(getter)
}

func (g *Get) validate(constructor any) (err error) {
	if g.params.Validate {
		return
	}

	typ := reflect.TypeOf(constructor)
	if reflect.Func != typ.Kind() { // 构造方法必须是方法，不能是其它类型
		err = exception.New().Message(message.ConstructorMustFunc).Field(field.New("constructor", typ.String())).Build()
	} else if 0 == typ.NumOut() { // 构造方法必须有返回值
		name := runtime.FuncForPC(reflect.ValueOf(constructor).Pointer()).Name()
		err = exception.New().Message(message.ConstructorMustReturn).Field(field.New("constructor", name)).Build()
	}

	return
}
