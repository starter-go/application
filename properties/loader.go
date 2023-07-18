package properties

import (
	"github.com/starter-go/application/resources"
	"github.com/starter-go/base/safe"
)

// LoadFromResource 从资源中加载属性表
func LoadFromResource(path string, res resources.Table, mode safe.Mode) (Table, error) {
	text, err := res.ReadText(path)
	if err != nil {
		return nil, err
	}
	return Parse(text, mode)
}
