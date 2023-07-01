package components

import "context"

////////////////////////////////////////////////////////////////////////////////

// Scope 表示组件的作用域
type Scope int

// 定义组件的各种作用域
const (
	ScopeSingleton Scope = 1
	ScopePrototype Scope = 2
)

////////////////////////////////////////////////////////////////////////////////

// ID 表示一个组件ID
type ID string

func (id ID) String() string {
	return string(id)
}

////////////////////////////////////////////////////////////////////////////////

// Selector 表示一个组件选择器
type Selector string

func (sel Selector) String() string {
	return string(sel)
}

// SelectorForAlias 创建 alias 选择器
func SelectorForAlias(alias string) Selector {
	str := "#" + alias
	return Selector(str)
}

// SelectorForClass 创建 class 选择器
func SelectorForClass(cname string) Selector {
	str := "." + cname
	return Selector(str)
}

// SelectorForID 创建 ID 选择器
func SelectorForID(id ID) Selector {
	str := "#" + id.String()
	return Selector(str)
}

////////////////////////////////////////////////////////////////////////////////

// Instance 用于创建组件实例
type Instance interface {
	Ready() bool
	Get() any
	Info() Info
	Inject(i Injection) error
}

// Factory 用于创建组件实例
type Factory interface {
	CreateInstance(info Info) (Instance, error)
}

// Info 提供组件的相关信息
type Info interface {
	ID() ID
	Aliases() []string
	Classes() []string
	Scope() Scope
}

// Holder 根据作用域，持有一个或一组组件实例
type Holder interface {
	Info() Info
	Factory() Factory
	GetInstance() (Instance, error)
	NewRef(sel Selector) Ref
}

// Ref 表示一个对 Holder 的引用
type Ref interface {
	Selector() Selector
	Holder() Holder
}

// Injection 表示一个依赖注入上下文
type Injection interface {
	context.Context

	Select(selector Selector) ([]Instance, error)

	SelectOne(selector Selector) (Instance, error)

	GetByID(id ID) (Instance, error)

	GetWithHolder(holder Holder) (Instance, error)
}
