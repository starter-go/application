package main

import (
	"embed"
	"os"
	"strings"

	"github.com/starter-go/application"
	"github.com/starter-go/application/boot"
	"github.com/starter-go/vlog"
)

//go:embed "res"
var theResFS embed.FS

func main() {
	vlog.Debug("hello")
	opt := &boot.Options{
		Args: os.Args,
	}
	mod := m3()
	err := boot.Run(mod, opt)
	if err != nil {
		panic(err)
	}
}

func m1() application.Module {
	mb := &application.ModuleBuilder{}
	mb.Name("demo.m1").Version("0.0.1").Revision(1)
	mb.EmbedResources(theResFS, "res")

	mb.Components(func(r application.ComponentRegistry) error {
		com1 := r.NewRegistration()
		com1.ID = "com-1"
		com1.Classes = "Demo Example"
		com1.Aliases = "demo-1 demo-2"
		com1.Scope = "singleton"

		com1.InjectFunc = func(c application.InjectionExt, instance any) error {
			c2 := c.GetComponent("#com-2").(*Com2)
			c1 := instance.(*Com1)
			c1.c2 = c2
			return nil
		}
		com1.NewFunc = func() any {
			return &Com1{}
		}

		return com1.Commit()
	})

	return mb.Create()
}

func m2() application.Module {
	mb := &application.ModuleBuilder{}
	mb.Name("demo.m2").Version("0.0.1").Revision(1)
	mb.EmbedResources(theResFS, "res")

	mb.Components(func(r application.ComponentRegistry) error {
		com2 := r.NewRegistration()
		com2.ID = "com-2"
		com2.Classes = "Demo Example"
		com2.Aliases = "demo-1 demo-2"
		com2.Scope = "singleton"

		com2.InjectFunc = func(c application.InjectionExt, instance any) error {
			c1 := c.GetComponent("#com-1").(*Com1)
			c2 := instance.(*Com2)
			c2.c1 = c1
			return nil
		}
		com2.NewFunc = func() any {
			return &Com2{}
		}

		return com2.Commit()
	})

	m1 := m1()
	mb.Depend(m1)
	return mb.Create()
}

func m3() application.Module {
	mb := &application.ModuleBuilder{}
	mb.Name("demo.m3").Version("0.0.1").Revision(1)
	mb.EmbedResources(theResFS, "res")

	mb.Components(func(r application.ComponentRegistry) error {
		com1 := r.NewRegistration()
		com1.ID = "demo3"
		com1.Classes = "Demo Example"
		com1.Aliases = "demo-1 demo-2"
		com1.Scope = "singleton"

		com1.NewFunc = func() any {
			return &strings.Builder{}
		}
		com1.InjectFunc = func(inj application.InjectionExt, instance any) error {

			// d1  :=  inj.GetComponent ("#com-1")
			str := ""
			builder := instance.(*strings.Builder)
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
