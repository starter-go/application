package components

import "context"

// Injection 表示一个依赖注入上下文
type Injection interface {
	context.Context

	Select(selector Selector) ([]Instance, error)

	SelectOne(selector Selector) (Instance, error)

	GetByID(id ID) (Instance, error)

	GetApplicationContext() context.Context

	GetProperty(selector Selector) (string, error)

	GetWithHolder(holder Holder) (Instance, error)

	Scope() Scope
}
