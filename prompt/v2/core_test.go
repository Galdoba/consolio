// [file name]: core_test.go
package v2

import (
	"fmt"
	"testing"
)

// TestPromptTypeValues tests that prompt type constants have expected values
func TestPromptTypeValues(t *testing.T) {
	tests := []struct {
		pt       promptType
		expected string
	}{
		{ptInput, "input"},
		{ptSelect, "select"},
		{ptSelectMulti, "select_multi"},
		{ptConfirm, "confirm"},
		{ptSearch, "search"},
	}

	for _, tt := range tests {
		t.Run(string(tt.pt), func(t *testing.T) {
			if string(tt.pt) != tt.expected {
				t.Errorf("promptType %v = %v, want %v", tt.pt, string(tt.pt), tt.expected)
			}
		})
	}
}

// TestOptionKeyValues tests that option key constants have expected values
func TestOptionKeyValues(t *testing.T) {
	tests := []struct {
		key      any
		expected string
	}{
		{KeyTitle, "title"},
		{KeyDescription, "description"},
		{KeyPrompt, "prompt"},
		{KeyPlaceholder, "placeholder"},
		{KeyStringValidatorFunc, "string_validator_func"},
		{KeyItems, "items"},
		{KeyItemValidatorFunc, "items_validator_func"},
		{KeyItemListValidatorFunc, "item_list_validator_func"},
		{KeyAffirmative, "affirmative"},
		{KeyNegative, "negative"},
		{KeyWidth, "width"},
		{KeyHeight, "height"},
		{KeyTheme, "theme"},
		{KeyCaseSensitiveFilter, "case_sensitive_filter"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if fmt.Sprintf("%v", tt.key) != tt.expected {
				t.Errorf("OptionKey = %v, want %v", tt.key, tt.expected)
			}
		})
	}
}

// TestPromptBuilderInitialization tests prompt builder creation
func TestPromptBuilderInitialization(t *testing.T) {
	pb := &promptBuilder{
		promptType:       ptInput,
		settings:         make(map[any]any),
		defaultsRegistry: defaultRegistry(),
	}

	if pb.promptType != ptInput {
		t.Errorf("promptType = %v, want %v", pb.promptType, ptInput)
	}

	if pb.settings == nil {
		t.Error("settings map should be initialized")
	}

	if pb.defaultsRegistry == nil {
		t.Error("defaultsRegistry should be initialized")
	}
}

// TestPromptOptionType tests that PromptOption is a function type
func TestPromptOptionType(t *testing.T) {
	// This test verifies that PromptOption is correctly defined as a function type
	option := WithTitle("test")

	if option == nil {
		t.Error("PromptOption function should not be nil")
	}

	// Verify it can be applied to promptBuilder
	pb := &promptBuilder{
		promptType:       ptInput,
		settings:         make(map[any]any),
		defaultsRegistry: defaultRegistry(),
	}

	option(pb)

	if pb.settings[KeyTitle] != "test" {
		t.Error("PromptOption should modify promptBuilder")
	}
}
