package boot

import "github.com/starter-go/base/safe"

// Options 包含启动选项
type Options struct {
	Args []string
	Mode safe.Mode
}
