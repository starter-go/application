package boot

import (
	"github.com/starter-go/application"
)

// Run 运行指定的模块
func Run(m application.Module) error {
	b := &Bootstrap{}
	return b.Run(m)
}

////////////////////////////////////////////////////////////////////////////////

// Bootstrap 是 starter 的应用启动器
type Bootstrap struct {
	profile string
	main    application.Module
	context application.Context
	modules []application.Module
}

// Run 运行
func (inst *Bootstrap) Run(m application.Module) error {
	inst.main = m
	steps := make([]func() error, 0)

	steps = append(steps, inst.loadModules)
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

func (inst *Bootstrap) loadContext() error {
	loader := &contextLoader{b: inst}
	return loader.load()
}

func (inst *Bootstrap) runMainLoop() error {
	lm := inst.context.GetLifeManager()
	runner := &mainLoopRunner{lm: lm}
	return runner.Run()
}
