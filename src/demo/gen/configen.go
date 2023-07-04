package gen

import (
	"io"

	"github.com/starter-go/application/components"
	"github.com/starter-go/application/src/demo/parts"
)

func Config(r components.Registry) error {

	c1 := &com1{}
	c2 := &com2{}

	c1.cfg(r)
	c2.cfg(r)

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type com1 struct {
}

func (inst *com1) cfg(r components.Registry) error {
	x := r.New()
	x.Classes = "parts.Com1"
	x.ID = "com-1"
	x.Aliases = "io.Closer"
	x.Scope = ""
	x.NewFunc = inst.new
	x.InjectFunc = inst.inject
	r.Register(x)
	return nil
}

func (inst *com1) new() any {
	return &parts.Com1{}
}

func (inst *com1) inject(ci components.Injection, instance any) error {
	o := instance.(*parts.Com1)
	cix := ci.Ext()
	if cix == nil || o == nil {
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type com2 struct {
}

func (inst *com2) cfg(r components.Registry) error {
	x := r.New()
	x.Classes = ""
	x.ID = "com-2"
	x.Aliases = ""
	x.Scope = ""
	x.NewFunc = inst.new
	x.InjectFunc = inst.inject
	r.Register(x)
	return nil
}

func (inst *com2) new() any {
	return &parts.Com2{}
}

func (inst *com2) inject(ci components.Injection, instance any) error {

	o := instance.(*parts.Com2)
	cix := ci.Ext()

	o.F1 = cix.GetString("${p.com2.string}")
	o.F2 = cix.GetInt("${p.com2.int}")
	o.F3 = cix.GetBool("${p.com2.bool}")
	o.F4 = cix.GetFloat64("${p.com2.float64}")
	o.F5 = cix.GetInt32("${p.com2.int32}")
	o.F6 = cix.GetByte("${p.com2.byte}")
	o.F7 = cix.GetComponent("#io.Closer").(io.Closer)
	o.F8 = inst.forF8(cix.ListComponents(".parts.Com1"))

	return nil
}

func (inst *com2) forF8(src []any) []*parts.Com1 {
	dst := make([]*parts.Com1, 0)
	for _, item1 := range src {
		item2 := item1.(*parts.Com1)
		dst = append(dst, item2)
	}
	return dst
}
