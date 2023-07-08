package application

import (
	"github.com/starter-go/application/arguments"
	"github.com/starter-go/application/attributes"
	"github.com/starter-go/application/environment"
	"github.com/starter-go/application/parameters"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/application/resources"
	"github.com/starter-go/base/safe"
)

// Collections 持有上下文中的各种资源集合
type Collections struct {
	Arguments   arguments.Table
	Attributes  attributes.Table
	Environment environment.Table
	Parameters  parameters.Table
	Properties  properties.Table
	Resources   resources.Table
}

// Complete 完善集合中的内容（如果需要）
func (inst *Collections) Complete(mode safe.Mode) {

	if mode == nil {
		mode = safe.Default()
	}

	if inst.Arguments == nil {
		inst.Arguments = arguments.NewTable(nil, mode)
	}

	if inst.Attributes == nil {
		inst.Attributes = attributes.NewTable(mode)
	}

	if inst.Environment == nil {
		inst.Environment = environment.NewTable(mode)
	}

	if inst.Parameters == nil {
		inst.Parameters = parameters.NewTable(mode)
	}

	if inst.Properties == nil {
		inst.Properties = properties.NewTable(mode)
	}

	if inst.Resources == nil {
		empty := make(map[string]resources.Resource)
		inst.Resources = resources.NewTable(empty, mode)
	}

}

// Clone 以指定模式克隆这个集合
func (inst *Collections) Clone(mode safe.Mode) *Collections {

	if mode == nil {
		mode = safe.Default()
	}

	co2 := &Collections{}
	co2.Complete(mode)

	args := inst.Arguments
	atts := inst.Attributes
	env := inst.Environment
	param := inst.Parameters
	props := inst.Properties
	res := inst.Resources

	if args != nil {
		co2.Arguments = arguments.NewTable(args.Raw(), mode)
	}

	if atts != nil {
		co2.Attributes.Import(atts.Export(nil))
	}

	if env != nil {
		co2.Environment.Import(env.Export(nil))
	}

	if param != nil {
		co2.Parameters.Import(param.Export(nil))
	}

	if props != nil {
		co2.Properties.Import(props.Export(nil))
	}

	if res != nil {
		co2.Resources.Import(res.Export(nil))
	}

	return co2
}
