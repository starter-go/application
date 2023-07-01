package boot

import "github.com/starter-go/application/components/implcom"

type contextLoader struct {
	b *Bootstrap
}

func (inst *contextLoader) load() error {

	mods := inst.b.modules
	builder := implcom.NewBuilder()

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
