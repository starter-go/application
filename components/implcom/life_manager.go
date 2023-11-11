package implcom

import (
	"fmt"
	"sort"

	"github.com/starter-go/application"
	"github.com/starter-go/vlog"
)

type lifeManager struct {
	items  []*application.Life
	master *application.Life
}

func (inst *lifeManager) _Impl() application.LifeManager {
	return inst
}

func (inst *lifeManager) init() {

}

func (inst *lifeManager) Add(l *application.Life) {
	if l != nil {
		inst.items = append(inst.items, l)
	}
}

func (inst *lifeManager) GetMaster() *application.Life {
	m := inst.master
	if m == nil {

		ml := &masterLife{}
		ml.init(inst.items)

		m = &application.Life{}
		m.OnCreate = ml.create
		m.OnStartPre = ml.start1
		m.OnStart = ml.start
		m.OnStartPost = ml.start3
		m.OnLoop = ml.loop
		m.OnStopPre = ml.stop1
		m.OnStop = ml.stop
		m.OnStopPost = ml.stop3
		m.OnDestroy = ml.destroy
		inst.master = m
	}
	return m
}

////////////////////////////////////////////////////////////////////////////////

type lifeWrapper struct {
	application.Life
	err      error
	createOk bool
	startOk  bool
}

func (inst *lifeWrapper) nop() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type masterLife struct {
	items []*lifeWrapper
}

func (inst *masterLife) init(src []*application.Life) {
	dst := inst.items
	for _, item1 := range src {
		if item1 == nil {
			continue
		}
		item2 := &lifeWrapper{}
		item2.Order = item1.Order
		item2.OnCreate = item1.OnCreate
		item2.OnStartPre = item1.OnStartPre
		item2.OnStart = item1.OnStart
		item2.OnStartPost = item1.OnStartPost
		item2.OnLoop = item1.OnLoop
		item2.OnStopPre = item1.OnStopPre
		item2.OnStop = item1.OnStop
		item2.OnStopPost = item1.OnStopPost
		item2.OnDestroy = item1.OnDestroy
		inst.prepareItem(item2)
		dst = append(dst, item2)
	}
	inst.items = dst
	inst.sort()
}

func (inst *masterLife) prepareItem(l *lifeWrapper) {

	if l.OnCreate == nil {
		l.OnCreate = l.nop
	}

	if l.OnStartPre == nil {
		l.OnStartPre = l.nop
	}
	if l.OnStart == nil {
		l.OnStart = l.nop
	}
	if l.OnStartPost == nil {
		l.OnStartPost = l.nop
	}

	if l.OnLoop == nil {
		l.OnLoop = l.nop
	}

	if l.OnStopPre == nil {
		l.OnStopPre = l.nop
	}
	if l.OnStop == nil {
		l.OnStop = l.nop
	}
	if l.OnStopPost == nil {
		l.OnStopPost = l.nop
	}

	if l.OnDestroy == nil {
		l.OnDestroy = l.nop
	}
}

func (inst *masterLife) handlePanic(p any, l *lifeWrapper) {
	if p == nil || l == nil {
		return
	}

	err, ok := p.(error)
	if ok && err != nil {
		l.err = err
		vlog.Error("Error: %v", err)
		return
	}

	msg, ok := p.(string)
	if ok {
		l.err = fmt.Errorf(msg)
		vlog.Error("Error: %s", msg)
		return
	}
}

func (inst *masterLife) invokeItem(l *lifeWrapper, fn func(l *lifeWrapper) error) error {
	if l == nil || fn == nil {
		return nil
	}
	defer func() {
		x := recover()
		inst.handlePanic(x, l)
	}()
	return fn(l)
}

func (inst *masterLife) invokeListOpen(fn func(l *lifeWrapper) error) error {
	list := inst.items
	size := len(list)
	for i := 0; i < size; i++ {
		item := list[i]
		err := inst.invokeItem(item, fn)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *masterLife) invokeListClose(fn func(l *lifeWrapper) error) error {
	list := inst.items
	size := len(list)
	for i := size - 1; i >= 0; i-- {
		item := list[i]
		err := inst.invokeItem(item, fn)
		if err != nil {
			vlog.Warn("Error: %v", err)
		}
	}
	return nil
}

func (inst *masterLife) create() error {
	return inst.invokeListOpen(func(l *lifeWrapper) error {
		err := l.OnCreate()
		if err == nil {
			l.createOk = true
		}
		return err
	})
}

func (inst *masterLife) start1() error {
	return inst.invokeListOpen(func(l *lifeWrapper) error {
		return l.OnStartPre()
	})
}
func (inst *masterLife) start() error {
	return inst.invokeListOpen(func(l *lifeWrapper) error {
		err := l.OnStart()
		if err == nil {
			l.startOk = true
		}
		return err
	})
}
func (inst *masterLife) start3() error {
	return inst.invokeListOpen(func(l *lifeWrapper) error {
		return l.OnStartPost()
	})
}

func (inst *masterLife) loop() error {
	return inst.invokeListOpen(func(l *lifeWrapper) error {
		return l.OnLoop()
	})
}

func (inst *masterLife) stop1() error {
	return inst.invokeListClose(func(l *lifeWrapper) error {
		return l.OnStopPre()
	})
}
func (inst *masterLife) stop() error {
	return inst.invokeListClose(func(l *lifeWrapper) error {
		if !l.startOk {
			return nil
		}
		return l.OnStop()
	})
}
func (inst *masterLife) stop3() error {
	return inst.invokeListClose(func(l *lifeWrapper) error {
		return l.OnStopPost()
	})
}

func (inst *masterLife) destroy() error {
	return inst.invokeListClose(func(l *lifeWrapper) error {
		if !l.createOk {
			return nil
		}
		return l.OnDestroy()
	})
}

func (inst *masterLife) Len() int {
	return len(inst.items)
}

func (inst *masterLife) Less(i1, i2 int) bool {
	list := inst.items
	return (list[i1].Order < list[i2].Order)
}

func (inst *masterLife) Swap(i1, i2 int) {
	list := inst.items
	list[i1], list[i2] = list[i2], list[i1]
}

func (inst *masterLife) sort() {
	sort.Sort(inst)
}
