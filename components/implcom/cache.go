package implcom

import "github.com/starter-go/application/components"

type cache interface {
	GetInstance(i components.Info, f components.Factory) (components.Instance, error)
}

////////////////////////////////////////////////////////////////////////////////

type singletonCache struct {
	instance components.Instance
}

func (inst *singletonCache) GetInstance(i components.Info, f components.Factory) (components.Instance, error) {
	o := inst.instance
	if o != nil {
		return o, nil
	}
	o, err := f.CreateInstance(i)
	if err != nil {
		return nil, err
	}
	inst.instance = o
	return o, nil
}

////////////////////////////////////////////////////////////////////////////////

type prototypeCache struct{}

func (inst *prototypeCache) GetInstance(i components.Info, f components.Factory) (components.Instance, error) {
	return f.CreateInstance(i)
}

////////////////////////////////////////////////////////////////////////////////
