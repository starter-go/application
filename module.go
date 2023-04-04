package application

// Module 表示一个 Starter 模块
type Module interface {
	Name() string
	Version() string
	Revision() int
	Dependencies() []Module

	// Resources() collection.Resources
	// Apply(cb ConfigBuilder) error

}
