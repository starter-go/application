package implcom

import (
	"fmt"

	"github.com/starter-go/application"
	"github.com/starter-go/application/components"
)

type comFactory struct {
	r *application.ComponentRegistration
}

func (inst *comFactory) _Impl() components.Factory {
	return inst
}

func (inst *comFactory) CreateInstance(info components.Info) (components.Instance, error) {

	reg := inst.r
	fn := reg.NewFunc
	if fn == nil {
		return nil, fmt.Errorf("no NewFunc for the component, id=%s", info.ID())
	}
	target := fn()

	ci := &comInstance{}
	ci.info = info
	ci.target = target
	ci.reg = reg
	return ci, nil
}
