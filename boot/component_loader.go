package boot

import (
	"fmt"
	"strings"

	"github.com/starter-go/application/components"
)

type componentsLoader struct {
	b *Bootstrap
}

func (inst *componentsLoader) load() error {
	mods := inst.b.modules
	builder := &componentTableBuilder{}
	for _, m := range mods {
		err := m.RegisterComponents(builder)
		if err != nil {
			return err
		}
	}
	inst.b.components = builder.Create()
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type componentTableBuilder struct {
	holders []components.Holder
}

func (inst *componentTableBuilder) _Impl() components.Registry {
	return inst
}

func (inst *componentTableBuilder) New() *components.Registration {
	return &components.Registration{}
}

func (inst *componentTableBuilder) Register(src *components.Registration) {

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
}

func (inst *componentTableBuilder) Create() components.Table {
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

type registrationNormalizer struct {
	r *components.Registration
}

func (inst *registrationNormalizer) GetID() components.ID {
	r := inst.r
	idstr := r.ID.String()
	idstr = strings.TrimSpace(idstr)
	if idstr == "" {
		idstr = "no-id"
	}
	return components.ID(idstr)
}

func (inst *registrationNormalizer) GetAliases(cid components.ID) []string {
	id := cid.String()
	r := inst.r
	text := r.Aliases
	list := inst.parseStringArray(text, func(s string) bool {
		return (s != "" && s != id)
	})
	return list
}

func (inst *registrationNormalizer) GetClasses() []string {
	r := inst.r
	text := r.Classes
	list := inst.parseStringArray(text, func(s string) bool {
		return (s != "")
	})
	return list
}

func (inst *registrationNormalizer) parseStringArray(text string, accept func(string) bool) []string {
	const (
		space = string(' ')
		tab   = "\t"
		nl    = "\n"
	)
	text = strings.ReplaceAll(text, tab, nl)
	text = strings.ReplaceAll(text, space, nl)
	items := strings.Split(text, nl)
	dst := make([]string, 0)
	for _, item := range items {
		item = strings.TrimSpace(item)
		if accept(item) {
			dst = append(dst, item)
		}
	}
	return dst
}

func (inst *registrationNormalizer) GetScope() components.Scope {
	text := inst.r.Scope
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	if text == "prototype" {
		return components.ScopePrototype
	}
	return components.ScopeSingleton
}

func (inst *registrationNormalizer) makeFactory() components.Factory {
	factory := &comFactory{
		r: inst.r,
	}
	return factory
}

////////////////////////////////////////////////////////////////////////////////

type comInfo struct {
	id      components.ID
	scope   components.Scope
	aliases []string
	classes []string
}

func (inst *comInfo) _Impl() components.Info {
	return inst
}

func (inst *comInfo) ID() components.ID {
	return inst.id
}

func (inst *comInfo) Aliases() []string {
	return inst.aliases
}

func (inst *comInfo) Classes() []string {
	return inst.classes
}

func (inst *comInfo) Scope() components.Scope {
	return inst.scope
}

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

type comFactory struct {
	r *components.Registration
}

func (inst *comFactory) _Impl() components.Factory {
	return inst
}

func (inst *comFactory) CreateInstance(info components.Info) (components.Instance, error) {

	reg := inst.r
	fn := reg.NewFunc
	if fn == nil {
		return nil, fmt.Errorf("no NewFunc for the component, id=%s", info.ID())
	}
	target := fn()

	ci := &comInstance{}
	ci.info = info
	ci.target = target
	ci.reg = reg
	return ci, nil
}

////////////////////////////////////////////////////////////////////////////////

type comInstance struct {
	reg    *components.Registration
	info   components.Info
	target any
}

func (inst *comInstance) _Impl() components.Instance {
	return inst
}

func (inst *comInstance) Get() any {
	return inst.target
}

func (inst *comInstance) Info() components.Info {
	return inst.info
}

func (inst *comInstance) Inject(i components.Injection) error {

	fn := inst.reg.InjectFunc
	t := inst.target

	if fn == nil {
		return nil
	}

	return fn(i, t)
}

////////////////////////////////////////////////////////////////////////////////
