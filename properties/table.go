package properties

// Table 表示一个属性的集合
type Table interface {
	Names() []string

	GetPropertyOptional(name string, defaultValue string) string

	GetProperty(name string) (string, error)

	SetProperty(name string, value string)
}
