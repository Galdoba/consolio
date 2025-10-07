package v2

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

// setTo safely sets a value in promptBuilder's settings map
func setTo[T OptionType](pb *promptBuilder, key OptionKey[T], value T) {
	if pb.settings == nil {
		pb.settings = make(map[any]any)
	}
	pb.settings[key] = value
}

// getRawValueFrom retrieves a raw value from promptBuilder's settings
func getRawValueFrom[T OptionType](pb *promptBuilder, key OptionKey[T]) (T, bool) {
	if raw, exists := pb.settings[key]; exists {
		if value, ok := raw.(T); ok {
			return value, true
		}
	}
	var zero T
	return zero, false
}

// getFrom retrieves a typed value from promptBuilder with fallback to defaults
func getFrom[T OptionType](pb *promptBuilder, key OptionKey[T]) (T, error) {
	var zero T

	// First try to get user-set value
	if value, exists := getRawValueFrom(pb, key); exists {
		return value, nil
	}

	// Check if defaults exist for this key
	defaults, exists := pb.defaultsRegistry.GetDefault(key, pb.promptType)
	if !exists {
		return zero, fmt.Errorf("option %s not registered", key)
	}

	return defaults.(T), nil
}

// mustGet retrieves a value or panics if not available
func mustGet[T OptionType](pb *promptBuilder, key OptionKey[T]) T {
	value, err := getFrom(pb, key)
	if err != nil {
		panic(fmt.Sprintf("Prompt config error: %v", err))
	}
	return value
}

// WithTitle sets the title for a prompt
func WithTitle(title string) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyTitle, title)
	}
}

func (pb *promptBuilder) getTitle() string {
	return mustGet[string](pb, KeyTitle)
}

// WithTitle sets the title for a prompt
func WithDescription(descr string) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyDescription, descr)
	}
}

func (pb *promptBuilder) getDescription() string {
	return mustGet[string](pb, KeyDescription)
}

// WithTitle sets the title for a prompt
func WithPrompt(prompt string) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyPrompt, prompt)
	}
}

func (pb *promptBuilder) getPrompt() string {
	return mustGet[string](pb, KeyPrompt)
}

// WithTitle sets the title for a prompt
func WithPlaceholder(placeholder string) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyPlaceholder, placeholder)
	}
}

func (pb *promptBuilder) getPlaceholder() string {
	return mustGet[string](pb, KeyPlaceholder)
}

// WithTitle sets the title for a prompt
func WithStringValidator(validator func(string) error) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyStringValidatorFunc, validator)
	}
}

func (pb *promptBuilder) getStringValidator() StringValidatorFunc {
	return mustGet[StringValidatorFunc](pb, KeyStringValidatorFunc)
}

func defaultStringValidator(string) error { return nil }

// WithWidth sets the width for a prompt
func WithWidth(w int) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyWidth, w)
	}
}
func (pb *promptBuilder) getWidth() int {
	return mustGet[int](pb, KeyWidth)
}

// WithHeight sets the height for a prompt
func WithHeight(h int) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyHeight, h)
	}
}

func (pb *promptBuilder) getHeight() int {
	return mustGet[int](pb, KeyHeight)
}

func WithTheme(theme *huh.Theme) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyTheme, theme)
	}
}

func (pb *promptBuilder) getTheme() *huh.Theme {
	return mustGet[*huh.Theme](pb, KeyTheme)
}

// WithTitle sets the title for a prompt
func WithItemValidator(validator func(*Item) error) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyItemValidatorFunc, validator)
	}
}

func (pb *promptBuilder) getItemValidator() ItemValidationFunc {
	return mustGet[ItemValidationFunc](pb, KeyItemValidatorFunc)
}

// WithTitle sets the title for a prompt
func WithItemListValidator(validator func([]*Item) error) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyItemListValidatorFunc, validator)
	}
}

func (pb *promptBuilder) getItemListValidator() ItemListValidationFunc {
	return mustGet(pb, KeyItemListValidatorFunc)
}

// FromItems sets the items for selection-based prompts
func FromItems(items []*Item) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyItems, items)
	}
}

func (pb *promptBuilder) getItems() []*Item {
	return mustGet[[]*Item](pb, KeyItems)
}

func WithAffirmative(affirmative string) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyAffirmative, affirmative)
	}
}
func (pb *promptBuilder) getAffirmative() string {
	return mustGet(pb, KeyAffirmative)
}

func WithNegative(negative string) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyNegative, negative)
	}
}
func (pb *promptBuilder) getNegative() string {
	return mustGet(pb, KeyNegative)
}
