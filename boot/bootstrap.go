package boot

import (
	"github.com/starter-go/application"
	"github.com/starter-go/application/arguments"
	"github.com/starter-go/base/safe"
	"github.com/starter-go/base/safe/ab"
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
	mode        ab.Mode
	profile     string
	args        []string
	modules     []application.Module
	main        application.Module
	context     application.Context
	collections application.Collections
}

func (inst *Bootstrap) init(opt *Options) {

	if opt == nil {
		opt = &Options{}
	}

	mode0 := safe.Fast()
	mode1 := opt.Mode
	args := opt.Args

	if args == nil {
		args = make([]string, 0)
	}

	if mode1 == nil {
		mode1 = safe.Fast()
	}

	mode := ab.New(mode0, mode1)
	mode.UseModeA()

	inst.args = args
	inst.mode = mode

	col := &application.Collections{}
	col.Arguments = arguments.NewTable(args, mode)
	col.Attributes = opt.Attributes
	col.Environment = opt.Environment
	col.Parameters = opt.Parameters
	col.Properties = opt.Properties
	col.Complete(mode)
	col = col.Clone(mode)
	inst.collections = *col
}

// Run 运行
func (inst *Bootstrap) Run(m application.Module) error {
	inst.main = m
	steps := make([]func() error, 0)

	steps = append(steps, inst.loadModules)
	steps = append(steps, inst.loadResources)

	steps = append(steps, inst.loadProperties)
	steps = append(steps, inst.loadParameters)
	steps = append(steps, inst.loadEnvironment)
	steps = append(steps, inst.loadAttributes)

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

func (inst *Bootstrap) loadEnvironment() error {
	loader := &envLoader{b: inst}
	return loader.Load()
}

func (inst *Bootstrap) loadParameters() error {
	loader := &parametersLoader{b: inst}
	return loader.Load()
}

func (inst *Bootstrap) loadAttributes() error {
	loader := &attributesLoader{b: inst}
	return loader.Load()
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

	inst.mode.UseModeB()

	lm := inst.context.GetLifeManager()
	runner := &mainLoopRunner{lm: lm}
	return runner.Run()
}
