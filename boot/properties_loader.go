package boot

import (
	"os"
	"strconv"
	"strings"

	"github.com/starter-go/afs/files"
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
		value, err := t.GetPropertyRequired(name)
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
	inst.loadPropertiesFromBootstrap(builder)
	inst.loadPropertiesFromLocalFile(builder)

	return builder.all
}

func (inst *propertiesLoader) loadPropertiesFromBootstrap(builder *propertiesLoaderBuilder) {
	table := inst.b.collections.Properties
	if table == nil {
		return
	}
	builder.add(table)
}

func (inst *propertiesLoader) loadPropertiesFromLocalFile(builder *propertiesLoaderBuilder) {
	const (
		keyFile    = "application.properties.file"
		keyEnabled = "application.properties.enabled"
	)
	strEnabled := builder.getProperty(keyEnabled)
	enabled, _ := strconv.ParseBool(strEnabled)
	if !enabled {
		return
	}
	strFile := builder.getProperty(keyFile)
	if strFile == "" {
		return
	}
	vlog.Info("load properties from file [%s]", strFile)
	path := files.FS().NewPath(strFile)
	text, err := path.GetIO().ReadText(nil)
	if err != nil {
		vlog.Warn(err.Error())
		return
	}
	table, err := properties.Parse(text, nil)
	if err != nil {
		vlog.Warn(err.Error())
		return
	}
	builder.add(table)
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
	// try get cached
	table := inst.propertiesFromExeDir
	if table != nil {
		builder.add(table)
		return
	}
	// locate 'application.properties' file
	exepath := ""
	for i, a := range os.Args {
		if i == 0 {
			exepath = a
			break
		}
	}
	if exepath == "" {
		return
	}
	exefile := files.FS().NewPath(exepath)
	if !exefile.IsFile() {
		return
	}
	apppropfile := exefile.GetParent().GetChild("application.properties")
	if !apppropfile.IsFile() {
		return
	}
	// load from file
	str, err := apppropfile.GetIO().ReadText(nil)
	if err != nil {
		panic(err)
	}
	table, err = properties.Parse(str, nil)
	if err != nil {
		panic(err)
	}
	builder.add(table)
	inst.propertiesFromExeDir = table
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
		p = inst.tryLoadByPath("/application.properties")
		inst.props1 = p
	}
	builder.add(p)
}

func (inst *propertiesLoaderModuleHolder) load2(builder *propertiesLoaderBuilder, profile string) {
	p := inst.props2
	if p == nil {
		p = inst.tryLoadByPath("/application-" + profile + ".properties")
		inst.props2 = p
	}
	builder.add(p)
}

func (inst *propertiesLoaderModuleHolder) tryLoadByPath(path string) properties.Table {
	table, err := inst.loadByPath(path)
	if err != nil {
		vlog.Debug("error: %v", err)
	}
	if table == nil {
		table = properties.NewTable(nil)
	}
	return table
}

func (inst *propertiesLoaderModuleHolder) loadByPath(path string) (properties.Table, error) {
	res := inst.mod.Resources()
	return properties.LoadFromResource(path, res, nil)
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

func (inst *propertiesLoaderBuilder) getProperty(name string) string {
	all := inst.all
	for i := len(all) - 1; i >= 0; i-- {
		t := all[i]
		if t == nil {
			continue
		}
		value := t.GetProperty(name)
		if value != "" {
			return value
		}
	}
	return ""
}

////////////////////////////////////////////////////////////////////////////////
