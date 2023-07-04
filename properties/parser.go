package properties

import (
	"fmt"
	"strings"

	"github.com/starter-go/base/safe"
)

// Parse 解析属性表
func Parse(str string, mode safe.Mode) (Table, error) {
	parser := &parser{}
	err := parser.parse(str)
	if err != nil {
		return nil, err
	}
	t := NewTable(mode)
	t.Import(parser.table)
	return t, nil
}

type parser struct {
	table   map[string]string
	segment string // like 'a.b.'
}

func (inst *parser) parse(str string) error {
	const (
		ch1 = "\r"
		ch2 = "\n"
	)
	inst.table = make(map[string]string)
	str = strings.ReplaceAll(str, ch1, ch2)
	rows := strings.Split(str, ch2)
	for i, row := range rows {
		row = strings.TrimSpace(row)
		err := inst.handleRow(row, i+1)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *parser) handleRow(row string, num int) error {
	if row == "" {
		return nil
	} else if strings.HasPrefix(row, "#") {
		return nil
	} else if strings.HasPrefix(row, "[") && strings.HasSuffix(row, "]") {
		return inst.handleSegment(row, num)
	} else if strings.ContainsRune(row, '=') {
		return inst.handleKeyValue(row, num)
	}
	return fmt.Errorf("bad properties row[%d]: %s", num, row)
}

func (inst *parser) handleSegment(row string, num int) error {
	const (
		ch1 = "["
		ch2 = "]"
		ch3 = "'"
		ch4 = "\""
		chx = "\n"
	)
	str := row
	str = strings.ReplaceAll(str, ch1, chx)
	str = strings.ReplaceAll(str, ch2, chx)
	str = strings.ReplaceAll(str, ch3, chx)
	str = strings.ReplaceAll(str, ch4, chx)
	parts := strings.Split(str, chx)
	builder := strings.Builder{}
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			builder.WriteString(part)
			builder.WriteRune('.')
		}
	}
	inst.segment = builder.String()
	return nil
}

func (inst *parser) handleKeyValue(row string, num int) error {
	i := strings.Index(row, "=")
	key := strings.TrimSpace(row[0:i])
	val := strings.TrimSpace(row[i+1:])
	name := inst.segment + key
	inst.table[name] = val
	return nil
}
