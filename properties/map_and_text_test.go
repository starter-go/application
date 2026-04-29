package properties

import (
	"testing"

	"github.com/starter-go/base/lang"
)

////////////////////////////////////////////////////////////////////////////////
// Map

func TestMap(t *testing.T) {

	ttt := &innerTestMapAndText{t: t}
	m := ttt.innerGetExampleMap()
	str := m.String()

	t.Logf("table = \n %s", str)
}

func TestMapDoInit(t *testing.T) {

	// ttt := &innerTestMapAndText{t: t}
	// m := ttt.innerGetExampleMap()
	// str := m.String()

	var m1 Map
	m2 := m1.Init()

	t.Logf("table[m1] = \n %v", m1)
	t.Logf("table[m2] = \n %v", m2)
}

func TestMapDoExport(t *testing.T) {

	ttt := &innerTestMapAndText{t: t}
	m := ttt.innerGetExampleMap()

	m2 := m.Export(nil)

	t.Logf("table = \n %v", m2)
}

func TestMapDoTable(t *testing.T) {

	ttt := &innerTestMapAndText{t: t}
	m := ttt.innerGetExampleMap()

	tab := m.Table(nil)

	t.Logf("table = \n %v", tab)
}

func TestMapDoTrim(t *testing.T) {

	ttt := &innerTestMapAndText{t: t}
	m := ttt.innerGetExampleMap()

	m1 := m.Trim()
	m2 := m.Trim()

	m2.Put("i", "")
	m2.Put("j", "")
	m2.Put("k", "")
	m2.Put("", "haha")

	m3 := m2.Trim()

	t.Logf("table(m1) = \n %v", m1)
	t.Logf("table(m2) = \n %v", m2)
	t.Logf("table(m3) = \n %v", m3)
}

////////////////////////////////////////////////////////////////////////////////
// Text

func TestText(t *testing.T) {

	ttt := &innerTestMapAndText{t: t}
	text := ttt.innerGetExampleText()
	str := text.String()

	t.Logf("table = \n %s", str)
}

func TestText2String(t *testing.T) {

	ttt := &innerTestMapAndText{t: t}
	text := ttt.innerGetExampleText()

	str := text.String()

	t.Logf("table = \n %s", str)
}

func TestText2Map(t *testing.T) {

	ttt := &innerTestMapAndText{t: t}
	text := ttt.innerGetExampleText()

	m := text.Map()

	str := m.String()
	t.Logf("table = \n %s", str)
}

////////////////////////////////////////////////////////////////////////////////
// class

type innerTestMapAndText struct {
	t *testing.T
}

func (inst *innerTestMapAndText) innerGetExampleText() Text {
	m := inst.innerGetExampleMap()
	return m.Text()
}

func (inst *innerTestMapAndText) innerGetExampleMap() Map {
	d := inst.innerGetExampleData()
	return d
}

func (inst *innerTestMapAndText) innerGetExampleData() map[string]string {

	now := lang.Now()
	m := make(map[string]string)

	m["a"] = "1"
	m["b"] = "2"
	m["c"] = "3"
	m["x"] = "77"
	m["y"] = "88"
	m["z"] = "99"

	m["now"] = now.String()

	return m
}

////////////////////////////////////////////////////////////////////////////////
// EOF
