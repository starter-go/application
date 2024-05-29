package properties

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Getter 属性值读取器
type Getter interface {
	Required() Getter
	Optional() Getter
	Error() error

	GetBool(name string, defaultValue ...bool) bool
	GetByte(name string, defaultValue ...byte) byte
	GetRune(name string, defaultValue ...rune) rune
	GetString(name string, defaultValue ...string) string
	GetStringList(name string, defaultValue ...string) []string

	GetFloat32(name string, defaultValue ...float32) float32
	GetFloat64(name string, defaultValue ...float64) float64

	GetInt(name string, defaultValue ...int) int
	GetInt8(name string, defaultValue ...int8) int8
	GetInt16(name string, defaultValue ...int16) int16
	GetInt32(name string, defaultValue ...int32) int32
	GetInt64(name string, defaultValue ...int64) int64

	GetUint(name string, defaultValue ...uint) uint
	GetUint8(name string, defaultValue ...uint8) uint8
	GetUint16(name string, defaultValue ...uint16) uint16
	GetUint32(name string, defaultValue ...uint32) uint32
	GetUint64(name string, defaultValue ...uint64) uint64

	// 根据给出的名称前缀和后缀，列出项目 id
	ListItems(prefix, suffix string) []string
}

////////////////////////////////////////////////////////////////////////////////

type myGetter struct {
	table    Table
	required bool
	errors   []error
}

func (g *myGetter) _impl(a Getter) {
	a = g
}

func (g *myGetter) Required() Getter {
	return &myGetter{
		table:    g.table,
		required: true,
	}
}

func (g *myGetter) Optional() Getter {
	return &myGetter{
		table:    g.table,
		required: false,
	}
}

func (g *myGetter) getString(name string) (string, bool) {
	str, err := g.table.GetPropertyRequired(name)
	if err != nil {
		if g.required {
			g.handleError(err)
		}
	}
	ok := (err == nil)
	return str, ok
}

func (g *myGetter) Error() error {
	list := g.errors
	if list == nil {
		return nil
	}
	if len(list) == 0 {
		return nil
	}
	b := strings.Builder{}
	for _, e := range list {
		b.WriteString(e.Error())
		b.WriteString("\n")
	}
	return fmt.Errorf("properties.Getter error list:\n%s", b.String())
}

func (g *myGetter) handleError(err error) {
	if err == nil {
		return
	}
	g.errors = append(g.errors, err)
}

func (g *myGetter) handleParsingError(name string, value string, err error) {
	if err == nil {
		return
	}
	const f = "parse property error, name:%s error:%s"
	err2 := fmt.Errorf(f, name, err.Error())
	g.errors = append(g.errors, err2)
}

func (g *myGetter) GetBool(name string, defaultValue ...bool) bool {
	str, ok := g.getString(name)
	if ok {
		b, err := strconv.ParseBool(str)
		g.handleParsingError(name, str, err)
		if err == nil {
			return b
		}
	}
	for _, value := range defaultValue {
		return value
	}
	return false
}

func (g *myGetter) GetByte(name string, defaultValue ...byte) byte {
	const (
		base    = 10
		bitSize = 8
	)
	str, ok := g.getString(name)
	if ok {
		n, err := strconv.ParseInt(str, base, bitSize)
		g.handleParsingError(name, str, err)
		if err == nil {
			return byte(n)
		}
	}
	for _, value := range defaultValue {
		return value
	}
	return 0
}

func (g *myGetter) GetRune(name string, defaultValue ...rune) rune {
	str, ok := g.getString(name)
	if ok {
		chs := []rune(str)
		for _, ch := range chs {
			return ch
		}
	}
	for _, value := range defaultValue {
		return value
	}
	return 0
}

func (g *myGetter) GetString(name string, defaultValue ...string) string {
	str, ok := g.getString(name)
	if ok {
		return str
	}
	for _, value := range defaultValue {
		return value
	}
	return ""
}

func (g *myGetter) GetStringList(name string, defaultValue ...string) []string {
	const sep = ","
	dst := make([]string, 0)
	src := defaultValue
	str, ok := g.getString(name)
	if ok {
		src = strings.Split(str, sep)
	}
	for _, part := range src {
		part = strings.TrimSpace(part)
		dst = append(dst, part)
	}
	return dst
}

func (g *myGetter) GetFloat32(name string, defaultValue ...float32) float32 {
	const bitSize = 32
	str, ok := g.getString(name)
	if ok {
		n, err := strconv.ParseFloat(str, bitSize)
		g.handleParsingError(name, str, err)
		if err == nil {
			return float32(n)
		}
	}
	for _, value := range defaultValue {
		return value
	}
	return 0
}

func (g *myGetter) GetFloat64(name string, defaultValue ...float64) float64 {
	const bitSize = 64
	str, ok := g.getString(name)
	if ok {
		n, err := strconv.ParseFloat(str, bitSize)
		g.handleParsingError(name, str, err)
		if err == nil {
			return n
		}
	}
	for _, value := range defaultValue {
		return value
	}
	return 0
}

