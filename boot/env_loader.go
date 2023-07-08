package boot

import (
	"os"
	"strings"
)

type envLoader struct {
	b *Bootstrap
}

func (inst *envLoader) Load() error {
	dst := inst.b.collections.Environment
	src1 := inst.loadFromSystem()
	src2 := dst.Export(nil)
	dst.Import(src1)
	dst.Import(src2)
	return nil
}

func (inst *envLoader) loadFromSystem() map[string]string {
	dst := make(map[string]string)
	src := os.Environ()
	for _, item := range src {
		i := strings.IndexRune(item, '=')
		if i >= 0 {
			key := strings.TrimSpace(item[0:i])
			val := strings.TrimSpace(item[i+1:])
			dst[key] = val
		}
	}
	return dst
}
