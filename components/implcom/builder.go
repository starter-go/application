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
	// args   arguments.Table
	// atts   attributes.Table
	// env    environment.Table
	// params parameters.Table
	// props  properties.Table
	// res    resources.Table
	// com    components.Table

	collections application.Collections

	registry componentTableBuilder

	mode safe.Mode
}

// SetArguments ...
func (inst *Builder) SetArguments(t arguments.Table) {
	inst.collections.Arguments = t
}

// SetAttributes ...
func (inst *Builder) SetAttributes(t attributes.Table) {
	inst.collections.Attributes = t
}

// SetEnv ...
func (inst *Builder) SetEnv(t environment.Table) {
	inst.collections.Environment = t
}

// SetParameters ...
func (inst *Builder) SetParameters(t parameters.Table) {
	inst.collections.Parameters = t
}

// SetProperties ...
func (inst *Builder) SetProperties(t properties.Table) {
	inst.collections.Properties = t
}

// SetCollections ...
func (inst *Builder) SetCollections(c *application.Collections) {
	if c == nil {
		return
	}
	inst.collections = *c
}

// SetResources ...
func (inst *Builder) SetResources(t resources.Table) {
	inst.collections.Resources = t
}

// Registry 用来获取组件注册接口
func (inst *Builder) Registry() application.ComponentRegistry {
	return &inst.registry
}

// Create 创建 application.Context
func (inst *Builder) Create() (application.Context, error) {

	coll := &inst.collections
	mode := inst.mode

	if mode == nil {
		mode = safe.Default()
	}
	coll.Complete(mode)

	args := coll.Arguments
	atts := coll.Attributes
	env := coll.Environment
	params := coll.Parameters
	props := coll.Properties
	res := coll.Resources

	comtab := inst.registry.Create(mode)

	ctx := &appContext{
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
	ctx := injection.GetContext()
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
func NewBuilder(mode safe.Mode) *Builder {
	return &Builder{mode: mode}
}
