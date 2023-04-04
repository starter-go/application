package application

// ModuleBuilder 是用来创建模块的工具
type ModuleBuilder struct {
	name     string
	version  string
	revision int
	deps     []Module
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

// Create 创建模块
func (inst *ModuleBuilder) Create() Module {
	return nil
}
