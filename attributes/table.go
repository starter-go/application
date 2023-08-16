package attributes

import (
	"fmt"
	"strings"

	"github.com/starter-go/base/safe"
)

// Table 表示 attributes 属性表
type Table interface {
	Mode() safe.Mode

	Names() []string

	GetAttribute(name string) any
	GetAttributeOptional(name string, defaultValue any) any
	GetAttributeRequired(name string) (any, error)

	SetAttribute(name string, value any)

	Export(dst map[string]any) map[string]any
	Import(src map[string]any)
}

////////////////////////////////////////////////////////////////////////////////

type table struct {
	t    map[string]any
	mode safe.Mode
	lock safe.Lock
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

func (inst *table) normalizeName(name string) string {
	name = strings.TrimSpace(name)
	return name
}

func (inst *table) attr(name string) any {
	name = inst.normalizeName(name)
	inst.lock.Lock()
	defer inst.lock.Unlock()
	return inst.t[name]
}

func (inst *table) GetAttribute(name string) any {
	return inst.attr(name)
}

func (inst *table) GetAttributeOptional(name string, defaultValue any) any {
	value := inst.attr(name)
	if value == "" {
		value = defaultValue
	}
	return value
}

func (inst *table) GetAttributeRequired(name string) (any, error) {
	value := inst.attr(name)
	if value == "" {
		return "", fmt.Errorf("no attribute with name: '%s'", name)
	}
	return value, nil
}

func (inst *table) SetAttribute(name string, value any) {
	name = inst.normalizeName(name)
	inst.lock.Lock()
	defer inst.lock.Unlock()
	inst.t[name] = value
}

func (inst *table) Export(dst map[string]any) map[string]any {
	inst.lock.Lock()
	defer inst.lock.Unlock()
	if dst == nil {
		dst = make(map[string]any)
	}
	src := inst.t
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func (inst *table) Import(src map[string]any) {
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
	t := make(map[string]any)
	return &table{
		t:    t,
		mode: mode,
		lock: mode.NewLock(),
	}
}

////////////////////////////////////////////////////////////////////////////////
