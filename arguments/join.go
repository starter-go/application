package arguments

import "strings"

// Join 把一组命令行参数连接成一个 string
func Join(args []string) string {
	dst := make([]string, len(args))
	for i, item := range args {
		dst[i] = prepareArgumentForJoin(item)
	}
	return strings.Join(dst, " ")
}

func prepareArgumentForJoin(part string) string {

	if part == "" {
		return "\"\""
	}

	countAll := 0
	countStr1 := 0 // count the '\''
	countStr2 := 0 // count the '"'

	chs := []rune(part)
	for _, ch := range chs {
		if ch == ' ' || ch == '\t' {
			countAll++
		} else if ch == '"' {
			countAll++
			countStr2++
		} else if ch == '\'' {
			countAll++
			countStr1++
		}
	}

	if countAll < 1 {
		return part // 不包含特殊字符
	}
	mark := "\""
	if countStr1 > 0 {
		if countStr2 > 0 {
			return "[bad_arg_string]"
		}
	} else {
		if countStr2 > 0 {
			mark = "'"
		}
	}
	return mark + part + mark
}
