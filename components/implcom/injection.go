package implcom

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/starter-go/application"
	"github.com/starter-go/application/components"
)

type injection struct {
	scope components.Scope

	parent application.Context
	lm     application.LifeManager
	ext    application.InjectionExt

	table map[components.ID]*hInstance
}

func (inst *injection) _Impl() application.Injection {
	return inst
}

func (inst *injection) init(parent application.Context, scope components.Scope) *injection {

	lm := &lifeManager{}
	lm.init()

	inst.parent = parent
	inst.scope = scope
	inst.table = make(map[components.ID]*hInstance)
	inst.lm = lm
	return inst
}

func (inst *injection) Deadline() (deadline time.Time, ok bool) {
	now := time.Now()
	return now, false
}

func (inst *injection) Done() <-chan struct{} {
	return nil
}

func (inst *injection) Err() error {
	return nil
}

func (inst *injection) Value(key any) any {
	return nil
}

func (inst *injection) Scope() components.Scope {
	return inst.scope
}

func (inst *injection) Parent() application.Context {
	return inst.parent
}

func (inst *injection) LifeManager() application.LifeManager {
	return inst.lm
}

func (inst *injection) Select(selector components.Selector) ([]components.Instance, error) {
	dst := make([]components.Instance, 0)
	holders, err := inst.parent.GetComponents().Select(selector)
	if err != nil {
		return nil, err
	}
	for _, h := range holders {
		instance, err := inst.GetWithHolder(h)
		if err != nil {
			return nil, err
		}
		dst = append(dst, instance)
	}
	return dst, nil
}

func (inst *injection) SelectOne(selector components.Selector) (components.Instance, error) {

	holders, err := inst.parent.GetComponents().Select(selector)
	if err != nil {
		return nil, err
	}

	if holders == nil {
		return nil, fmt.Errorf("no component with selector: %s", selector)
	}
	count := len(holders)
	if count < 1 {
		return nil, fmt.Errorf("no component with selector: %s", selector)
	}
	if count > 1 {
		return nil, fmt.Errorf("%d components with selector: %s", count, selector)
	}
	holder := holders[0]
	if holder == nil {
		return nil, fmt.Errorf("holder is nil for component selector: %s", selector)
	}

	return inst.GetWithHolder(holder)
}

func (inst *injection) GetByID(id components.ID) (components.Instance, error) {
	hi := inst.table[id]
	if hi == nil {
		holder, err := inst.parent.GetComponents().Get(id)
		if err != nil {
			return nil, err
		}
		return inst.GetWithHolder(holder)
	}
	err := hi.check()
	if err != nil {
		return nil, err
	}
	return hi.instance, nil
}

func (inst *injection) GetWithHolder(holder components.Holder) (components.Instance, error) {
	if holder == nil {
		panic("holder is nil")
	}
	id := holder.Info().ID()
	hi := inst.table[id]
	if hi == nil {
		hi = &hInstance{}
		err := hi.init(holder)
		if err != nil {
			return nil, err
		}
		inst.table[id] = hi
	}
	err := hi.check()
	if err != nil {
		return nil, err
	}
	return hi.instance, nil
}

func (inst *injection) Complete() error {
	for {
		count, err := inst.tryCompleteOnce()
		if err != nil {
			return err
		}
		if count <= 0 {
			return nil
		}
	}
}

func (inst *injection) tryCompleteOnce() (int, error) {
	count := 0
	tab := inst.table
	ids := make([]components.ID, 0)
	for id := range tab {
		ids = append(ids, id)
	}
	for _, id := range ids {
		hi := tab[id]
		if hi.injected {
			continue
		}
		hi.injected = true
		count++
		err := hi.instance.Inject(inst)
		if err != nil {
			return 0, err
		}
	}
	return count, nil
}

func (inst *injection) GetProperty(selector components.Selector) (string, error) {
	const (
		prefix = "${"
		suffix = "}"
	)
	name := selector.String()
	name = strings.TrimSpace(name)
	if strings.HasPrefix(name, prefix) && strings.HasSuffix(name, suffix) {
		name = name[len(prefix) : len(name)-len(suffix)]
		name = strings.TrimSpace(name)
		return inst.parent.GetProperties().GetProperty(name)
	}
	return name, nil
}

func (inst *injection) Ext() application.InjectionExt {
	ext := inst.ext
	if ext == nil {
		ie := &injectionExt{injection: inst}
		ext = ie.init()
		inst.ext = ext
	}
	return ext
}

func (inst *injection) GetContext() application.Context {
	return inst.parent
}

func (inst *injection) GetApplicationContext() context.Context {
	return inst.parent
}

////////////////////////////////////////////////////////////////////////////////

type hInstance struct {
	id       components.ID
	instance components.Instance
	holder   components.Holder
	injected bool
}

func (inst *hInstance) init(holder components.Holder) error {
	newInstance, err := holder.GetInstance()
	if err != nil {
		return err
	}
	inst.id = holder.Info().ID()
	inst.holder = holder
	inst.instance = newInstance
	inst.injected = false
	return nil
}

func (inst *hInstance) check() error {
	if inst.instance == nil {
		return fmt.Errorf("hInstance.instance is nil")
	}
	if inst.holder == nil {
		return fmt.Errorf("hInstance.holder is nil")
	}
	return nil
}
