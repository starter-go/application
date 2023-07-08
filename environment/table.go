package environment

import (
	"fmt"
	"strings"

	"github.com/starter-go/base/safe"
)

// Table 表示环境变量表
type Table interface {
	Mode() safe.Mode

	Names() []string

	SetEnv(name, value string)

	GetEnv(name string) string
	GetEnvOptional(name, defaultValue string) string
	GetEnvRequired(name string) (string, error)

	Export(dst map[string]string) map[string]string
	Import(src map[string]string)
}

////////////////////////////////////////////////////////////////////////////////

type table struct {
	t    map[string]string
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

func (inst *table) SetEnv(name, value string) {
	name = inst.normalizeName(name)
	inst.lock.Lock()
	defer inst.lock.Unlock()
	inst.t[name] = value
}

func (inst *table) env(name string) string {
	name = inst.normalizeName(name)
	inst.lock.Lock()
	defer inst.lock.Unlock()
	return inst.t[name]
}

func (inst *table) GetEnv(name string) string {
	return inst.env(name)
}

func (inst *table) GetEnvOptional(name, defaultValue string) string {
	value := inst.env(name)
	if value == "" {
		value = defaultValue
	}
	return value
}

func (inst *table) GetEnvRequired(name string) (string, error) {
	value := inst.env(name)
	if value == "" {
		return "", fmt.Errorf("no env value with name: '%s'", name)
	}
	return value, nil
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
