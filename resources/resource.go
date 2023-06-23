package resources

import "io"

// Resource 表示一个资源
type Resource interface {
	Path() string

	SimpleName() string

	Size() int64

	ReadBinary() ([]byte, error)

	ReadText() (string, error)

	Open() (io.ReadCloser, error)
}
