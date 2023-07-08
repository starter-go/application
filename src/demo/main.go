package main

import (
	"embed"
	"os"

	"github.com/starter-go/application"
	"github.com/starter-go/application/boot"
	"github.com/starter-go/application/src/demo/gen"
	"github.com/starter-go/base/safe"
)

func main() {

	mode := safe.Safe()
	coll := &application.Collections{}
	coll.Complete(mode)

	coll.Attributes.SetAttribute("demo-attr-name", "demo-attr-obj")
	coll.Environment.SetEnv("demo-env-name", "demo-env-value")
	coll.Properties.SetProperty("demo-prop-name", "demo-prop-value")
	coll.Parameters.SetParam("demo-param-name", "demo-param-value")

	opt := &boot.Options{}
	opt.Mode = mode
	opt.Args = os.Args
	opt.Attributes = coll.Attributes
	opt.Environment = coll.Environment
	opt.Parameters = coll.Parameters
	opt.Properties = coll.Properties

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
