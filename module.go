package application

import (
	"github.com/starter-go/application/components"
	"github.com/starter-go/application/resources"
)

// Module 表示一个 Starter 模块
type Module interface {
	Name() string
	Version() string
	Revision() int
	Dependencies() []Module
	Resources() resources.Table
	RegisterComponents(r components.Registry) error
}

////////////////////////////////////////////////////////////////////////////////

type myModule struct {
	name     string
	version  string
	revision int
	deps     []Module
	res      resources.Table
	comRegFn components.RegistryHandlerFunc
}

func (inst *myModule) _Impl() Module {
	return inst
}

func (inst *myModule) Name() string {
	return inst.name
}

func (inst *myModule) Version() string {
	return inst.version
}

func (inst *myModule) Revision() int {
	return inst.revision
}

func (inst *myModule) Dependencies() []Module {
	return inst.deps
}

func (inst *myModule) Resources() resources.Table {
	return inst.res
}

func (inst *myModule) RegisterComponents(r components.Registry) error {
	fn := inst.comRegFn
	if fn == nil || r == nil {
		return nil
	}
	return fn(r)
}
