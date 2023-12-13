package properties

import (
	"sort"
	"strings"

	"github.com/starter-go/base/util"
)

type groupsFormatter struct{}

func (inst *groupsFormatter) fmt(t Table) string {

	// vlog.Warn("groupsFormatter: 目前没有实现分组格式化功能，暂时用 simpleFormatter 代替")

	src := t.Export(nil)
	b := new(groupsTextBuilder)
	for k, v := range src {
		b.add(k, v)
	}
	return b.create()
}

////////////////////////////////////////////////////////////////////////////////

type groupsTextBuilder struct {
	groups map[string]*group // map[prefix] * group
}

func (b *groupsTextBuilder) add(name, value string) {
	class, id, field := b.parseName(name)
	g := b.getGroup(class, id)
	g.put(field, value)
}

func (b *groupsTextBuilder) prefixOf(class, id string) string {
	pb := new(strings.Builder)
	if class != "" {
		pb.WriteString(class)
		pb.WriteRune('.')
	}
	if id != "" {
		pb.WriteString(id)
		pb.WriteRune('.')
	}
	return pb.String()
}

func (b *groupsTextBuilder) getGroup(class, id string) *group {
	all := b.getGroups()
	prefix := b.prefixOf(class, id)
	g := all[prefix]
	if g == nil {
		g = new(group)
		g.init(prefix, class, id)
		all[prefix] = g
	}
	return g
}

func (b *groupsTextBuilder) getGroups() map[string]*group {
	t := b.groups
	if t == nil {
		t = make(map[string]*group)
		b.groups = t
	}
	return t
}

// parse name to (class,id,field)
func (b *groupsTextBuilder) parseName(name string) (string, string, string) {

	parts := strings.Split(name, ".")
	cnt := len(parts)
	if cnt == 1 {
		return "", "", name
	} else if cnt == 2 {
		return parts[0], "", parts[1]
	} else if cnt == 3 {
		return parts[0], parts[1], parts[2]
	} else if cnt < 1 {
		return "", "", ""
	}

	// cnt >= 3
	sep := ""
	class := ""
	field := ""
	id := new(strings.Builder)
	end := cnt - 1
	for i, part := range parts {
		if i == 0 {
			class = part
		} else if i == end {
			field = part
		} else {
			id.WriteString(sep)
			id.WriteString(part)
			sep = "."
		}
	}
	return class, id.String(), field
}

func (b *groupsTextBuilder) create() string {

	table := b.getGroups()
	list := make([]*group, 0)
	for _, group := range table {
		list = append(list, group)
	}

	sorter := &util.Sorter{
		OnSwap: func(i1, i2 int) { list[i1], list[i2] = list[i2], list[i1] },
		OnLen:  func() int { return len(list) },
		OnLess: func(i1, i2 int) bool { return b.less(list[i1], list[i2]) },
	}
	sorter.Sort()

	sb := new(strings.Builder)
	for _, group := range list {
		group.writeHead(sb)
		group.writeKeyValues(sb)
	}
	return sb.String()
}

func (b *groupsTextBuilder) less(x, y *group) bool {
	if x == nil || y == nil {
		return false
	}
	v := strings.Compare(x.prefix, y.prefix)
	return v < 0
}

////////////////////////////////////////////////////////////////////////////////

type group struct {
	prefix string            // class + '.' + id
	class  string            // the group class-name
	id     string            // the group id-name
	table  map[string]string // map[ simple_name ] value
}

func (inst *group) init(prefix, class, id string) {
	inst.class = class
	inst.id = id
	inst.prefix = prefix
	inst.table = make(map[string]string)
}

func (inst *group) put(simpleName, value string) {
	inst.table[simpleName] = value
}

func (inst *group) writeHead(b *strings.Builder) {

	if inst.class == "" {
		b.WriteString("[]\n")
		return
	}

	b.WriteString("[")
	b.WriteString(inst.class)
	if inst.id != "" {
		b.WriteString(" \"")
		b.WriteString(inst.id)
		b.WriteString("\"")
	}
	b.WriteString("]\n")
}

func (inst *group) writeKeyValues(b *strings.Builder) {
	table := inst.table
	keys := []string{}
	for key := range table {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		val := table[key]
		b.WriteString("\t")
		b.WriteString(key)
		b.WriteString(" = ")
		b.WriteString(val)
		b.WriteString("\n")
	}
}

////////////////////////////////////////////////////////////////////////////////
