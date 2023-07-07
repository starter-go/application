package application

import (
	"github.com/starter-go/application/components"
)

// Injection 表示一个注入上下文
type Injection interface {
	components.Injection

	Ext() InjectionExt

	GetContext() Context

	LifeManager() LifeManager

	Complete() error
}

// InjectionExt 是 Injection 的扩展部分
type InjectionExt interface {
	GetInjection() Injection

	GetContext() Context

	GetString(selector components.Selector) string
	GetBool(selector components.Selector) bool
	GetRune(selector components.Selector) rune
	GetByte(selector components.Selector) byte
	GetAny(selector components.Selector) any

	GetInt(selector components.Selector) int
	GetInt8(selector components.Selector) int8
	GetInt16(selector components.Selector) int16
	GetInt32(selector components.Selector) int32
	GetInt64(selector components.Selector) int64

	GetUint(selector components.Selector) uint
	GetUint8(selector components.Selector) uint8
	GetUint16(selector components.Selector) uint16
	GetUint32(selector components.Selector) uint32
	GetUint64(selector components.Selector) uint64

	GetFloat32(selector components.Selector) float32
	GetFloat64(selector components.Selector) float64

	GetComponent(selector components.Selector) any

	ListComponents(selector components.Selector) []any
}
