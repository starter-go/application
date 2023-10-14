package properties

import (
	"strconv"
	"strings"
)

// Setter 属性值设置器
type Setter interface {
	SetBool(name string, value bool)
	SetByte(name string, value byte)
	SetRune(name string, value rune)
	SetString(name string, value string)
	SetStringList(name string, value []string)

	SetFloat32(name string, value float32)
	SetFloat64(name string, value float64)

	SetInt(name string, value int)
	SetInt8(name string, value int8)
	SetInt16(name string, value int16)
	SetInt32(name string, value int32)
	SetInt64(name string, value int64)

	SetUint(name string, value uint)
	SetUint8(name string, value uint8)
	SetUint16(name string, value uint16)
	SetUint32(name string, value uint32)
	SetUint64(name string, value uint64)
}

////////////////////////////////////////////////////////////////////////////////

type mySetter struct {
	table Table
}

func (s *mySetter) _impl(a Setter) {
	a = s
}

func (s *mySetter) SetBool(name string, value bool) {
	str := strconv.FormatBool(value)
	s.table.SetProperty(name, str)
}

func (s *mySetter) SetByte(name string, value byte) {
	var buffer [1]byte
	buffer[0] = value
	str := string(buffer[:])
	s.table.SetProperty(name, str)
}

func (s *mySetter) SetRune(name string, value rune) {
	str := string(value)
	s.table.SetProperty(name, str)
}

func (s *mySetter) SetString(name string, value string) {
	s.table.SetProperty(name, value)
}

func (s *mySetter) SetStringList(name string, value []string) {
	b := strings.Builder{}
	for i, item := range value {
		if i > 0 {
			b.WriteRune(',')
		}
		b.WriteString(item)
	}
	s.table.SetProperty(name, b.String())
}

func (s *mySetter) SetFloat32(name string, value float32) {
	const (
		fmt     = 'f'
		prec    = 10
		bitSize = 32
	)
	str := strconv.FormatFloat(float64(value), fmt, prec, bitSize)
	s.table.SetProperty(name, str)
}

func (s *mySetter) SetFloat64(name string, value float64) {
	const (
		fmt     = 'f'
		prec    = 20
		bitSize = 64
	)
	str := strconv.FormatFloat(value, fmt, prec, bitSize)
	s.table.SetProperty(name, str)
}

func (s *mySetter) SetInt(name string, value int) {
	const base = 10
	str := strconv.FormatInt(int64(value), base)
	s.table.SetProperty(name, str)
}

func (s *mySetter) SetInt8(name string, value int8) {
	const base = 10
	str := strconv.FormatInt(int64(value), base)
	s.table.SetProperty(name, str)
}

func (s *mySetter) SetInt16(name string, value int16) {
	const base = 10
	str := strconv.FormatInt(int64(value), base)
	s.table.SetProperty(name, str)
}

func (s *mySetter) SetInt32(name string, value int32) {
	const base = 10
	str := strconv.FormatInt(int64(value), base)
	s.table.SetProperty(name, str)
}

func (s *mySetter) SetInt64(name string, value int64) {
	const base = 10
	str := strconv.FormatInt(value, base)
	s.table.SetProperty(name, str)
}

func (s *mySetter) SetUint(name string, value uint) {
	const base = 10
	str := strconv.FormatUint(uint64(value), base)
	s.table.SetProperty(name, str)
}

func (s *mySetter) SetUint8(name string, value uint8) {
	const base = 10
	str := strconv.FormatUint(uint64(value), base)
	s.table.SetProperty(name, str)
}

func (s *mySetter) SetUint16(name string, value uint16) {
	const base = 10
	str := strconv.FormatUint(uint64(value), base)
	s.table.SetProperty(name, str)
}

func (s *mySetter) SetUint32(name string, value uint32) {
	const base = 10
	str := strconv.FormatUint(uint64(value), base)
	s.table.SetProperty(name, str)
}

func (s *mySetter) SetUint64(name string, value uint64) {
	const base = 10
	str := strconv.FormatUint(value, base)
	s.table.SetProperty(name, str)
}
