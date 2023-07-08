package boot

type resourcesLoader struct {
	b *Bootstrap
}

func (inst *resourcesLoader) load() error {
	src := inst.b.modules
	dst := inst.b.collections.Resources
	for _, m := range src {
		items := m.Resources()
		dst.Import(items.Export(nil))
	}
	return nil
}
