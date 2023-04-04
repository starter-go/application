package application

import (
	"context"
	"io"
)

// Context 表示一个 starter 应用上下文
type Context interface {
	io.Closer
	context.Context

	NewChild() Context
}
