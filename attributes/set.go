package attributes

type Table interface {
	Attr(name string) string
}
