package v2

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

// NewInput creates a text input prompt
func NewInput(opts ...PromptOption) (string, error) {
	pb := &promptBuilder{
		promptType:       ptInput,
		settings:         map[any]any{},
		defaultsRegistry: defaultRegistry(),
	}
	for _, modify := range opts {
		modify(pb)
	}
	val := ""
	input := huh.NewInput().
		Title(pb.getTitle()).
		Description(pb.getDescription()).
		Prompt(pb.getPrompt()).
		Placeholder(pb.getPlaceholder()).
		Validate(pb.getStringValidator()).
		Value(&val)

	form := huh.NewForm(huh.NewGroup(input)).
		WithHeight(pb.getHeight()).
		WithWidth(pb.getWidth()).
		WithTheme(pb.getTheme())
	if err := form.Run(); err != nil {
		return "no val", err
	}

	return val, nil
}

// NewSelect creates a single selection prompt
func NewSelect(opts ...PromptOption) (*Item, error) {
	pb := &promptBuilder{
		promptType:       ptSelect,
		settings:         map[any]any{},
		defaultsRegistry: defaultRegistry(),
	}
	for _, modify := range opts {
		modify(pb)
	}
	val := new(Item)
	items := mustGet[[]*Item](pb, KeyItems)
	switch len(items) {
	case 0:
		return nil, fmt.Errorf("nothing to select from")
	case 1:
		return items[0], nil
	}
	options := huh.NewOptions[*Item]()
	for i, item := range items {
		if err := defaultItemValidationFunc(item); err != nil {
			return nil, fmt.Errorf("bad item list: item %v: %v", i, err)
		}
		options = append(options, huh.NewOption(item.Key(), item))
	}

	selector := huh.NewSelect[*Item]().
		Title(pb.getTitle()).
		Description(pb.getDescription()).
		Value(&val).
		Validate(pb.getItemValidator()).
		Options(options...)

	form := huh.NewForm(huh.NewGroup(selector)).
		WithHeight(pb.getHeight()).
		WithWidth(pb.getWidth()).
		WithTheme(pb.getTheme())
	if err := form.Run(); err != nil {
		return nil, err
	}

	return val, nil
}

// NewSelectMulti creates a multiselection prompt
// func NewSelectMulti(opts ...PromptOption) (*Item, error) {
// 	pb := &promptBuilder{
// 		promptType:       ptSelect,
// 		settings:         map[any]any{},
// 		defaultsRegistry: defaultRegistry(),
// 	}
// 	for _, modify := range opts {
// 		modify(pb)
// 	}
// 	val := new(Item)
// 	items := pb.getItems() // mustGet[[]*Item](pb, KeyItems)
// 	switch len(items) {
// 	case 0:
// 		return nil, fmt.Errorf("nothing to select from")
// 	case 1:
// 		return items[0], nil
// 	}
// 	options := huh.NewOptions[*Item]()
// 	validatorFunc := pb.getItemValidator()
// 	validatorFunc(items[0])
// 	for _, item := range items {
// 		fmt.Println("item", item.key)
// 		if err := validatorFunc(item); err != nil {
// 			return nil, err
// 		}
// 		options = append(options, huh.NewOption(item.Key(), item))
// 	}

// 	selector := huh.NewSelect[*Item]().
// 		Title(pb.getTitle()).
// 		Description(pb.getDescription()).
// 		Validate(pb.getItemValidator()).
// 		Value(&val).
// 		Options(options...)

// 	form := huh.NewForm(huh.NewGroup(selector)).
// 		WithHeight(pb.getHeight()).
// 		WithWidth(pb.getWidth()).
// 		WithTheme(pb.getTheme())
// 	if err := form.Run(); err != nil {
// 		return nil, err
// 	}

// 	return val, nil
// }
