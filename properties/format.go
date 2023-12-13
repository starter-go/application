package properties

// Format 把 Table 格式化为字符串
func Format(t Table, options ...FormatOptionsF) string {

	if t == nil {
		return ""
	}

	opt := new(FormatOptions)
	for _, f := range options {
		f(opt)
	}

	if opt.useGroups {
		formatter := new(groupsFormatter)
		return formatter.fmt(t)
	}

	formatter := new(simpleFormatter)
	return formatter.fmt(t)
}

////////////////////////////////////////////////////////////////////////////////

// FormatOptions ...
type FormatOptions struct {
	useGroups bool
}

// FormatOptionsF ...
type FormatOptionsF func(opt *FormatOptions)

// FormatWithGroups 启用分组功能
func FormatWithGroups(opt *FormatOptions) {
	opt.useGroups = true
}

////////////////////////////////////////////////////////////////////////////////
