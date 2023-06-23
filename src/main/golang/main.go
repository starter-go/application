package main

import (
	"embed"

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
		com1.ID = "demo1"
		com1.Classes = "Demo Example"
		com1.Aliases = "demo-1 demo-2"
		com1.Scope = "singleton"
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
		com1 := r.New()
		com1.ID = "demo2"
		com1.Classes = "Demo Example"
		com1.Aliases = "demo-1 demo-2"
		com1.Scope = "singleton"
		r.Register(com1)
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
		r.Register(com1)
		return nil
	})

	m1 := m1()
	m2 := m2()
	mb.Depend(m1, m2)
	return mb.Create()
}
