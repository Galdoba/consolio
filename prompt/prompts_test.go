package prompt_test

import (
	"fmt"
	"testing"

	"github.com/Galdoba/consolio/prompt"
)

// func TestGetInput(t *testing.T) {
// 	answer, err := prompt.Input(prompt.WithTextValidator(prompt.Number))
// 	fmt.Println(answer, err)
// 	fmt.Println("")
// }

// func TestSingleSelect(t *testing.T) {
// 	chosen, err := prompt.SelectSingle(prompt.WithDescription("abbacab"),
// 		prompt.FromItems(
// 			prompt.NewItem("one", 1),
// 			prompt.NewItem("two", "II"),
// 			prompt.NewItem("theree", "три"),
// 			prompt.NewItem("four", "****"),
// 		),
// 		prompt.WithInlineSelection(false),
// 		prompt.WithTheme(huh.ThemeBase16()))
// 	fmt.Println(err)
// 	fmt.Println(chosen.Key, fmt.Sprintf("%v", chosen.Data))
// 	fmt.Println("")
// }

// func TestConfirm(t *testing.T) {
// 	confirmed, err := prompt.Confirm(prompt.WithTitle("me need confirm"), prompt.WithDescription("descr confr"))
// 	fmt.Println(err)
// 	fmt.Println(confirmed)

func TestSelectMultiple(t *testing.T) {
	items := []*prompt.Item{}
	for i := 0; i < 50; i++ {
		items = append(items, prompt.CreateItem(fmt.Sprintf("object %v", i)))
	}
	chosen, err := prompt.SelectMultiple(
		prompt.FromItems(items...),
		prompt.WithDescription("this is test desctiption for multiselect\n with multiline"),
	)
	fmt.Println(err)
	for i, selected := range chosen {
		fmt.Println(i, selected.Key, fmt.Sprintf("%v", selected.Data))

	}
	fmt.Println("")
}

func TestSearch(t *testing.T) {
	items := []*prompt.Item{}
	for i := 0; i < 500000; i++ {
		items = append(items, prompt.CreateItem(fmt.Sprintf("object %v", i)))
	}

	item, err := prompt.SearchItem(
		prompt.FromItems(items...),
	)
	fmt.Println(err)
	fmt.Println(item)
}
