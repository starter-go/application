package boot

import "github.com/starter-go/application/components"

type instanceCache interface {
	GetInstance(i components.Info, f components.Factory) (components.Instance, error)
}

////////////////////////////////////////////////////////////////////////////////

type comHolder struct {
	info    components.Info
	factory components.Factory
	cache   instanceCache
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

////////////////////////////////////////////////////////////////////////////////

type comRef struct {
	sel    components.Selector
	holder components.Holder
}

func (inst *comRef) _Impl() components.Ref {
	return inst
}

func (inst *comRef) Selector() components.Selector {
	return inst.sel
}

func (inst *comRef) Holder() components.Holder {
	return inst.holder
}

////////////////////////////////////////////////////////////////////////////////

type singletonCache struct {
	instance components.Instance
}

func (inst *singletonCache) GetInstance(i components.Info, f components.Factory) (components.Instance, error) {
	o := inst.instance
	if o != nil {
		return o, nil
	}
	o, err := f.CreateInstance(i)
	if err != nil {
		return nil, err
	}
	inst.instance = o
	return o, nil
}

////////////////////////////////////////////////////////////////////////////////

type prototypeCache struct{}

func (inst *prototypeCache) GetInstance(i components.Info, f components.Factory) (components.Instance, error) {
	return f.CreateInstance(i)
}

////////////////////////////////////////////////////////////////////////////////
