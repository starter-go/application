package implcom

import (
	"fmt"
	"time"

	"github.com/starter-go/application"
	"github.com/starter-go/application/arguments"
	"github.com/starter-go/application/attributes"
	"github.com/starter-go/application/components"
	"github.com/starter-go/application/environment"
	"github.com/starter-go/application/parameters"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/application/resources"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/base/safe"
)

// context 实现 application.Context
type context struct {
	mode safe.Mode

	moduleMain application.Module
	modules    []application.Module

	lm application.LifeManager

	createdAt lang.Time
	startedAt lang.Time
	stoppedAt lang.Time
	destroyAt lang.Time

	args   arguments.Table
	atts   attributes.Table
	com    components.Table
	env    environment.Table
	params parameters.Table
	props  properties.Table
	res    resources.Table
}

func (ctx *context) _Impl() application.Context {
	return ctx
}

func (ctx *context) Deadline() (deadline time.Time, ok bool) {
	now := time.Now()
	return now, false
}

func (ctx *context) Done() <-chan struct{} {

	return nil
}

func (ctx *context) Err() error {
	return nil
}

func (ctx *context) Value(key any) any {
	return nil
}

func (ctx *context) Close() error {
	return nil
}

func (ctx *context) Mode() safe.Mode {
	return ctx.mode
}

func (ctx *context) NewChild() application.Context {
	child := &context{}
	// todo ...
	return child
}

func (ctx *context) NewInjection(scope components.Scope) application.Injection {
	ai1 := &injection{}
	ai2 := ai1.init(ctx, scope)
	return ai2
}

func (ctx *context) GetArguments() arguments.Table {
	return ctx.args
}

func (ctx *context) GetAttributes() attributes.Table {
	return ctx.atts
}

func (ctx *context) GetComponents() components.Table {
	return ctx.com
}

func (ctx *context) GetEnvironment() environment.Table {
	return ctx.env
}

func (ctx *context) GetProperties() properties.Table {
	return ctx.props
}

func (ctx *context) GetResources() resources.Table {
	return ctx.res
}

func (ctx *context) GetModules() []application.Module {
	src := ctx.modules
	size := len(src)
	dst := make([]application.Module, size)
	copy(dst, src)
	return dst
}

func (ctx *context) GetMainModule() application.Module {
	return ctx.moduleMain
}

func (ctx *context) SelectComponent(selector components.Selector) (any, error) {
	hlist, err := ctx.com.Select(selector)
	if err != nil {
		return nil, err
	}
	olist, err := ctx.activeComponents(hlist...)
	if err != nil {
		return nil, err
	}
	return ctx.onlyOneComponent(olist, selector)
}

func (ctx *context) SelectComponents(selector components.Selector) ([]any, error) {
	hlist, err := ctx.com.Select(selector)
	if err != nil {
		return nil, err
	}
	return ctx.activeComponents(hlist...)
}

func (ctx *context) GetComponent(id components.ID) (any, error) {
	h, err := ctx.com.Get(id)
	if err != nil {
		return nil, err
	}
	olist, err := ctx.activeComponents(h)
	if err != nil {
		return nil, err
	}
	sel := components.SelectorForID(id)
	return ctx.onlyOneComponent(olist, sel)
}

func (ctx *context) activeComponents(holders ...components.Holder) ([]any, error) {

	instanceList := make([]components.Instance, 0)
	injection := ctx.NewInjection(0) // todo: scope ...

	for _, h := range holders {
		instance, err := injection.GetWithHolder(h)
		if err != nil {
			return nil, err
		}
		instanceList = append(instanceList, instance)
	}

	err := injection.Complete()
	if err != nil {
		return nil, err
	}

	dst := make([]any, 0)
	for _, instance := range instanceList {
		obj := instance.Get()
		dst = append(dst, obj)
	}
	return dst, nil
}

func (ctx *context) onlyOneComponent(olist []any, sel components.Selector) (any, error) {
	count := 0
	if olist != nil {
		count = len(olist)
		if count == 1 {
			obj := olist[0]
			if obj != nil {
				return obj, nil
			}
		}
	}
	return nil, fmt.Errorf("there are %d components with selector:%s", count, sel.String())
}

func (ctx *context) ListComponentIDs() []components.ID {
	return ctx.com.ListIDs()
}

func (ctx *context) GetLifeManager() application.LifeManager {
	return ctx.lm
}
