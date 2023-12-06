package properties

import (
	"sort"
	"strings"

	"github.com/starter-go/vlog"
)

// Format 把 Table 格式化为字符串
func Format(t Table, options ...FormatOptionsF) string {

	if t == nil {
		return ""
	}

	opt := new(FormatOptions)
	for _, f := range options {
		f(opt)
	}

	if opt.useGroups {
		formatter := new(groupsFormatter)
		return formatter.fmt(t)
	}

	formatter := new(simpleFormatter)
	return formatter.fmt(t)
}

////////////////////////////////////////////////////////////////////////////////

// FormatOptions ...
type FormatOptions struct {
	useGroups bool
}

// FormatOptionsF ...
type FormatOptionsF func(opt *FormatOptions)

// FormatWithGroups 启用分组功能
func FormatWithGroups(opt *FormatOptions) {
	opt.useGroups = true
}

////////////////////////////////////////////////////////////////////////////////

type simpleFormatter struct{}

func (inst *simpleFormatter) fmt(t Table) string {
	list := make([]string, 0)
	src := t.Export(nil)
	for k, v := range src {
		list = append(list, k+" = "+v+"\n")
	}
	sort.Strings(list)
	b := new(strings.Builder)
	for _, str := range list {
		b.WriteString(str)
	}
	return b.String()
}

////////////////////////////////////////////////////////////////////////////////

type groupsFormatter struct{}

func (inst *groupsFormatter) fmt(t Table) string {

	vlog.Warn("groupsFormatter: 目前没有实现分组格式化功能，暂时用 simpleFormatter 代替")

	sf := new(simpleFormatter)
	return sf.fmt(t)
}

////////////////////////////////////////////////////////////////////////////////
