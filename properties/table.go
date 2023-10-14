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

	Exists(name string) bool

	Clear()

	Remove(name string) bool

	GetProperty(name string) string
	GetPropertyRequired(name string) (string, error)
	GetPropertyOptional(name string, defaultValue string) string

	SetProperty(name string, value string)

	Export(dst map[string]string) map[string]string

	Import(src map[string]string)

	Getter() Getter
	Setter() Setter
}

////////////////////////////////////////////////////////////////////////////////

type property struct {
	name  string
	value string
}

////////////////////////////////////////////////////////////////////////////////

type table struct {
	t    map[string]*property
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

func (inst *table) property(name string) (*property, error) {
	name = inst.normalizeName(name)
	inst.lock.Lock()
	defer inst.lock.Unlock()
	item := inst.t[name]
	if item == nil {
		return nil, fmt.Errorf("no property with name: '%s'", name)
	}
	return item, nil
}

func (inst *table) GetProperty(name string) string {
	item, err := inst.property(name)
	if err != nil {
		return ""
	}
	return item.value
}

func (inst *table) GetPropertyOptional(name string, defaultValue string) string {
	item, _ := inst.property(name)
	if item == nil {
		return defaultValue
	}
	return item.value
}

func (inst *table) GetPropertyRequired(name string) (string, error) {
	item, err := inst.property(name)
	if item != nil {
		return item.value, nil
	}
	return "", err
}

func (inst *table) SetProperty(name string, value string) {
	name = inst.normalizeName(name)
	inst.lock.Lock()
	defer inst.lock.Unlock()
	item := &property{}
	item.name = name
	item.value = value
	inst.t[name] = item
}

func (inst *table) Export(dst map[string]string) map[string]string {

	inst.lock.Lock()
	defer inst.lock.Unlock()

	src := inst.t
	if dst == nil {
		dst = make(map[string]string)
	}

	for name, item := range src {
		if item != nil {
			dst[name] = item.value
		}
	}
	return dst
}

func (inst *table) Import(src map[string]string) {

	inst.lock.Lock()
	defer inst.lock.Unlock()

	dst := inst.t
	if dst == nil {
		dst = make(map[string]*property)
	}
	for name, value := range src {
		item := &property{}
		item.name = name
		item.value = value
		dst[name] = item
	}
	inst.t = dst
}

func (inst *table) normalizeName(name string) string {
	name = strings.TrimSpace(name)
	return name
}

// func (inst *table) copyItems(dst, src map[string]string) {
// 	if dst == nil || src == nil {
// 		return
// 	}
// 	for k, v := range src {
// 		if k == "" || v == "" {
// 			continue
// 		}
// 		dst[k] = v
// 	}
// }

func (inst *table) Mode() safe.Mode {
	return inst.mode
}

func (inst *table) Clear() {
	t := make(map[string]*property)
	inst.t = t
}

func (inst *table) Exists(name string) bool {
	inst.lock.Lock()
	defer inst.lock.Unlock()
	item := inst.t[name]
	return (item != nil)
}

func (inst *table) Remove(name string) bool {
	inst.lock.Lock()
	defer inst.lock.Unlock()
	item := inst.t[name]
	if item == nil {
		return false
	}
	inst.t[name] = nil
	return true
}

func (inst *table) Getter() Getter {
	return &myGetter{table: inst}
}

func (inst *table) Setter() Setter {
	return &mySetter{table: inst}
}

////////////////////////////////////////////////////////////////////////////////

// NewTable 新建属性表
func NewTable(mode safe.Mode) Table {
	if mode == nil {
		mode = safe.Default()
	}
	t := make(map[string]*property)
	return &table{
		t:    t,
		mode: mode,
		lock: mode.NewLock(),
	}
}
