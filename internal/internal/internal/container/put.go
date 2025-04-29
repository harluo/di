package container

import (
	"reflect"
	"runtime"

	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/harluo/di/internal/internal/constant"
	"github.com/harluo/di/internal/internal/internal/message"
	"github.com/harluo/di/internal/internal/internal/param"
	"go.uber.org/dig"
)

type Put struct {
	container *dig.Container
	params    *param.Put
}

func NewPut(container *dig.Container, params *param.Put) *Put {
	return &Put{
		container: container,
		params:    params,
	}
}

func (p *Put) Apply() {
	if err := p.Inject(); nil != err {
		panic(err)
	}
}

func (p *Put) Inject() (err error) {
	for _, name := range p.params.Names {
		for _, group := range p.params.Groups {
			for _, constructor := range p.params.Constructors {
				if ve := p.validate(constructor); nil != ve {
					err = ve
				} else if pe := p.provide(constructor, name, group); nil != pe {
					err = pe
				}
				if nil != err {
					break
				}
			}
		}
	}

	return
}

func (p *Put) provide(constructor any, name string, group string) error {
	options := make([]dig.ProvideOption, 0)
	if constant.DependencyNone != name {
		options = append(options, dig.Name(name))
	}
	if constant.DependencyNone != group {
		options = append(options, dig.Group(group))
	}

	return p.container.Provide(constructor, options...)
}

func (p *Put) validate(constructor any) (err error) {
	if p.params.Validate {
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
