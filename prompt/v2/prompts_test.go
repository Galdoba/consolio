package v2

import (
	"fmt"
	"testing"
)

func TestInput(t *testing.T) {
	fmt.Println(NewSelect(
		FromItems([]*Item{
			NewItem("one", 1.2),
			NewItem("Two", "II"),
			NewItem("THEEE", 3),
			NewItem("FOUR"),
			NewItem("Five"),
			NewItem("6")}),
		WithDescription("selectuion"), WithTitle("me is selection"), WithItemValidator(NoNumbers)))

	// fmt.Println(NewSelect(WithTitle("this is select"), FromItems(
	// 	NewItem("item 1"),
	// 	NewItem("item 2"),
	// )))
	//

}
