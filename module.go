package application

import (
	"github.com/starter-go/application/resources"
)

const (
	theModuleName     = "github.com/starter-go/application"
	theModuleVersion  = "v0.9.24"
	theModuleRevision = 8
)

////////////////////////////////////////////////////////////////////////////////

// Module 表示一个 Starter 模块
type Module interface {
	Name() string
	Version() string
	Revision() int
	Dependencies() []Module
	Resources() resources.Table
	RegisterComponents(r ComponentRegistry) error
}

////////////////////////////////////////////////////////////////////////////////

type myModule struct {
	name     string
	version  string
	revision int
	deps     []Module
	res      resources.Table
	registry ComponentRegistryFunc
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

func (inst *myModule) RegisterComponents(r ComponentRegistry) error {
	fn := inst.registry
	if fn == nil || r == nil {
		return nil
	}
	return fn(r)
}
