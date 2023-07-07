package implcom

import (
	"fmt"

	"github.com/starter-go/application"
	"github.com/starter-go/application/components"
	"github.com/starter-go/base/safe"
)

////////////////////////////////////////////////////////////////////////////////

type comTable struct {
	holders    map[components.ID]components.Holder   // map[id]holder
	selections map[components.Selector]*comSelection // map[selector]selection
}

func (inst *comTable) _Impl() components.Table {
	return inst
}

func (inst *comTable) init() {
	inst.holders = make(map[components.ID]components.Holder)
}

func (inst *comTable) Get(id components.ID) (components.Holder, error) {
	tab := inst.holders
	h := tab[id]
	if h == nil {
		return nil, fmt.Errorf("no component with id:%s", id)
	}
	return h, nil
}

func (inst *comTable) Put(h components.Holder) error {

	if h == nil {
		return nil
	}

	tab := inst.holders
	id := h.Info().ID()
	older := tab[id]

	if older != nil {
		return fmt.Errorf("the components are duplicated, id:%s", id)
	}

	tab[id] = h
	inst.selections = nil
	return nil
}

func (inst *comTable) Select(selector components.Selector) ([]components.Holder, error) {

	selections := inst.selections
	if selections == nil {
		s2, err := inst.loadSelections()
		if err != nil {
			return nil, err
		}
		selections = s2
		inst.selections = selections
	}

	selection := selections[selector]
	if selection == nil {
		return nil, fmt.Errorf("no component found with selector:%s", selector)
	}

	list := selection.listHolders()
	return list, nil
}

func (inst *comTable) loadSelections() (map[components.Selector]*comSelection, error) {
	builder := &comSelectionTableBuilder{}
	builder.init()
	src := inst.holders
	for _, h := range src {
		err := builder.put(h)
		if err != nil {
			return nil, err
		}
	}
	return builder.create()
}

func (inst *comTable) ListIDs() []components.ID {
	src := inst.holders
	dst := make([]components.ID, 0)
	for id := range src {
		dst = append(dst, id)
	}
	return dst
}

////////////////////////////////////////////////////////////////////////////////

type componentTableBuilder struct {
	holders []components.Holder
}

func (inst *componentTableBuilder) _impl() application.ComponentRegistry {
	return inst
}

func (inst *componentTableBuilder) NewRegistration() *application.ComponentRegistration {
	r := inst._impl()
	return &application.ComponentRegistration{
		Registry: r,
	}
}

func (inst *componentTableBuilder) Register(src *application.ComponentRegistration) error {

	rn := &registrationNormalizer{r: src}
	id := rn.GetID()
	factory := rn.makeFactory()

	info := &comInfo{}
	info.id = id
	info.aliases = rn.GetAliases(id)
	info.classes = rn.GetClasses()
	info.scope = rn.GetScope()

	h := &comHolder{}
	h.init(info, factory)

	inst.holders = append(inst.holders, h)

	return nil
}

func (inst *componentTableBuilder) Create(mode safe.Mode) components.Table {
	src := inst.holders
	dst := make(map[components.ID]components.Holder)
	for _, h := range src {
		id := h.Info().ID()
		dst[id] = h
	}
	tab := &comTable{}
	tab.holders = dst
	return tab
}

////////////////////////////////////////////////////////////////////////////////
