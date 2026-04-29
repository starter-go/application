package properties

import (
	"fmt"
	"sort"
	"strings"

	"github.com/starter-go/base/safe"
)

////////////////////////////////////////////////////////////////////////////////
// Map

type Map map[string]string

func (m Map) String() string {
	t := m.Text()
	return t.String()
}

func (m Map) Get(name string) string {
	if m == nil {
		return ""
	}
	return m[name]
}

func (m Map) Put(name, value string) error {

	if m == nil {
		return fmt.Errorf("properties.Map is nil")
	}

	m[name] = value
	return nil
}

func (m Map) Trim() Map {

	src := m
	dst := make(Map)

	for k, v := range src {
		if k == "" || v == "" {
			continue
		}
		dst[k] = v
	}

	return dst
}

func (m Map) Text() Text {

	if m == nil {
		return ""
	}

	// format

	keys := []string{}
	builder := new(strings.Builder)

	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		v := m[k]
		builder.WriteString(k)
		builder.WriteRune('=')
		builder.WriteString(v)
		builder.WriteRune('\n')
	}

	str := builder.String()
	return Text(str)
}

func (m Map) Init() Map {
	if m == nil {
		m = make(Map)
	}
	return m
}

func (m Map) Table(mode safe.Mode) Table {

	dst := NewTable(mode)

	if m == nil {
		return dst
	}

	for k, v := range m {
		dst.SetProperty(k, v)
	}

	return dst
}

func (m Map) Export(dst map[string]string) map[string]string {

	if dst == nil {
		dst = make(map[string]string)
	}

	if m == nil {
		return dst
	}

	for k, v := range m {
		dst[k] = v
	}

	return dst
}

////////////////////////////////////////////////////////////////////////////////
// Text

type Text string

func (t Text) String() string {
	return string(t)
}

func (t Text) Map() Map {

	// parse

	str := t.String()
	mode := safe.Safe()
	table, err := Parse(str, mode)
	if err != nil {
		return make(Map)
	}
	m := table.Export(nil)
	return m
}

////////////////////////////////////////////////////////////////////////////////
// EOF
