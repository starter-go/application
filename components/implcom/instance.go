package implcom

import (
	"github.com/starter-go/application"
	"github.com/starter-go/application/components"
)

type comInstance struct {
	reg          *components.Registration
	info         components.Info
	target       any
	hasInject    bool
	hasAddToLife bool
}

func (inst *comInstance) _Impl() components.Instance {
	return inst
}

func (inst *comInstance) Get() any {
	return inst.target
}

func (inst *comInstance) Info() components.Info {
	return inst.info
}

func (inst *comInstance) Ready() bool {
	return inst.hasInject
}

func (inst *comInstance) Inject(i components.Injection) error {

	if inst.hasInject {
		return nil
	}

	inst.hasInject = true
	fn := inst.reg.InjectFunc
	t := inst.target

	if fn == nil || t == nil {
		return nil
	}

	err := fn(i, t)
	if err != nil {
		return err
	}

	return inst.addToLife(i)
}

func (inst *comInstance) addToLife(i components.Injection) error {

	if inst.hasAddToLife {
		return nil
	}

	inst.hasAddToLife = true
	ix := i.(application.Injection)
	t := inst.target
	lc, ok := t.(application.Lifecycle)
	if !ok || lc == nil {
		return nil
	}

	life := lc.Life()
	if life == nil {
		return nil
	}

	ix.LifeManager().Add(life)
	return nil
}
