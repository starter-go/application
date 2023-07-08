package boot

import (
	"github.com/starter-go/application/attributes"
	"github.com/starter-go/application/environment"
	"github.com/starter-go/application/parameters"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/safe"
)

// Options 包含启动选项
type Options struct {
	Mode safe.Mode

	Args []string

	Attributes  attributes.Table
	Environment environment.Table
	Parameters  parameters.Table
	Properties  properties.Table
}
