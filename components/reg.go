package components

// Registry 提供组件注册服务
type Registry interface {
	New() *Registration
	Register(r *Registration)
}

////////////////////////////////////////////////////////////////////////////////

// RegistryHandlerFunc 是组件注册的处理函数
type RegistryHandlerFunc func(r Registry) error

////////////////////////////////////////////////////////////////////////////////

// Registration 组件注册信息
type Registration struct {
	ID         ID
	Classes    string // 类名列表，项与项之间以空格符(SPACE)分隔
	Aliases    string // 别名列表，项与项之间以空格符(SPACE)分隔
	Scope      string // 作用域 'singleton'|'prototype'
	NewFunc    func() any
	InjectFunc func(c Injection, instance any) error
}
