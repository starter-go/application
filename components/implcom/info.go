package implcom

import "github.com/starter-go/application/components"

type comInfo struct {
	id      components.ID
	scope   components.Scope
	aliases []string
	classes []string
}

func (inst *comInfo) _Impl() components.Info {
	return inst
}

func (inst *comInfo) ID() components.ID {
	return inst.id
}

func (inst *comInfo) Aliases() []string {
	return inst.aliases
}

func (inst *comInfo) Classes() []string {
	return inst.classes
}

func (inst *comInfo) Scope() components.Scope {
	return inst.scope
}
