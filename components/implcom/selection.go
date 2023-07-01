package implcom

import "github.com/starter-go/application/components"

////////////////////////////////////////////////////////////////////////////////

type comSelectionTableBuilder struct {
	table map[components.Selector]*comSelection
}

func (inst *comSelectionTableBuilder) init() {
	inst.table = make(map[components.Selector]*comSelection)
}

func (inst *comSelectionTableBuilder) create() (map[components.Selector]*comSelection, error) {
	t := inst.table
	if t == nil {
		t = make(map[components.Selector]*comSelection)
	}
	return t, nil
}

func (inst *comSelectionTableBuilder) addRef(ref components.Ref) error {
	tab := inst.table
	sel := ref.Selector()
	selection := tab[sel]
	if selection == nil {
		selection = &comSelection{
			selector: sel,
		}
		tab[sel] = selection
	}
	selection.items = append(selection.items, ref)
	return nil
}

func (inst *comSelectionTableBuilder) put(h components.Holder) error {

	info := h.Info()
	err := inst.putHolderWithID(h, info.ID())
	if err != nil {
		return err
	}

	err = inst.putHolderWithClasses(h, info.Classes())
	if err != nil {
		return err
	}

	return inst.putHolderWithAliases(h, info.Aliases())
}

func (inst *comSelectionTableBuilder) putHolderWithID(h components.Holder, id components.ID) error {
	sel := components.SelectorForID(id)
	ref := h.NewRef(sel)
	return inst.addRef(ref)
}

func (inst *comSelectionTableBuilder) putHolderWithAliases(h components.Holder, aliases []string) error {
	for _, alias := range aliases {
		sel := components.SelectorForAlias(alias)
		ref := h.NewRef(sel)
		err := inst.addRef(ref)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *comSelectionTableBuilder) putHolderWithClasses(h components.Holder, classes []string) error {
	for _, cname := range classes {
		sel := components.SelectorForClass(cname)
		ref := h.NewRef(sel)
		err := inst.addRef(ref)
		if err != nil {
			return err
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// comSelection 是一组具有相同 key 的 Ref
type comSelection struct {
	selector components.Selector
	items    []components.Ref
}

func (inst *comSelection) listHolders() []components.Holder {
	src := inst.items
	dst := make([]components.Holder, 0)
	for _, ref := range src {
		h := ref.Holder()
		dst = append(dst, h)
	}
	return dst
}

////////////////////////////////////////////////////////////////////////////////
