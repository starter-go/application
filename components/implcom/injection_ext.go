package implcom

import (
	"context"
	"strconv"

	"github.com/starter-go/application/components"
)

type injectionExt struct {
	injection components.Injection
}

func (inst *injectionExt) init() components.InjectionExt {
	return inst
}

func (inst *injectionExt) Injection() components.Injection {
	return inst.injection
}

func (inst *injectionExt) GetString(selector components.Selector) string {
	value, err := inst.injection.GetProperty(selector)
	if err != nil {
		panic(err)
	}
	return value
}

func (inst *injectionExt) GetBool(selector components.Selector) bool {
	str := inst.GetString(selector)
	value, err := strconv.ParseBool(str)
	if err != nil {
		panic(err)
	}
	return value
}

func (inst *injectionExt) GetRune(selector components.Selector) rune {
	str := inst.GetString(selector)
	array := []rune(str)
	if array != nil {
		if len(array) == 1 {
			return array[0]
		}
	}
	panic("bad rune property: '" + str + "'")
}

func (inst *injectionExt) GetByte(selector components.Selector) byte {
	n := inst.GetUint64(selector)
	return byte(n & 0x00ff)
}

func (inst *injectionExt) GetInt(selector components.Selector) int {
	str := inst.GetString(selector)
	value, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return value
}

func (inst *injectionExt) GetInt8(selector components.Selector) int8 {
	n := inst.GetInt64(selector)
	return int8(n)
}

func (inst *injectionExt) GetInt16(selector components.Selector) int16 {
	n := inst.GetInt64(selector)
	return int16(n)
}

func (inst *injectionExt) GetInt32(selector components.Selector) int32 {
	n := inst.GetInt64(selector)
	return int32(n)
}

func (inst *injectionExt) GetInt64(selector components.Selector) int64 {
	str := inst.GetString(selector)
	value, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return value
}

func (inst *injectionExt) GetUint(selector components.Selector) uint {
	n := inst.GetUint64(selector)
	return uint(n)
}

func (inst *injectionExt) GetUint8(selector components.Selector) uint8 {
	n := inst.GetUint64(selector)
	return uint8(n)
}

func (inst *injectionExt) GetUint16(selector components.Selector) uint16 {
	n := inst.GetUint64(selector)
	return uint16(n)
}

func (inst *injectionExt) GetUint32(selector components.Selector) uint32 {
	n := inst.GetUint64(selector)
	return uint32(n)
}

func (inst *injectionExt) GetUint64(selector components.Selector) uint64 {
	str := inst.GetString(selector)
	value, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return value
}

func (inst *injectionExt) GetFloat32(selector components.Selector) float32 {
	str := inst.GetString(selector)
	value, err := strconv.ParseFloat(str, 32)
	if err != nil {
		panic(err)
	}
	return float32(value)
}

func (inst *injectionExt) GetFloat64(selector components.Selector) float64 {
	str := inst.GetString(selector)
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}
	return value
}

func (inst *injectionExt) GetComponent(selector components.Selector) any {
	i, err := inst.injection.SelectOne(selector)
	if err != nil {
		panic(err)
	}
	return i.Get()
}

func (inst *injectionExt) ListComponents(selector components.Selector) []any {
	src, err := inst.injection.Select(selector)
	if err != nil {
		panic(err)
	}
	dst := make([]any, 0)
	for _, i := range src {
		dst = append(dst, i.Get())
	}
	return dst
}

// GetAny 是 GetComponent 的别名
func (inst *injectionExt) GetAny(selector components.Selector) any {
	return inst.GetComponent(selector)
}

func (inst *injectionExt) GetApplicationContext() context.Context {
	return inst.injection.GetApplicationContext()
}
