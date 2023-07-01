package parameters

// Table ... 表示一个参数表
type Table interface {
	Param(name string) string
}
