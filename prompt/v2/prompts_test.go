package v2

import (
	"fmt"
	"testing"
)

func TestPrompt(t *testing.T) {
	fmt.Println(SearchItem(FromItems([]*Item{NewItem("Item 1"), NewItem("второй")})))
}
