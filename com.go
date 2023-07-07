package application

import "github.com/starter-go/application/components"

// ComponentRegistration 表示组件注册信息
type ComponentRegistration struct {
	ID components.ID

	Classes string // 类名列表，项与项之间以空格符(SPACE)分隔
	Aliases string // 别名列表，项与项之间以空格符(SPACE)分隔
	Scope   string // 作用域 'singleton'|'prototype'

	Registry ComponentRegistry // 提交后，这个字段会重置为nil

	NewFunc    func() any
	InjectFunc func(c InjectionExt, instance any) error
}

// ComponentRegistry 是组件注册
type ComponentRegistry interface {
	NewRegistration() *ComponentRegistration
	Register(r *ComponentRegistration) error
}

// ComponentRegistryFunc 是组件的注册函数
type ComponentRegistryFunc func(r ComponentRegistry) error

////////////////////////////////////////////////////////////////////////////////

// Commit 把注册信息提交到 ComponentRegistry
func (inst *ComponentRegistration) Commit() error {
	r := inst.Registry
	inst.Registry = nil
	if r == nil {
		return nil
	}
	return r.Register(inst)
}
