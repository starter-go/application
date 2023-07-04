package main

import (
	"embed"
	"os"

	"github.com/starter-go/application"
	"github.com/starter-go/application/boot"
	"github.com/starter-go/application/src/demo/gen"
)

func main() {
	opt := &boot.Options{
		Args: os.Args,
	}
	mod := theModule()
	err := boot.Run(mod, opt)
	if err != nil {
		panic(err)
	}
}

////////////////////////////////////////////////////////////////////////////////

const (
	theModuleName     = "src/demo"
	theModuleVersion  = "v1"
	theModuleRevision = 1
	theModuleResPath  = "res"
)

//go:embed "res"
var theModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

func theModule() application.Module {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName).Version(theModuleVersion).Revision(theModuleRevision)
	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Depend(nil)
	mb.Components(gen.Config)
	return mb.Create()
}
