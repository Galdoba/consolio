package prompt

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

const (
	formInput   = "input"
	formSingle  = "single"
	formMulti   = "multi"
	formConfirm = "confirm"
	formSearch  = "search"
)

func Input(opts ...PromptOption) (string, error) {
	pb, err := newPromptBuilder(formInput)
	if err != nil {
		return "", fmt.Errorf("failed to create prompt builder: %v", err)
	}
	for _, modify := range opts {
		modify(pb)
	}
	inputPrompt := huh.NewInput().
		Title(pb.title).
		Description(pb.description).
		Prompt(pb.prompt).
		Placeholder(pb.placeholder).
		Validate(pb.textValidator).
		Value(&pb.userInput)

	form := huh.NewForm(huh.NewGroup(inputPrompt)).
		WithTheme(pb.theme).
		WithHeight(pb.height).
		WithWidth(pb.width)

	if err := form.Run(); err != nil {
		return "", err
	}
	return pb.userInput, nil
}

func SelectSingle(opts ...PromptOption) (*Item, error) {
	pb, err := newPromptBuilder(formSingle)
	if err != nil {
		return nil, fmt.Errorf("failed to create prompt builder: %v", err)
	}
	for _, modify := range opts {
		modify(pb)
	}
	singlePrompt := huh.NewSelect[*Item]().
		Title(pb.title).
		Description(pb.description).
		Inline(pb.selectInline).
		Validate(pb.itemValidator).
		Value(&pb.outputSingle)

	items := []huh.Option[*Item]{}
	for _, item := range pb.items {
		items = append(items, huh.NewOption(item.Key, item))
	}
	singlePrompt.Options(items...)
	if len(pb.items) == 1 {
		return pb.items[0], pb.itemValidator(pb.items[0])
	}

	form := huh.NewForm(huh.NewGroup(singlePrompt)).
		WithTheme(pb.theme).
		WithHeight(pb.height).
		WithWidth(pb.width)

	if err := form.Run(); err != nil {
		return nil, err
	}
	return pb.outputSingle, nil
}

func SelectMultiple(opts ...PromptOption) ([]*Item, error) {
	pb, err := newPromptBuilder(formMulti)
	if err != nil {
		return nil, fmt.Errorf("failed to create prompt builder: %v", err)
	}
	for _, modify := range opts {
		modify(pb)
	}
	multiPrompt := huh.NewMultiSelect[*Item]().
		Title(pb.title).
		Description(pb.description).
		Validate(pb.itemSetValidator).
		Value(&pb.outputMultiple)

	items := []huh.Option[*Item]{}
	for _, item := range pb.items {
		items = append(items, huh.NewOption(item.Key, item))
	}
	multiPrompt.Options(items...)

	form := huh.NewForm(huh.NewGroup(multiPrompt)).
		WithTheme(pb.theme).
		WithHeight(pb.height).
		WithWidth(pb.width)

	if err := form.Run(); err != nil {
		return nil, err
	}
	return pb.outputMultiple, nil
}

func Confirm(opts ...PromptOption) (bool, error) {
	pb, err := newPromptBuilder(formConfirm)
	if err != nil {
		return false, fmt.Errorf("failed to create prompt builder: %v", err)
	}
	for _, modify := range opts {
		modify(pb)
	}

	confirmPrompt := huh.NewConfirm().
		Title(pb.title).
		Description(pb.description).
		Affirmative(pb.confirmAffirmative).
		Negative(pb.confirmNegative).
		Value(&pb.outputConfirmation)

	form := huh.NewForm(huh.NewGroup(confirmPrompt)).
		WithTheme(pb.theme).
		WithHeight(pb.height).
		WithWidth(pb.width)

	if err := form.Run(); err != nil {
		return false, err
	}
	return pb.outputConfirmation, nil
}

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
		return filteredModel.selectedItem, nil
	}
	return nil, fmt.Errorf("unexpected endpoint reached")

}