func (g *myGetter) GetInt(name string, defaultValue ...int) int {
	const (
		base    = 10
		bitSize = 0
	)
	str, ok := g.getString(name)
	if ok {
		n, err := strconv.ParseInt(str, base, bitSize)
		g.handleParsingError(name, str, err)
		if err == nil {
			return int(n)
		}
	}
	for _, value := range defaultValue {
		return value
	}
	return 0
}

func (g *myGetter) GetInt8(name string, defaultValue ...int8) int8 {
	const (
		base    = 10
		bitSize = 8
	)
	str, ok := g.getString(name)
	if ok {
		n, err := strconv.ParseInt(str, base, bitSize)
		g.handleParsingError(name, str, err)
		if err == nil {
			return int8(n)
		}
	}
	for _, value := range defaultValue {
		return value
	}
	return 0
}

func (g *myGetter) GetInt16(name string, defaultValue ...int16) int16 {
	const (
		base    = 10
		bitSize = 16
	)
	str, ok := g.getString(name)
	if ok {
		n, err := strconv.ParseInt(str, base, bitSize)
		g.handleParsingError(name, str, err)
		if err == nil {
			return int16(n)
		}
	}
	for _, value := range defaultValue {
		return value
	}
	return 0
}

func (g *myGetter) GetInt32(name string, defaultValue ...int32) int32 {
	const (
		base    = 10
		bitSize = 32
	)
	str, ok := g.getString(name)
	if ok {
		n, err := strconv.ParseInt(str, base, bitSize)
		g.handleParsingError(name, str, err)
		if err == nil {
			return int32(n)
		}
	}
	for _, value := range defaultValue {
		return value
	}
	return 0
}

func (g *myGetter) GetInt64(name string, defaultValue ...int64) int64 {
	const (
		base    = 10
		bitSize = 64
	)
	str, ok := g.getString(name)
	if ok {
		n, err := strconv.ParseInt(str, base, bitSize)
		g.handleParsingError(name, str, err)
		if err == nil {
			return n
		}
	}
	for _, value := range defaultValue {
		return value
	}
	return 0
}

func (g *myGetter) GetUint(name string, defaultValue ...uint) uint {
	const (
		base    = 10
		bitSize = 0
	)
	str, ok := g.getString(name)
	if ok {
		n, err := strconv.ParseUint(str, base, bitSize)
		g.handleParsingError(name, str, err)
		if err == nil {
			return uint(n)
		}
	}
	for _, value := range defaultValue {
		return value
	}
	return 0
}

func (g *myGetter) GetUint8(name string, defaultValue ...uint8) uint8 {
	const (
		base    = 10
		bitSize = 8
	)
	str, ok := g.getString(name)
	if ok {
		n, err := strconv.ParseUint(str, base, bitSize)
		g.handleParsingError(name, str, err)
		if err == nil {
			return uint8(n)
		}
	}
	for _, value := range defaultValue {
		return value
	}
	return 0
}

func (g *myGetter) GetUint16(name string, defaultValue ...uint16) uint16 {
	const (
		base    = 10
		bitSize = 16
	)
	str, ok := g.getString(name)
	if ok {
		n, err := strconv.ParseUint(str, base, bitSize)
		g.handleParsingError(name, str, err)
		if err == nil {
			return uint16(n)
		}
	}
	for _, value := range defaultValue {
		return value
	}
	return 0
}

func (g *myGetter) GetUint32(name string, defaultValue ...uint32) uint32 {
	const (
		base    = 10
		bitSize = 32
	)
	str, ok := g.getString(name)
	if ok {
		n, err := strconv.ParseUint(str, base, bitSize)
		g.handleParsingError(name, str, err)
		if err == nil {
			return uint32(n)
		}
	}
	for _, value := range defaultValue {
		return value
	}
	return 0
}

func (g *myGetter) GetUint64(name string, defaultValue ...uint64) uint64 {
	const (
		base    = 10
		bitSize = 64
	)
	str, ok := g.getString(name)
	if ok {
		n, err := strconv.ParseUint(str, base, bitSize)
		g.handleParsingError(name, str, err)
		if err == nil {
			return n
		}
	}
	for _, value := range defaultValue {
		return value
	}
	return 0
}

func (g *myGetter) ListItems(prefix, suffix string) []string {

	dst := make([]string, 0)
	src := g.table.Export(nil)

	len1 := len(prefix)
	len2 := len(suffix)

	for k, v := range src {
		len3 := len(k)
		if len3 > len1+len2 {
			if strings.HasPrefix(k, prefix) && strings.HasSuffix(k, suffix) {
				if v != "" {
					id := k[len1 : len3-len2]
					dst = append(dst, id)
				}
			}
		}
	}
	sort.Strings(dst)
	return dst
}
