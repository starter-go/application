package boot

import (
	"github.com/starter-go/application"
	"github.com/starter-go/vlog"
)

type mainLoopRunner struct {
	lm application.LifeManager
}

func (inst *mainLoopRunner) handleError(err error) {
	if err == nil {
		return
	}
	vlog.Warn("Error: %v", err)
}

func (inst *mainLoopRunner) Run() error {

	life := inst.lm.GetMaster()

	err := life.OnCreate()
	if err != nil {
		return err
	}

	defer func() {
		err := life.OnDestroy()
		inst.handleError(err)
	}()

	return inst.run2(life)
}

func (inst *mainLoopRunner) run2(life *application.Life) error {

	err := life.OnStartPre()
	if err != nil {
		return err
	}

	err = life.OnStart()
	if err != nil {
		return err
	}

	err = life.OnStartPost()
	if err != nil {
		return err
	}

	defer func() {
		err = life.OnStopPre()
		inst.handleError(err)

		err = life.OnStop()
		inst.handleError(err)

		err = life.OnStopPost()
		inst.handleError(err)
	}()

	return inst.run3(life)
}

func (inst *mainLoopRunner) run3(life *application.Life) error {
	return life.OnLoop()
}
