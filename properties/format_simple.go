package properties

import (
	"sort"
	"strings"
)

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
