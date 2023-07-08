package application

import (
	"context"
	"io"

	"github.com/starter-go/application/arguments"
	"github.com/starter-go/application/attributes"
	"github.com/starter-go/application/components"
	"github.com/starter-go/application/environment"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/application/resources"
	"github.com/starter-go/base/safe"
)

// Context 表示一个 starter 应用上下文
type Context interface {
	io.Closer
	context.Context

	Mode() safe.Mode

	// NewChild() Context
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
