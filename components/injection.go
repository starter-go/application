package components

import "context"

// Injection 表示一个依赖注入上下文
type Injection interface {
	context.Context

	Select(selector Selector) ([]Instance, error)

	SelectOne(selector Selector) (Instance, error)

	GetByID(id ID) (Instance, error)

	GetApplicationContext() context.Context

	GetProperty(selector Selector) (string, error)

	GetWithHolder(holder Holder) (Instance, error)

	Ext() InjectionExt
}

// InjectionExt 是 Injection 的扩展部分
type InjectionExt interface {
	Injection() Injection

	GetContext() context.Context // 返回一个 application.Context

	GetString(selector Selector) string
	GetBool(selector Selector) bool
	GetRune(selector Selector) rune
	GetByte(selector Selector) byte
	GetAny(selector Selector) any

	GetInt(selector Selector) int
	GetInt8(selector Selector) int8
	GetInt16(selector Selector) int16
	GetInt32(selector Selector) int32
	GetInt64(selector Selector) int64

	GetUint(selector Selector) uint
	GetUint8(selector Selector) uint8
	GetUint16(selector Selector) uint16
	GetUint32(selector Selector) uint32
	GetUint64(selector Selector) uint64

	GetFloat32(selector Selector) float32
	GetFloat64(selector Selector) float64

	GetComponent(selector Selector) any

	ListComponents(selector Selector) []any
}
