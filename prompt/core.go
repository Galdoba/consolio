package prompt

import "github.com/charmbracelet/huh"

// promptType defines the types of supported prompts in the library.
// It's used to distinguish between different prompt categories like input, selection, confirmation etc.
// (ai generated comment)
type promptType string

const (
	ptInput       promptType = "input"        // Single line text input prompt
	ptSelect      promptType = "select"       // Single item selection from a list
	ptSelectMulti promptType = "select_multi" // Multiple items selection from a list
	ptConfirm     promptType = "confirm"      // Yes/No confirmation dialog
	ptSearch      promptType = "search"       // Interactive search through filtered items
)

// OptionType constrains allowed types for prompt configuration options.
// This generic constraint ensures type safety when working with prompt options.
// (ai generated comment)
type OptionType interface {
	string | int | bool | []*Item | StringValidatorFunc | ItemValidationFunc | ItemListValidationFunc | *huh.Theme
}

// OptionKey represents a typed option key for prompt configuration.
// It uses generics to ensure type safety when accessing configuration values.
// (ai generated comment)
type OptionKey[T OptionType] string

const (
	KeyTitle                 OptionKey[string]                 = "title"                    // Main title displayed for the prompt
	KeyDescription           OptionKey[string]                 = "description"              // Additional description text
	KeyPrompt                OptionKey[string]                 = "prompt"                   // Input prompt text
	KeyPlaceholder           OptionKey[string]                 = "placeholder"              // Placeholder text for input fields
	KeyStringValidatorFunc   OptionKey[StringValidatorFunc]    = "string_validator_func"    // Function to validate string input
	KeyItems                 OptionKey[[]*Item]                = "items"                    // List of selectable items
	KeyItemValidatorFunc     OptionKey[ItemValidationFunc]     = "items_validator_func"     // Function to validate individual items
	KeyItemListValidatorFunc OptionKey[ItemListValidationFunc] = "item_list_validator_func" // Function to validate item lists
	KeyAffirmative           OptionKey[string]                 = "affirmative"              // "Yes" button text for confirmation
	KeyNegative              OptionKey[string]                 = "negative"                 // "No" button text for confirmation
	KeyWidth                 OptionKey[int]                    = "width"                    // Prompt display width
	KeyHeight                OptionKey[int]                    = "height"                   // Prompt display height
	KeyTheme                 OptionKey[*huh.Theme]             = "theme"                    // Visual theme for the prompt
	KeyCaseSensitiveFilter   OptionKey[bool]                   = "case_sensitive_filter"    // Case sensitivity for search filters
)

// PromptOption is a function type for configuring prompts through the builder pattern.
// Each option function modifies the promptBuilder configuration.
// (ai generated comment)
type PromptOption func(*promptBuilder)

// promptBuilder is the main configuration structure for building prompts.
// It maintains prompt type, custom settings, and default values registry.
// (ai generated comment)
type promptBuilder struct {
	promptType       promptType       // Type of prompt being built
	settings         map[any]any      // Custom settings overriding defaults
	defaultsRegistry DefaultsRegistry // Registry of default values for prompt types
}
