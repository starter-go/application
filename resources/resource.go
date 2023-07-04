package resources

import (
	"io"
	"strings"
)

// Resource 表示一个资源
type Resource interface {
	Path() string

	SimpleName() string

	Size() int64

	ReadBinary() ([]byte, error)

	ReadText() (string, error)

	Open() (io.ReadCloser, error)
}

// 标准化资源路径
func normalizePath(path string) string {
	const keyword = "://"
	i := strings.Index(path, keyword)
	if i > 0 {
		path = path[i+len(keyword):]
	}
	elements := strings.Split(path, "/")
	builder := strings.Builder{}
	for _, el := range elements {
		el = strings.TrimSpace(el)
		if el != "" {
			builder.WriteString("/")
			builder.WriteString(el)
		}
	}
	return builder.String()
}
