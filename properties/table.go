package properties

import (
	"fmt"
	"strings"

	"github.com/starter-go/base/safe"
)

// Table 表示一个属性的集合
type Table interface {
	Mode() safe.Mode

	Names() []string

	GetProperty(name string) string
	GetPropertyRequired(name string) (string, error)
	GetPropertyOptional(name string, defaultValue string) string

	SetProperty(name string, value string)

	Export(dst map[string]string) map[string]string

	Import(src map[string]string)
}

////////////////////////////////////////////////////////////////////////////////

type table struct {
	t    map[string]string
	mode safe.Mode
	lock safe.Lock
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

func (inst *table) property(name string) string {
	name = inst.normalizeName(name)
	inst.lock.Lock()
	defer inst.lock.Unlock()
	return inst.t[name]
}

func (inst *table) GetProperty(name string) string {
	return inst.property(name)
}

func (inst *table) GetPropertyOptional(name string, defaultValue string) string {
	value := inst.property(name)
	if value == "" {
		value = defaultValue
	}
	return value
}

func (inst *table) GetPropertyRequired(name string) (string, error) {
	value := inst.property(name)
	if value == "" {
		return "", fmt.Errorf("no property with name: '%s'", name)
	}
	return value, nil
}

func (inst *table) SetProperty(name string, value string) {
	name = inst.normalizeName(name)
	inst.lock.Lock()
	defer inst.lock.Unlock()
	inst.t[name] = value
}

func (inst *table) Export(dst map[string]string) map[string]string {

	if dst == nil {
		dst = make(map[string]string)
	}

	inst.lock.Lock()
	defer inst.lock.Unlock()

	src := inst.t
	inst.copyItems(dst, src)
	return dst
}

func (inst *table) Import(src map[string]string) {
	inst.lock.Lock()
	defer inst.lock.Unlock()
	dst := inst.t
	inst.copyItems(dst, src)
}

func (inst *table) normalizeName(name string) string {
	name = strings.TrimSpace(name)
	return name
}

func (inst *table) copyItems(dst, src map[string]string) {
	if dst == nil || src == nil {
		return
	}
	for k, v := range src {
		if k == "" || v == "" {
			continue
		}
		dst[k] = v
	}
}

func (inst *table) Mode() safe.Mode {
	return inst.mode
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
