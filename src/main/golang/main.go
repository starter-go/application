package main

import (
	"embed"
	"strings"

	"github.com/starter-go/application"
	"github.com/starter-go/application/boot"
	"github.com/starter-go/application/components"
	"github.com/starter-go/vlog"
)

//go:embed "res"
var theResFS embed.FS

func main() {
	vlog.Debug("hello")
	mod := m3()
	err := boot.Run(mod)
	if err != nil {
		panic(err)
	}
}

func m1() application.Module {
	mb := &application.ModuleBuilder{}
	mb.Name("demo.m1").Version("0.0.1").Revision(1)
	mb.EmbedResources(theResFS, "res")

	mb.Components(func(r components.Registry) error {
		com1 := r.New()
		com1.ID = "com-1"
		com1.Classes = "Demo Example"
		com1.Aliases = "demo-1 demo-2"
		com1.Scope = "singleton"

		com1.InjectFunc = func(c components.Injection, instance any) error {
			c2inst, err := c.GetByID("com-2")
			if err != nil {
				return err
			}
			c2 := c2inst.Get().(*Com2)
			c1 := instance.(*Com1)
			c1.c2 = c2
			return nil
		}
		com1.NewFunc = func() any {
			return &Com1{}
		}

		r.Register(com1)
		return nil
	})

	return mb.Create()
}

func m2() application.Module {
	mb := &application.ModuleBuilder{}
	mb.Name("demo.m2").Version("0.0.1").Revision(1)
	mb.EmbedResources(theResFS, "res")

	mb.Components(func(r components.Registry) error {
		com2 := r.New()
		com2.ID = "com-2"
		com2.Classes = "Demo Example"
		com2.Aliases = "demo-1 demo-2"
		com2.Scope = "singleton"

		com2.InjectFunc = func(c components.Injection, instance any) error {
			c1inst, err := c.GetByID("com-1")
			if err != nil {
				return err
			}
			c1 := c1inst.Get().(*Com1)
			c2 := instance.(*Com2)
			c2.c1 = c1
			return nil
		}
		com2.NewFunc = func() any {
			return &Com2{}
		}

		r.Register(com2)
		return nil
	})

	m1 := m1()
	mb.Depend(m1)
	return mb.Create()
}

func m3() application.Module {
	mb := &application.ModuleBuilder{}
	mb.Name("demo.m3").Version("0.0.1").Revision(1)
	mb.EmbedResources(theResFS, "res")

	mb.Components(func(r components.Registry) error {
		com1 := r.New()
		com1.ID = "demo3"
		com1.Classes = "Demo Example"
		com1.Aliases = "demo-1 demo-2"
		com1.Scope = "singleton"

		com1.NewFunc = func() any {
			return &strings.Builder{}
		}
		com1.InjectFunc = func(c components.Injection, instance any) error {

			d1, err := c.SelectOne("#com-1")
			if err != nil {
				return err
			}

			builder := instance.(*strings.Builder)
			str := d1.Info().ID().String()
			builder.WriteString(str)

			return nil
		}

		r.Register(com1)
		return nil
	})

	m1 := m1()
	m2 := m2()
	mb.Depend(m1, m2)
	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////

type Com1 struct {
	c2 *Com2
}

func (inst *Com1) _Impl() application.Lifecycle {
	return inst
}

func (inst *Com1) Life() *application.Life {
	return &application.Life{
		OnCreate:  inst.create,
		OnStart:   inst.start,
		OnLoop:    inst.loop,
		OnStop:    inst.stop,
		OnDestroy: inst.destroy,
	}
}

func (inst *Com1) nop() error {
	return nil
}

func (inst *Com1) create() error {
	vlog.Info("com1.create()")
	return nil
}

func (inst *Com1) start() error {
	vlog.Info("com1.start()")
	return nil
}

func (inst *Com1) loop() error {
	vlog.Info("com1.loop()")
	return nil
}

func (inst *Com1) stop() error {
	vlog.Info("com1.stop()")
	return nil
}

func (inst *Com1) destroy() error {
	vlog.Info("com1.destroy()")
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type Com2 struct {
	c1 *Com1
}

////////////////////////////////////////////////////////////////////////////////
