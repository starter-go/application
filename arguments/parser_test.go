package arguments

import "testing"

func TestParseCommandLine(t *testing.T) {
	rows := make([]string, 0)

	rows = append(rows, "")
	rows = append(rows, " cmd.exe /c \"wt.exe\"  new-tab -p \"Ubuntu-18.04\"  focus-tab -t 1 ")
	rows = append(rows, "git push --force-with-lease=master:base master:master")

	for _, row := range rows {
		args, err := Parse(row)
		if err != nil {
			t.Error(err)
		}
		t.Logf("command-line: %s", row)
		t.Logf("        args:")
		for i, a := range args {
			t.Logf("    args[%d] = %s", i, a)
		}
	}
}
