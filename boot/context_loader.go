package boot

import "github.com/starter-go/application/components/implcom"

type contextLoader struct {
	b *Bootstrap
}

func (inst *contextLoader) load() error {

	boot := inst.b
	mods := boot.modules
	builder := implcom.NewBuilder()
	builder.SetCollections(&boot.collections)

	for _, m := range mods {
		err := m.RegisterComponents(builder.Registry())
		if err != nil {
			return err
		}
	}

	ctx, err := builder.Create()
	if err != nil {
		return err
	}

	inst.b.context = ctx
	return nil
}

////////////////////////////////////////////////////////////////////////////////
