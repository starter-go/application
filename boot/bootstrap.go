package boot

import (
	"github.com/starter-go/application"
	"github.com/starter-go/application/arguments"
)

// Run 运行指定的模块
func Run(m application.Module, opt *Options) error {
	b := &Bootstrap{}
	b.init(opt)
	return b.Run(m)
}

////////////////////////////////////////////////////////////////////////////////

// Bootstrap 是 starter 的应用启动器
type Bootstrap struct {
	profile     string
	main        application.Module
	context     application.Context
	collections application.Collections
	Options     Options
	modules     []application.Module
}

func (inst *Bootstrap) init(opt *Options) {

	if opt == nil {
		opt = &Options{}
	}

	args := opt.Args

	inst.Options = *opt
	inst.collections.Complete(nil)
	inst.collections.Arguments = arguments.NewTable(args, nil)
}

// Run 运行
func (inst *Bootstrap) Run(m application.Module) error {
	inst.main = m
	steps := make([]func() error, 0)

	steps = append(steps, inst.loadModules)
	steps = append(steps, inst.loadResources)
	steps = append(steps, inst.loadProperties)
	steps = append(steps, inst.loadContext)
	steps = append(steps, inst.runMainLoop)

	for _, step := range steps {
		err := step()
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *Bootstrap) loadModules() error {
	loader := &modulesLoader{b: inst}
	return loader.load()
}

func (inst *Bootstrap) loadProperties() error {
	loader := &propertiesLoader{b: inst}
	return loader.load()
}

func (inst *Bootstrap) loadResources() error {
	loader := &resourcesLoader{b: inst}
	return loader.load()
}

func (inst *Bootstrap) loadContext() error {
	loader := &contextLoader{b: inst}
	return loader.load()
}

func (inst *Bootstrap) runMainLoop() error {
	lm := inst.context.GetLifeManager()
	runner := &mainLoopRunner{lm: lm}
	return runner.Run()
}
