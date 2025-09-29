package prompt

import "github.com/charmbracelet/huh"

type PromptOption func(*promptBuilder)

func WithInitialPrompt(initial string) PromptOption {
	return func(pb *promptBuilder) {
		pb.prompt = initial
	}
}
func WithPlaceholder(initial string) PromptOption {
	return func(pb *promptBuilder) {
		pb.placeholder = initial
	}
}

func WithTextValidator(validationFunc func(string) error) PromptOption {
	return func(pb *promptBuilder) {
		pb.textValidator = validationFunc
	}
}

func WithItemValidator(validationFunc func(*Item) error) PromptOption {
	return func(pb *promptBuilder) {
		pb.itemValidator = validationFunc
	}
}

func WithDescription(descr string) PromptOption {
	return func(pb *promptBuilder) {
		pb.description = descr
	}
}

func WithTitle(title string) PromptOption {
	return func(pb *promptBuilder) {
		pb.title = title
	}
}

func FromItems(items ...*Item) PromptOption {
	return func(pb *promptBuilder) {
		pb.items = items
	}
}

func WithTheme(theme *huh.Theme) PromptOption {
	return func(pb *promptBuilder) {
		pb.theme = theme
	}
}

func WithInlineSelection(isel bool) PromptOption {
	return func(pb *promptBuilder) {
		pb.selectInline = isel
	}
}

func WithSize(width, height int) PromptOption {
	return func(pb *promptBuilder) {
		pb.width = width
		pb.height = height

	}
}
