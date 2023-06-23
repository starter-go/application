package resources

// Table 表示一个资源的集合
type Table interface {
	Paths() []string

	GetResource(path string) (Resource, error)
}
