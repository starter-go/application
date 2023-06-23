package environment

type Table interface {
	Env(name string) string
}
