package boot

import (
	"strings"

	"github.com/starter-go/application"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/vlog"
)

type propertiesLoader struct {
	b                    *Bootstrap
	profile              string // the profile name
	moduleHolders        []*propertiesLoaderModuleHolder
	propertiesFromArgs   properties.Table
	propertiesFromExeDir properties.Table
}

func (inst *propertiesLoader) load() error {
	err := inst.loadProfileName()
	if err != nil {
		return err
	}
	vlog.Info("application.profiles.active: %s", inst.profile)
	return inst.loadAllProperties()
}

func (inst *propertiesLoader) loadProfileName() error {
	const name = "application.profiles.active"
	list := inst.listProperties()
	profile := inst.profile
	for _, t := range list {
		value, err := t.GetProperty(name)
		if err == nil {
			profile = value
		}
	}
	if profile == "" {
		profile = "default"
	}
	inst.profile = profile
	return nil
}

func (inst *propertiesLoader) loadAllProperties() error {
	dst := properties.NewTable(nil)
	list := inst.listProperties()
	for _, t := range list {
		dst.Import(t.Export(nil))
	}
	inst.b.collections.Properties.Import(dst.Export(nil))
	return nil
}

func (inst *propertiesLoader) listProperties() []properties.Table {

	builder := &propertiesLoaderBuilder{}

	inst.loadPropertiesFromRes(builder)
	inst.loadPropertiesFromExeDir(builder)
	inst.loadPropertiesFromArgs(builder)

	return builder.all
}

func (inst *propertiesLoader) loadPropertiesFromArgs(builder *propertiesLoaderBuilder) {
	const (
		prefix = "--"
		eq     = "="
	)
	props := inst.propertiesFromArgs
	if props != nil {
		builder.add(props)
		return
	}
	props = properties.NewTable(nil)
	args := inst.b.collections.Arguments.Raw()
	for _, a := range args {
		i := strings.Index(a, eq)
		if strings.HasPrefix(a, prefix) && i > 0 {
			key := strings.TrimSpace(a[len(prefix):i])
			val := strings.TrimSpace(a[i+1:])
			props.SetProperty(key, val)
		}
	}
	inst.propertiesFromArgs = props
	builder.add(props)
}

func (inst *propertiesLoader) loadPropertiesFromExeDir(builder *propertiesLoaderBuilder) {
	// todo ...
}

func (inst *propertiesLoader) loadPropertiesFromRes(builder *propertiesLoaderBuilder) {
	holders := inst.moduleHolders
	if holders == nil {
		mods := inst.b.modules
		for _, m := range mods {
			h := &propertiesLoaderModuleHolder{mod: m}
			holders = append(holders, h)
		}
		inst.moduleHolders = holders
	}
	profile := inst.profile
	for _, h := range holders {
		h.load1(builder)
		if profile != "" {
			h.load2(builder, profile)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////

type propertiesLoaderModuleHolder struct {
	mod    application.Module
	props1 properties.Table // 'application.properties'
	props2 properties.Table // 'application-x.properties'
}

func (inst *propertiesLoaderModuleHolder) load1(builder *propertiesLoaderBuilder) {
	p := inst.props1
	if p == nil {
		p = inst.loadByPath("/application.properties")
		inst.props1 = p
	}
	builder.add(p)
}

func (inst *propertiesLoaderModuleHolder) load2(builder *propertiesLoaderBuilder, profile string) {
	p := inst.props2
	if p == nil {
		p = inst.loadByPath("/application-" + profile + ".properties")
		inst.props2 = p
	}
	builder.add(p)
}

func (inst *propertiesLoaderModuleHolder) loadByPath(path string) properties.Table {
	table := properties.NewTable(nil)
	res, err := inst.mod.Resources().GetResource(path)
	if err != nil {
		vlog.Warn("%v", err)
		return table
	}
	text, err := res.ReadText()
	if err != nil {
		vlog.Warn("%v", err)
		return table
	}
	table, err = properties.Parse(text, nil)
	if err != nil {
		vlog.Warn("%v", err)
		return table
	}
	return table
}

////////////////////////////////////////////////////////////////////////////////

type propertiesLoaderBuilder struct {
	all []properties.Table
}

func (inst *propertiesLoaderBuilder) add(t properties.Table) {
	if t == nil {
		return
	}
	inst.all = append(inst.all, t)
}

////////////////////////////////////////////////////////////////////////////////
