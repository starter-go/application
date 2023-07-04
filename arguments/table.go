package arguments

import "github.com/starter-go/base/safe"

type Table interface {
	Raw() []string
}

////////////////////////////////////////////////////////////////////////////////

type table struct {
	args []string
	mode safe.Mode
	lock safe.Lock
}

func (inst *table) Raw() []string {
	list := inst.args
	if list == nil {
		list = make([]string, 0)
	}
	return list
}

////////////////////////////////////////////////////////////////////////////////

// NewTable 新建属性表
func NewTable(args []string, mode safe.Mode) Table {
	if mode == nil {
		mode = safe.Default()
	}
	if args == nil {
		args = make([]string, 0)
	}
	return &table{
		args: args,
		mode: mode,
		lock: mode.NewLock(),
	}
}

////////////////////////////////////////////////////////////////////////////////
