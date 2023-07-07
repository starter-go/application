package application

import (
	"context"
	"io"

	"github.com/starter-go/application/arguments"
	"github.com/starter-go/application/attributes"
	"github.com/starter-go/application/components"
	"github.com/starter-go/application/environment"
	"github.com/starter-go/application/parameters"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/application/resources"
	"github.com/starter-go/base/safe"
)

// Context 表示一个 starter 应用上下文
type Context interface {
	io.Closer
	context.Context

	Mode() safe.Mode

	NewChild() Context
	NewInjection(scope components.Scope) Injection

	GetArguments() arguments.Table
	GetAttributes() attributes.Table
	GetComponents() components.Table
	GetEnvironment() environment.Table
	GetProperties() properties.Table
	GetResources() resources.Table

	GetModules() []Module
	GetMainModule() Module

	GetLifeManager() LifeManager

	SelectComponents(selector components.Selector) ([]any, error)
	SelectComponent(selector components.Selector) (any, error)
	GetComponent(id components.ID) (any, error)
	ListComponentIDs() []components.ID
}

////////////////////////////////////////////////////////////////////////////////

// Collections 持有上下文中的各种资源集合
type Collections struct {
	Arguments   arguments.Table
	Attributes  attributes.Table
	Environment environment.Table
	Parameters  parameters.Table
	Properties  properties.Table
	Resources   resources.Table
}

// Complete 完善集合中的内容（如果需要）
func (inst *Collections) Complete(mode safe.Mode) {

	if inst.Arguments == nil {
		inst.Arguments = arguments.NewTable(nil, mode)
	}

	if inst.Attributes == nil {
		inst.Attributes = attributes.NewTable(mode)
	}

	if inst.Environment == nil {
		inst.Environment = environment.NewTable(mode)
	}

	if inst.Parameters == nil {
		inst.Parameters = parameters.NewTable(mode)
	}

	if inst.Properties == nil {
		inst.Properties = properties.NewTable(mode)
	}

	if inst.Resources == nil {
		empty := make(map[string]resources.Resource)
		inst.Resources = resources.NewTable(empty, mode)
	}

}
