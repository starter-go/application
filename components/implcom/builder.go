package implcom

import (
	"github.com/starter-go/application"
	"github.com/starter-go/application/arguments"
	"github.com/starter-go/application/attributes"
	"github.com/starter-go/application/components"
	"github.com/starter-go/application/environment"
	"github.com/starter-go/application/parameters"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/application/resources"
	"github.com/starter-go/base/safe"
)

// Builder 用来创建 application.Context
type Builder struct {
	args   arguments.Table
	atts   attributes.Table
	env    environment.Table
	params parameters.Table
	props  properties.Table
	res    resources.Table
	// com    components.Table

	registry componentTableBuilder

	mode safe.Mode
}

// Registry 用来获取组件注册接口
func (inst *Builder) Registry() components.Registry {
	return &inst.registry
}

// Create 创建 application.Context
func (inst *Builder) Create() (application.Context, error) {

	mode := inst.mode
	if mode == nil {

	}

	args := inst.args
	atts := inst.atts
	env := inst.env
	params := inst.params
	props := inst.props
	res := inst.res

	comtab := inst.registry.Create(mode)

	ctx := &context{
		mode: mode,

		args:   args,
		atts:   atts,
		com:    comtab,
		env:    env,
		params: params,
		props:  props,
		res:    res,
	}

	injection := ctx.NewInjection(components.ScopeSingleton)
	err := inst.loadSingletonComponents(injection)
	if err != nil {
		return nil, err
	}
	ctx.lm = injection.LifeManager()

	return ctx, nil
}

func (inst *Builder) loadSingletonComponents(injection application.Injection) error {
	const singleton = components.ScopeSingleton
	ctx := injection.Parent()
	ctable := ctx.GetComponents()
	ids := ctable.ListIDs()
	for _, id := range ids {
		holder, err := ctable.Get(id)
		if err != nil {
			return err
		}
		scope := holder.Info().Scope()
		if scope == singleton {
			_, err := injection.GetWithHolder(holder)
			if err != nil {
				return err
			}
		}
	}
	return injection.Complete()
}

////////////////////////////////////////////////////////////////////////////////

// NewBuilder 新建一个上下文创建器
func NewBuilder() *Builder {
	return &Builder{}
}
