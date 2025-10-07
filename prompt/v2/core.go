package v2

import "github.com/charmbracelet/huh"

// promptType defines the types of supported prompts
type promptType string

const (
	ptInput       promptType = "input"
	ptSelect      promptType = "select"
	ptSelectMulti promptType = "select_multi"
)

// OptionType constrains allowed types for options
type OptionType interface {
	string | int | []*Item | StringValidatorFunc | *huh.Theme
}

// OptionKey represents a typed option key
type OptionKey[T OptionType] string

const (
	KeyTitle               OptionKey[string]              = "title"
	KeyDescription         OptionKey[string]              = "description"
	KeyPrompt              OptionKey[string]              = "prompt"
	KeyPlaceholder         OptionKey[string]              = "placeholder"
	KeyStringValidatorFunc OptionKey[StringValidatorFunc] = "string_validator_func"
	KeyItems               OptionKey[[]*Item]             = "items"
	KeyWidth               OptionKey[int]                 = "width"
	KeyHeight              OptionKey[int]                 = "height"
	KeyTheme               OptionKey[*huh.Theme]          = "theme"
)

// PromptOption is a function type for configuring prompts
type PromptOption func(*promptBuilder)

// promptBuilder is the main configuration structure for prompts
type promptBuilder struct {
	promptType       promptType
	settings         map[any]any
	defaultsRegistry DefaultsRegistry
}
