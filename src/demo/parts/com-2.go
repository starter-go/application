package parts

import "io"

type Com2 struct {
	F1 string
	F2 int
	F3 bool
	F4 float64
	F5 rune
	F6 byte
	F7 io.Closer
	F8 []*Com1
}
