package application

import (
	"embed"

	"github.com/starter-go/application/resources"
)

// ModuleBuilder 是用来创建模块的工具
type ModuleBuilder struct {
	name     string
	version  string
	revision int
	res      resources.Table
	deps     []Module
	registry ComponentRegistryFunc
}

// Name 方法用来设置模块名称
func (inst *ModuleBuilder) Name(name string) *ModuleBuilder {
	inst.name = name
	return inst
}

// Version 方法用来设置模块的版本号
func (inst *ModuleBuilder) Version(version string) *ModuleBuilder {
	inst.version = version
	return inst
}

// Revision 方法用来设置模块的版本编号
func (inst *ModuleBuilder) Revision(rev int) *ModuleBuilder {
	inst.revision = rev
	return inst
}

// Depend 方法用来添加依赖项
func (inst *ModuleBuilder) Depend(deps ...Module) *ModuleBuilder {
	inst.deps = append(inst.deps, deps...)
	return inst
}

// EmbedResources 方法用来添加嵌入的资源
func (inst *ModuleBuilder) EmbedResources(fs embed.FS, basepath string) *ModuleBuilder {
	inst.res = resources.NewEmbed(fs, basepath)
	return inst
}

// Components 用于设置组件注册函数
// func (inst *ModuleBuilder) Components(fn components.RegistryHandlerFunc) *ModuleBuilder {
// 	inst.comRegFn = fn
// 	return inst
// }

// Components 用于设置组件注册函数
func (inst *ModuleBuilder) Components(fn ComponentRegistryFunc) *ModuleBuilder {
	inst.registry = fn
	return inst
}

// Create 创建模块
func (inst *ModuleBuilder) Create() Module {

	m := &myModule{}

	m.name = inst.name
	m.version = inst.version
	m.revision = inst.revision

	m.deps = inst.deps
	m.res = inst.res
	m.registry = inst.registry

	if m.name == "" {
		m.name = "unnamed"
	}

	if m.version == "" {
		m.version = "0.0.0"
	}

	return m
}
