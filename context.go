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
)

// Context 表示一个 starter 应用上下文
type Context interface {
	io.Closer

	context.Context

	NewChild() Context

	GetArguments() arguments.Table
	GetAttributes() attributes.Table
	GetComponents() components.Table
	GetEnvironment() environment.Table
	GetProperties() properties.Table
	GetResources() resources.Table

	GetComponent(selector string) (any, error)
	GetComponentByID(id string) (any, error)
	ListComponentIDs() []string
}
