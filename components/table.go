package components

// Table 表示一个组件的集合
type Table interface {
	Get(id ID) (Holder, error)
	Put(h Holder) error
	Select(selector Selector) ([]Holder, error)
	ListIDs() []ID

	Export(dst map[ID]Holder) map[ID]Holder
	Import(src map[ID]Holder)
}
