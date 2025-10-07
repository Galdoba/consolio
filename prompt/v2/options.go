package v2

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

// setTo stores a configuration value in the prompt builder's settings map.
// This is a generic helper function used by all option setters to maintain type safety.
// (ai generated comment)
func setTo[T OptionType](pb *promptBuilder, key OptionKey[T], value T) {
	if pb.settings == nil {
		pb.settings = make(map[any]any)
	}
	pb.settings[key] = value
}

// getRawValueFrom retrieves a raw value from the prompt builder's settings map.
// Returns the value and a boolean indicating if the key was found.
// This function performs type assertion to ensure type safety.
// (ai generated comment)
func getRawValueFrom[T OptionType](pb *promptBuilder, key OptionKey[T]) (T, bool) {
	if raw, exists := pb.settings[key]; exists {
		if value, ok := raw.(T); ok {
			return value, true
		}
	}
	var zero T
	return zero, false
}

// getFrom retrieves a configuration value, first checking user settings then falling back to defaults.
// Returns an error if the key is not found in either location.
// This is the main lookup function used by all getter methods.
// (ai generated comment)
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

// mustGet retrieves a configuration value and panics if the key is not found.
// Used internally when a configuration value is required for prompt operation.
// (ai generated comment)
func mustGet[T OptionType](pb *promptBuilder, key OptionKey[T]) T {
	value, err := getFrom(pb, key)
	if err != nil {
		panic(fmt.Sprintf("Prompt config error: %v", err))
	}
	return value
}

// WithTitle sets the main title text for the prompt.
// The title is displayed prominently at the top of the prompt.
// (ai generated comment)
func WithTitle(title string) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyTitle, title)
	}
}

func (pb *promptBuilder) getTitle() string {
	return mustGet(pb, KeyTitle)
}

// WithDescription sets the description text for the prompt.
// Description provides additional context or instructions below the title.
// (ai generated comment)
func WithDescription(descr string) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyDescription, descr)
	}
}

func (pb *promptBuilder) getDescription() string {
	return mustGet(pb, KeyDescription)
}

// WithPrompt sets the prompt text for input fields.
// This is the text that appears immediately before the input area.
// (ai generated comment)
func WithPrompt(prompt string) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyPrompt, prompt)
	}
}

func (pb *promptBuilder) getPrompt() string {
	return mustGet(pb, KeyPrompt)
}

// WithPlaceholder sets the placeholder text for input fields.
// Placeholder text appears in the input field when it's empty.
// (ai generated comment)
func WithPlaceholder(placeholder string) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyPlaceholder, placeholder)
	}
}

func (pb *promptBuilder) getPlaceholder() string {
	return mustGet(pb, KeyPlaceholder)
}

// WithStringValidator sets a validation function for string input.
// The validator function should return an error if the input is invalid.
// (ai generated comment)
func WithStringValidator(validator func(string) error) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyStringValidatorFunc, validator)
	}
}

func (pb *promptBuilder) getStringValidator() StringValidatorFunc {
	return mustGet(pb, KeyStringValidatorFunc)
}

// defaultStringValidator is the default validator that accepts any string input.
// Always returns nil, indicating all input is valid.
// (ai generated comment)
func defaultStringValidator(string) error { return nil }

// WithWidth sets the display width for the prompt.
// Width of 0 uses the terminal's default width.
// (ai generated comment)
func WithWidth(w int) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyWidth, w)
	}
}

func (pb *promptBuilder) getWidth() int {
	return mustGet(pb, KeyWidth)
}

// WithHeight sets the display height for the prompt.
// Height of 0 uses the terminal's default height.
// (ai generated comment)
func WithHeight(h int) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyHeight, h)
	}
}

func (pb *promptBuilder) getHeight() int {
	return mustGet(pb, KeyHeight)
}

// WithTheme sets the visual theme for the prompt.
// The theme controls colors, styles, and overall appearance of the prompt.
// (ai generated comment)
func WithTheme(theme *huh.Theme) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyTheme, theme)
	}
}

func (pb *promptBuilder) getTheme() *huh.Theme {
	return mustGet(pb, KeyTheme)
}

// WithItemValidator sets a validation function for individual items in selection prompts.
// The validator is called for each item in the selection list.
// (ai generated comment)
func WithItemValidator(validator func(*Item) error) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyItemValidatorFunc, validator)
	}
}

func (pb *promptBuilder) getItemValidator() ItemValidationFunc {
	return mustGet(pb, KeyItemValidatorFunc)
}

// WithItemListValidator sets a validation function for entire item lists.
// Used in multi-select prompts to validate the complete selection set.
// (ai generated comment)
func WithItemListValidator(validator func([]*Item) error) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyItemListValidatorFunc, validator)
	}
}

func (pb *promptBuilder) getItemListValidator() ItemListValidationFunc {
	return mustGet(pb, KeyItemListValidatorFunc)
}

// FromItems sets the list of selectable items for selection prompts.
// (ai generated comment)
func FromItems(items []*Item) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyItems, items)
	}
}

func (pb *promptBuilder) getItems() []*Item {
	return mustGet(pb, KeyItems)
}

// WithAffirmative sets the text for the affirmative (Yes) button in confirmation prompts.
// Default is "Yes" if not specified.
// (ai generated comment)
func WithAffirmative(affirmative string) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyAffirmative, affirmative)
	}
}

func (pb *promptBuilder) getAffirmative() string {
	return mustGet(pb, KeyAffirmative)
}

// WithNegative sets the text for the negative (No) button in confirmation prompts.
// Default is "No" if not specified.
// (ai generated comment)
func WithNegative(negative string) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyNegative, negative)
	}
}

func (pb *promptBuilder) getNegative() string {
	return mustGet(pb, KeyNegative)
}

// WithCaseSensitiveFilter sets whether search filtering should be case sensitive.
// When false (default), search is case insensitive.
// (ai generated comment)
func WithCaseSensitiveFilter(csf bool) PromptOption {
	return func(pb *promptBuilder) {
		setTo(pb, KeyCaseSensitiveFilter, csf)
	}
}

func (pb *promptBuilder) getCaseSensitive() bool {
	return mustGet(pb, KeyCaseSensitiveFilter)
}
