package implcom

import "github.com/starter-go/application/components"

type comHolder struct {
	info    components.Info
	factory components.Factory
	cache   cache
}

func (inst *comHolder) init(i components.Info, f components.Factory) {

	scope := i.Scope()

	if scope == components.ScopePrototype {
		inst.cache = &prototypeCache{}
	} else {
		inst.cache = &singletonCache{}
	}

	inst.info = i
	inst.factory = f
}

func (inst *comHolder) _Impl() components.Holder {
	return inst
}

func (inst *comHolder) Info() components.Info {
	return inst.info
}

func (inst *comHolder) Factory() components.Factory {
	return inst.factory
}

func (inst *comHolder) GetInstance() (components.Instance, error) {
	i := inst.info
	f := inst.factory
	c := inst.cache
	if c == nil {
		c := &singletonCache{}
		inst.cache = c
	}
	return c.GetInstance(i, f)
}

func (inst *comHolder) NewRef(sel components.Selector) components.Ref {
	r := &comRef{
		holder: inst,
		sel:    sel,
	}
	return r
}
