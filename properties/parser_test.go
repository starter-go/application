package properties

import "testing"

func TestParse(t *testing.T) {

	t1 := make(map[string]string)

	t1[""] = "0"
	t1["foo"] = "foo1"
	t1["bar"] = "bar1"
	t1["core.a"] = "1"
	t1["core.b"] = "2"
	t1["core.c"] = "3"
	t1["remote.a.name"] = "a"
	t1["remote.a.url"] = "git@foo.com:bar"
	t1["remote.bbb.name"] = "bbb"
	t1["remote.bbb.url"] = "https://foo.com/bar"
	t1["branch.zzz.up"] = "+++"
	t1["branch.zzz.down"] = "---z"
	t1["branch.yyy.zzz.down"] = "--y-z"
	t1["branch.xxx.yyy.zzz.down"] = "-x-y-z"

	// table 2
	t2 := NewTable(nil)
	t2.Import(t1)
	s2 := Format(t2, FormatWithGroups)
	t.Logf("table[2]: \n %s", s2)

	// table 3
	t3, err := Parse(s2, nil)
	if err != nil {
		t.Error(err)
		return
	}
	s3 := Format(t3, FormatWithGroups)
	t.Logf("table[3]: \n %s", s3)
}
