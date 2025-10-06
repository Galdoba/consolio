package v2

// promptType defines the types of supported prompts
type promptType string

const (
	ptInput  promptType = "input"
	ptSelect promptType = "select"
)

// OptionType constrains allowed types for options
type OptionType interface {
	string | int | []*Item | StringValidatorFunc
}

// OptionKey represents a typed option key
type OptionKey[T OptionType] string

const (
	KeyTitle               OptionKey[string]              = "title"
	KeyDescription         OptionKey[string]              = "description"
	KeyPrompt              OptionKey[string]              = "prompt"
	KeyPlaceholder         OptionKey[string]              = "placeholder"
	KeyStringValidatorFunc OptionKey[StringValidatorFunc] = "string_validator_func"
	KeyWidth               OptionKey[int]                 = "width"
	KeyHeight              OptionKey[int]                 = "height"
	KeyItems               OptionKey[[]*Item]             = "items"
)

// PromptOption is a function type for configuring prompts
type PromptOption func(*promptBuilder)

// promptBuilder is the main configuration structure for prompts
type promptBuilder struct {
	promptType      promptType
	settings        map[any]any
	defaltsRegistry DefaultsRegistry
}
