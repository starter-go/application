package properties

import (
	"fmt"
	"testing"
)

func TestGetterListItems(t *testing.T) {

	const (
		prefix = "item."
		suffix = ".tag"
	)

	t1 := NewTable(nil)

	for i := 0; i < 8; i++ {
		name12 := fmt.Sprintf("item.x-item-%d", i)
		t1.SetProperty(name12+".name", "name_xxx")
		t1.SetProperty(name12+".value", "value_xxx")
		t1.SetProperty(name12+".tag", "tag_xxx")
		t1.SetProperty(name12+".label", "label_xxx")
	}

	formatted := Format(t1)
	fmt.Println(formatted)

	g1 := t1.Getter()
	list := g1.ListItems(prefix, suffix)
	for idx, id := range list {
		fmt.Printf("item[%d].id = %s\n", idx, id)
	}

}
