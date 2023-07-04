package attributes

import "github.com/starter-go/base/safe"

type Table interface {
	Attr(name string) any
}

////////////////////////////////////////////////////////////////////////////////

type table struct {
	t    map[string]any
	mode safe.Mode
	lock safe.Lock
}

func (inst *table) Attr(name string) any {

	inst.lock.Lock()
	defer inst.lock.Lock()

	value := inst.t[name]
	return value
}

////////////////////////////////////////////////////////////////////////////////

// NewTable 新建属性表
func NewTable(mode safe.Mode) Table {
	if mode == nil {
		mode = safe.Default()
	}
	t := make(map[string]any)
	return &table{
		t:    t,
		mode: mode,
		lock: mode.NewLock(),
	}
}

////////////////////////////////////////////////////////////////////////////////
