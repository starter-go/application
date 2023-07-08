package parameters

import (
	"fmt"
	"strings"

	"github.com/starter-go/base/safe"
)

// Table ... 表示一个参数表
type Table interface {
	Mode() safe.Mode

	Names() []string

	GetParam(name string) string
	GetParamOptional(name, defaultValue string) string
	GetParamRequired(name string) (string, error)

	SetParam(name, value string)

	Export(dst map[string]string) map[string]string
	Import(src map[string]string)
}

////////////////////////////////////////////////////////////////////////////////

type table struct {
	t    map[string]string
	mode safe.Mode
	lock safe.Lock
}

func (inst *table) GetParam(name string) string {
	inst.lock.Lock()
	defer inst.lock.Unlock()
	value := inst.t[name]
	return value
}

func (inst *table) Mode() safe.Mode {
	return inst.mode
}

func (inst *table) Names() []string {

	inst.lock.Lock()
	defer inst.lock.Unlock()

	src := inst.t
	dst := make([]string, 0)
	for key := range src {
		dst = append(dst, key)
	}

	return dst
}

func (inst *table) param(name string) string {
	name = inst.normalizeName(name)
	inst.lock.Lock()
	defer inst.lock.Unlock()
	return inst.t[name]
}

func (inst *table) GetParamOptional(name, defaultValue string) string {
	value := inst.param(name)
	if value == "" {
		value = defaultValue
	}
	return value
}

func (inst *table) GetParamRequired(name string) (string, error) {
	value := inst.param(name)
	if value == "" {
		return "", fmt.Errorf("no parameter with name: '%s'", name)
	}
	return value, nil
}

func (inst *table) SetParam(name, value string) {
	name = inst.normalizeName(name)
	inst.lock.Lock()
	defer inst.lock.Unlock()
	inst.t[name] = value
}

func (inst *table) normalizeName(name string) string {
	name = strings.TrimSpace(name)
	return name
}

func (inst *table) Export(dst map[string]string) map[string]string {
	inst.lock.Lock()
	defer inst.lock.Unlock()
	if dst == nil {
		dst = make(map[string]string)
	}
	src := inst.t
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func (inst *table) Import(src map[string]string) {
	inst.lock.Lock()
	defer inst.lock.Unlock()
	dst := inst.t
	for k, v := range src {
		dst[k] = v
	}
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
