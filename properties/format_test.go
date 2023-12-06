package properties

import "testing"

func TestFormat(t *testing.T) {

	t1 := make(map[string]string)

	t1["core.a"] = "1"
	t1["core.b"] = "2"
	t1["core.c"] = "3"
	t1["remote.a.name"] = "a"
	t1["remote.a.url"] = "git@foo.com:bar"
	t1["remote.bbb.name"] = "bbb"
	t1["remote.bbb.url"] = "https://foo.com/bar"
	t1["branch.zzz.up"] = "+++"
	t1["branch.zzz.down"] = "---"

	t2 := NewTable(nil)
	t2.Import(t1)
	s1 := Format(t2)
	s2 := Format(t2, FormatWithGroups)
	t.Logf("format.props.style1:\n%s", s1)
	t.Logf("format.props.style2:\n%s", s2)
}
