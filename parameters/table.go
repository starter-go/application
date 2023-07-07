package parameters

import "github.com/starter-go/base/safe"

// Table ... 表示一个参数表
type Table interface {
	Param(name string) string
}

////////////////////////////////////////////////////////////////////////////////

type table struct {
	t    map[string]string
	mode safe.Mode
	lock safe.Lock
}

func (inst *table) Param(name string) string {
	inst.lock.Lock()
	defer inst.lock.Unlock()
	value := inst.t[name]
	return value
}

////////////////////////////////////////////////////////////////////////////////

// NewTable 新建属性表
func NewTable(mode safe.Mode) Table {
	if mode == nil {
		mode = safe.Default()
	}
	t := make(map[string]string)
	return &table{
		t:    t,
		mode: mode,
		lock: mode.NewLock(),
	}
}

////////////////////////////////////////////////////////////////////////////////