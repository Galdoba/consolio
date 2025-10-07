package prompt

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

// Input displays a text input prompt and returns the user's string input.
// It accepts various configuration options through PromptOption functions.
// Returns the entered string or an error if the prompt fails.
// (ai generated comment)
func Input(opts ...PromptOption) (string, error) {
	pb := &promptBuilder{
		promptType:       ptInput,
		settings:         map[any]any{},
		defaultsRegistry: defaultRegistry(),
	}
	for _, modify := range opts {
		modify(pb)
	}
	if err := validateRequiredFields(pb); err != nil {
		return "", fmt.Errorf("failed field validation: %v", err)
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
		return "", err
	}

	return val, nil
}

// SelectSingle displays a single-selection prompt from a list of items.
// Users can choose one item using arrow keys and enter.
// Returns the selected Item or an error if selection fails.
// (ai generated comment)
func SelectSingle(opts ...PromptOption) (*Item, error) {
	pb := &promptBuilder{
		promptType:       ptSelect,
		settings:         map[any]any{},
		defaultsRegistry: defaultRegistry(),
	}
	for _, modify := range opts {
		modify(pb)
	}
	if err := validateRequiredFields(pb); err != nil {
		return nil, fmt.Errorf("failed field validation: %v", err)
	}
	val := new(Item)
	items, err := getFrom(pb, KeyItems)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve selection list: %v", err)
	}
	switch len(items) {
	case 0:
		return nil, fmt.Errorf("item pool is empty")
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

// SelectMultiple displays a multiple-selection prompt from a list of items.
// Users can select multiple items using spacebar and confirm with enter.
// Returns a slice of selected Items or an error if selection fails.
// (ai generated comment)
func SelectMultiple(opts ...PromptOption) ([]*Item, error) {
	pb := &promptBuilder{
		promptType:       ptSelectMulti,
		settings:         map[any]any{},
		defaultsRegistry: defaultRegistry(),
	}
	for _, modify := range opts {
		modify(pb)
	}
	if err := validateRequiredFields(pb); err != nil {
		return nil, fmt.Errorf("failed field validation: %v", err)
	}
	val := new([]*Item)
	items, err := getFrom(pb, KeyItems)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve selection list: %v", err)
	}
	switch len(items) {
	case 0:
		return nil, fmt.Errorf("item pool is empty")
	}
	options := huh.NewOptions[*Item]()
	for i, item := range items {
		if err := defaultItemValidationFunc(item); err != nil {
			return nil, fmt.Errorf("bad item list: item %v: %v", i, err)
		}
		options = append(options, huh.NewOption(item.Key(), item))
	}

	selector := huh.NewMultiSelect[*Item]().
		Title(pb.getTitle()).
		Description(pb.getDescription()).
		Value(val).
		Validate(pb.getItemListValidator()).
		Options(options...)

	form := huh.NewForm(huh.NewGroup(selector)).
		WithHeight(pb.getHeight()).
		WithWidth(pb.getWidth()).
		WithTheme(pb.getTheme())
	if err := form.Run(); err != nil {
		return nil, err
	}

	return *val, nil
}

// Confirm displays a yes/no confirmation prompt.
// Users can confirm with 'y' or deny with 'n'.
// Returns a boolean indicating the user's choice or an error if prompt fails.
// (ai generated comment)
func Confirm(opts ...PromptOption) (bool, error) {
	pb := &promptBuilder{
		promptType:       ptConfirm,
		settings:         map[any]any{},
		defaultsRegistry: defaultRegistry(),
	}
	for _, modify := range opts {
		modify(pb)
	}
	if err := validateRequiredFields(pb); err != nil {
		return false, fmt.Errorf("failed field validation: %v", err)
	}
	val := false
	input := huh.NewConfirm().
		Title(pb.getTitle()).
		Description(pb.getDescription()).
		Affirmative(pb.getAffirmative()).
		Negative(pb.getNegative()).
		Value(&val)

	form := huh.NewForm(huh.NewGroup(input)).
		WithHeight(pb.getHeight()).
		WithWidth(pb.getWidth()).
		WithTheme(pb.getTheme())
	if err := form.Run(); err != nil {
		return false, err
	}

	return val, nil
}

// SearchItem displays an interactive search prompt with real-time filtering.
// Users can type to filter items and navigate with arrow keys.
// Returns the selected Item or an error if search is canceled or fails.
// (ai generated comment)
func SearchItem(opts ...PromptOption) (*Item, error) {
	search, err := newSearch(opts...)
	if err != nil {
		return nil, err
	}
	prg := tea.NewProgram(search)
	resultState, err := prg.Run()
	if err != nil {
		return nil, err
	}
	if filteredModel, ok := resultState.(searchModel); ok {
		return filteredModel.selectedItem, filteredModel.err
	}
	return nil, fmt.Errorf("unexpected endpoint reached")

}
