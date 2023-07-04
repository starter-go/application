package boot

import (
	"github.com/starter-go/application"
	"github.com/starter-go/base/util"
	"github.com/starter-go/vlog"
)

type modulesLoader struct {
	b    *Bootstrap
	mods map[string]*moduleHolder
}

func (inst *modulesLoader) load() error {
	inst.mods = make(map[string]*moduleHolder)
	m := inst.b.main
	err := inst.loadModuleR(m, 0)
	if err != nil {
		return err
	}
	res := inst.makeResult()
	inst.b.modules = res
	inst.log()
	return nil
}

func (inst *modulesLoader) log() {
	all := inst.b.modules
	for _, m := range all {
		name := m.Name()
		ver := m.Version()
		rev := m.Revision()
		vlog.Info("use module %s@%s-r%d", name, ver, rev)
	}
}

func (inst *modulesLoader) sort(list []*moduleHolder) {
	sorter := util.Sorter{
		OnLen:  func() int { return len(list) },
		OnLess: func(i1, i2 int) bool { return list[i1].depthSum > list[i2].depthSum },
		OnSwap: func(i1, i2 int) { list[i1], list[i2] = list[i2], list[i1] },
	}
	sorter.Sort()
}

func (inst *modulesLoader) makeResult() []application.Module {

	dst := make([]application.Module, 0)
	list := make([]*moduleHolder, 0)
	src := inst.mods

	for _, h := range src {
		list = append(list, h)
	}

	inst.sort(list)

	for _, h := range list {
		com := h.getLatest()
		if com != nil {
			dst = append(dst, com)
		}
	}

	return dst
}

func (inst *modulesLoader) loadModuleR(m application.Module, depth int) error {

	// for self
	if m == nil {
		return nil
	}

	h := inst.holderForModule(m)
	h.depthSum += depth
	h.add(m)

	// for deps
	deplist := m.Dependencies()
	for _, dep := range deplist {
		err := inst.loadModuleR(dep, depth+1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *modulesLoader) keyForModule(m application.Module) string {
	name := m.Name()
	return name
}

func (inst *modulesLoader) holderForModule(m application.Module) *moduleHolder {
	key := inst.keyForModule(m)
	h := inst.mods[key]
	if h == nil {
		h = &moduleHolder{key: key}
		h.init()
		inst.mods[key] = h
	}
	return h
}

////////////////////////////////////////////////////////////////////////////////

type moduleHolder struct {
	versions map[string]application.Module
	key      string
	depthSum int
}

func (inst *moduleHolder) init() {
	inst.versions = make(map[string]application.Module)
}

func (inst *moduleHolder) add(m application.Module) {
	ver := m.Version()
	inst.versions[ver] = m
}

func (inst *moduleHolder) getLatest() application.Module {
	all := inst.versions
	var a application.Module = nil
	for _, b := range all {
		if a == nil {
			a = b
			continue
		}
		if a.Revision() < b.Revision() {
			a = b
		}
	}
	return a
}
