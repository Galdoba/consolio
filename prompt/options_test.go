package prompt

import (
	"reflect"
	"testing"

	"github.com/charmbracelet/huh"
)

// TestOptionSetters tests the option setter functions
func TestOptionSetters(t *testing.T) {
	tests := []struct {
		name     string
		option   PromptOption
		key      any
		expected any
	}{
		{
			"WithTitle",
			WithTitle("Test Title"),
			KeyTitle,
			"Test Title",
		},
		{
			"WithDescription",
			WithDescription("Test Description"),
			KeyDescription,
			"Test Description",
		},
		{
			"WithPrompt",
			WithPrompt("Test Prompt"),
			KeyPrompt,
			"Test Prompt",
		},
		{
			"WithPlaceholder",
			WithPlaceholder("Test Placeholder"),
			KeyPlaceholder,
			"Test Placeholder",
		},
		{
			"WithWidth",
			WithWidth(100),
			KeyWidth,
			100,
		},
		{
			"WithHeight",
			WithHeight(50),
			KeyHeight,
			50,
		},
		{
			"WithAffirmative",
			WithAffirmative("Yep"),
			KeyAffirmative,
			"Yep",
		},
		{
			"WithNegative",
			WithNegative("Nope"),
			KeyNegative,
			"Nope",
		},
		{
			"WithCaseSensitiveFilter",
			WithCaseSensitiveFilter(true),
			KeyCaseSensitiveFilter,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pb := &promptBuilder{
				promptType:       ptInput,
				settings:         make(map[any]any),
				defaultsRegistry: defaultRegistry(),
			}

			// Apply the option
			tt.option(pb)

			// Verify the value was set
			value, exists := pb.settings[tt.key]
			if !exists {
				t.Errorf("Option %s did not set the value in settings", tt.name)
			}

			if value != tt.expected {
				t.Errorf("Option %s set value = %v, want %v", tt.name, value, tt.expected)
			}
		})
	}
}

// TestWithStringValidator tests the string validator option
func TestWithStringValidator(t *testing.T) {
	validator := func(s string) error { return nil }

	pb := &promptBuilder{
		promptType:       ptInput,
		settings:         make(map[any]any),
		defaultsRegistry: defaultRegistry(),
	}

	WithStringValidator(validator)(pb)

	setValidator, exists := pb.settings[KeyStringValidatorFunc]
	if !exists {
		t.Error("WithStringValidator did not set the validator")
	}

	// Compare function pointers (they should be the same)
	if reflect.ValueOf(setValidator).Pointer() != reflect.ValueOf(validator).Pointer() {
		t.Error("WithStringValidator did not set the correct validator function")
	}
}

// TestWithTheme tests the theme option
func TestWithTheme(t *testing.T) {
	theme := huh.ThemeCatppuccin()

	pb := &promptBuilder{
		promptType:       ptInput,
		settings:         make(map[any]any),
		defaultsRegistry: defaultRegistry(),
	}

	WithTheme(theme)(pb)

	setTheme, exists := pb.settings[KeyTheme]
	if !exists {
		t.Error("WithTheme did not set the theme")
	}

	if setTheme != theme {
		t.Error("WithTheme did not set the correct theme")
	}
}

// TestFromItems tests the items option
func TestFromItems(t *testing.T) {
	items := []*Item{NewItem("a"), NewItem("b")}

	pb := &promptBuilder{
		promptType:       ptSelect,
		settings:         make(map[any]any),
		defaultsRegistry: defaultRegistry(),
	}

	FromItems(items)(pb)

	setItems, exists := pb.settings[KeyItems]
	if !exists {
		t.Error("FromItems did not set the items")
	}

	if len(setItems.([]*Item)) != len(items) {
		t.Error("FromItems did not set the correct number of items")
	}
}

// TestGetterMethods tests the getter methods of promptBuilder
func TestGetterMethods(t *testing.T) {
	pb := &promptBuilder{
		promptType:       ptInput,
		settings:         make(map[any]any),
		defaultsRegistry: defaultRegistry(),
	}

	// Test default values
	if pb.getTitle() == "" {
		t.Error("getTitle() should return default title")
	}

	if pb.getDescription() != "" {
		t.Error("getDescription() should return empty string by default")
	}

	if pb.getWidth() != 0 {
		t.Error("getWidth() should return 0 by default")
	}

	// Test with custom values
	pb.settings[KeyTitle] = "Custom Title"
	pb.settings[KeyWidth] = 200

	if pb.getTitle() != "Custom Title" {
		t.Error("getTitle() should return custom title")
	}

	if pb.getWidth() != 200 {
		t.Error("getWidth() should return custom width")
	}
}

// TestMustGetPanic tests that mustGet panics when key is not found
func TestMustGetPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("mustGet should panic when key is not found")
		}
	}()

	pb := &promptBuilder{
		promptType:       ptInput,
		settings:         make(map[any]any),
		defaultsRegistry: newMapDefaultsRegistry(), // Empty registry
	}

	// This should panic
	pb.getTitle()
}
