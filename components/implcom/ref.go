package implcom

import "github.com/starter-go/application/components"

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
