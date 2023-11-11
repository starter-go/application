package arguments

import (
	"fmt"
	"strings"
)

// Parse 函数把一个命令行字符串解析为 []string 的形式
func Parse(line string) ([]string, error) {
	p := &parser{}
	return p.parse(line)
}

////////////////////////////////////////////////////////////////////////////////

type parser struct {
}

func (inst *parser) parse(line string) ([]string, error) {
	reader := &parserReader{}
	reader.init(line)
	dst := make([]string, 0)
	for {
		reader.skipSpace()
		if reader.eof() {
			break
		}
		part, err := reader.read()
		if err != nil {
			return nil, err
		}
		dst = append(dst, part)
	}
	return dst, nil
}

////////////////////////////////////////////////////////////////////////////////

type parserReader struct {
	buffer []rune
	readAt int // 这是一个指向 buffer 的 index, 表示当前读取的位置
	size   int // size = len(buffer)
}

func (inst *parserReader) init(line string) {
	data := []rune(line)
	inst.readAt = 0
	inst.buffer = data
	inst.size = len(data)
}

func (inst *parserReader) eof() bool {
	return inst.size <= inst.readAt
}

func (inst *parserReader) read() (string, error) {
	i := inst.readAt
	size := inst.size
	var ch rune
	if 0 <= i && i < size {
		ch = inst.buffer[i]
	} else {
		return "", fmt.Errorf("the read position at buffer([]rune) is overflow")
	}
	if ch == '\'' || ch == '"' {
		return inst.readAsString(ch)
	}
	return inst.readAsWord()
}

func (inst *parserReader) readAsWord() (string, error) {
	i := inst.readAt
	data := inst.buffer
	size := inst.size
	b := &strings.Builder{}
	for ; i < size; i++ {
		ch := data[i]
		if inst.isWordRune(ch) {
			b.WriteRune(ch)
		} else {
			break
		}
	}
	inst.readAt = i
	return b.String(), nil
}

func (inst *parserReader) isWordRune(ch1 rune) bool {
	const ending = " \t\r\n'\"" // 这些字符不会出现在 word 中
	return !strings.ContainsRune(ending, ch1)
}

func (inst *parserReader) readAsString(mark rune) (string, error) {
	i := inst.readAt
	data := inst.buffer
	size := inst.size
	b := &strings.Builder{}
	countMark := 0
	for ; i < size; i++ {
		ch := data[i]
		if ch == mark {
			countMark++
			if countMark == 1 {
				continue
			} else if countMark == 2 {
				i++
				break // close string
			}
		}
		b.WriteRune(ch)
	}
	inst.readAt = i
	if countMark == 2 {
		return b.String(), nil // string has been open & closed
	}
	return "", fmt.Errorf("the argument string is not closed:[%s]", b.String())
}

func (inst *parserReader) skipSpace() {
	i := inst.readAt
	data := inst.buffer
	size := inst.size
	for ; i < size; i++ {
		ch := data[i]
		if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
			continue
		} else {
			break
		}
	}
	inst.readAt = i
}

////////////////////////////////////////////////////////////////////////////////
