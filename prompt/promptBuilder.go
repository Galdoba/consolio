package prompt

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/x/term"
)

type promptBuilder struct {
	promptType string
	input      *huh.Input
	single     *huh.Select[*Item]
	multi      *huh.MultiSelect[*Item]
	confirm    *huh.Confirm
	// search     *SearchModel

	// Общие поля для всех типов промптов
	theme            *huh.Theme
	items            []*Item
	title            string
	description      string
	prompt           string
	placeholder      string
	helpText         string
	textValidator    func(string) error
	itemValidator    func(*Item) error
	itemSetValidator func([]*Item) error
	height           int
	width            int

	// Специфичные поля
	selectInline        bool
	confirmAffirmative  string
	confirmNegative     string
	searchMaxDisplay    int
	searchAutoSelect    bool
	searchCaseSensitive bool
	searchHelpText      string
	searchInitialFilter string

	// Поля вывода
	userInput          string
	outputSingle       *Item
	outputMultiple     []*Item
	outputConfirmation bool
}

func newPromptBuilder(form string) (*promptBuilder, error) {
	width, height, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		return nil, fmt.Errorf("failed to recieve console dimentions: %v", err)
	}

	pb := &promptBuilder{
		promptType:         form,
		description:        defaultDescription,
		prompt:             defaultPrompt,
		placeholder:        defaultPlaceholder,
		theme:              huh.ThemeBase16(),
		textValidator:      defaultTextValiidatorFunc,
		itemValidator:      defaultItemValiidatorFunc,
		itemSetValidator:   defaultItemSetValiidatorFunc,
		selectInline:       false,
		confirmAffirmative: "Yes",
		confirmNegative:    "No",
		// searchMaxDisplay:    defaultDisplayHeight,
		searchAutoSelect:    true,
		searchCaseSensitive: false,
		// searchHelpText:      defaultHelpText,
		title: defaultPromptTitle(form),
		// input:               &huh.Input{},
		// single:              &huh.Select[*Item]{},
		// multi:               &huh.MultiSelect[*Item]{},
		// confirm:             &huh.Confirm{},
		// items:               []*Item{},
		helpText:            "",
		height:              height - 1,
		width:               width,
		searchMaxDisplay:    0,
		searchHelpText:      "",
		searchInitialFilter: "",
		userInput:           "",
		// outputSingle:        &Item{},
		// outputMultiple:      []*Item{},
		// outputConfirmation:  false,
	}
	return pb, nil
}

const (
	defaultTitleInput   = "user text input:"
	defaultTitleSingle  = "select single option:"
	defaultTitleMulti   = "select multiple options"
	defaultTitleConfirm = "confirm"
	defaultTitleSearch  = "search from list with filter"
	defaultDescription  = ""
	defaultPlaceholder  = "input"
	defaultPrompt       = "> "
	defaultHelpInput    = "this is Input Help text"
	defaultHelpSingle   = "this is Single Help text"
	defaultHelpMulti    = "this is Multi Help text"
	defaultHelpConfirm  = "this is Confirm Help text"
	defaultHelpSearch   = "↑/↓: move cursor • Enter: submit • Esc: exit"
)

func defaultPromptTitle(form string) string {
	switch form {
	case formInput:
		return defaultTitleInput
	case formSingle:
		return defaultTitleSingle
	case formMulti:
		return defaultTitleMulti
	case formConfirm:
		return defaultTitleConfirm
	case formSearch:
		return defaultTitleSearch
	}
	return "unknown form type"
}

func defaultPrompthelp(form string) string {
	switch form {
	case formInput:
		return defaultHelpInput
	case formSingle:
		return defaultHelpSingle
	case formMulti:
		return defaultHelpMulti
	case formConfirm:
		return defaultHelpConfirm
	case formSearch:
		return defaultHelpSearch
	}
	return "unknown form type"
}
