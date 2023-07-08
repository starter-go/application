package resources

import (
	"fmt"

	"github.com/starter-go/base/safe"
)

// Table 表示一个资源的集合
type Table interface {
	Paths() []string

	GetResource(path string) (Resource, error)

	Export(dst map[string]Resource) map[string]Resource
	Import(src map[string]Resource)
}

////////////////////////////////////////////////////////////////////////////////

type table struct {
	t    map[string]Resource
	mode safe.Mode
	lock safe.Lock
}

func (inst *table) Paths() []string {
	src := inst.t
	dst := make([]string, 0)
	inst.lock.Lock()
	defer inst.lock.Unlock()
	for k, v := range src {
		if v != nil {
			dst = append(dst, k)
		}
	}
	return dst
}

func (inst *table) GetResource(path string) (Resource, error) {
	path = normalizePath(path)
	inst.lock.Lock()
	defer inst.lock.Unlock()
	res := inst.t[path]
	if res == nil {
		return nil, fmt.Errorf("no resource with path [%s]", path)
	}
	return res, nil
}

func (inst *table) Export(dst map[string]Resource) map[string]Resource {
	inst.lock.Lock()
	defer inst.lock.Unlock()
	if dst == nil {
		dst = make(map[string]Resource)
	}
	src := inst.t
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func (inst *table) Import(src map[string]Resource) {
	inst.lock.Lock()
	defer inst.lock.Unlock()
	dst := inst.t
	for k, v := range src {
		dst[k] = v
	}
}

////////////////////////////////////////////////////////////////////////////////

// NewTable 新建属性表
func NewTable(src map[string]Resource, mode safe.Mode) Table {
	if mode == nil {
		mode = safe.Default()
	}
	dst := make(map[string]Resource)
	for _, item := range src {
		path := item.Path()
		dst[path] = item
	}
	return &table{
		t:    dst,
		mode: mode,
		lock: mode.NewLock(),
	}
}

////////////////////////////////////////////////////////////////////////////////
