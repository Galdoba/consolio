package v2

import (
	"testing"
)

// TestIntegrationBasic tests basic integration of the library components
func TestIntegrationBasic(t *testing.T) {
	// Create items
	items := []*Item{
		NewItem("Apple"),
		NewItem("Banana"),
		NewItem("Cherry"),
	}

	// Test that items can be used with options
	option := FromItems(items)

	pb := &promptBuilder{
		promptType:       ptSelect,
		settings:         make(map[any]any),
		defaultsRegistry: defaultRegistry(),
	}

	option(pb)

	retrievedItems := pb.getItems()
	if len(retrievedItems) != len(items) {
		t.Errorf("Integration: got %d items, want %d", len(retrievedItems), len(items))
	}
}

// TestIntegrationWithValidators tests integration with validators
func TestIntegrationWithValidators(t *testing.T) {
	items := []*Item{
		NewItem("Valid1"),
		NewItem("Valid2"),
	}

	// Test item validator integration
	pb := &promptBuilder{
		promptType:       ptSelect,
		settings:         make(map[any]any),
		defaultsRegistry: defaultRegistry(),
	}

	// WithItemValidator(NoNumbers)(pb)
	FromItems(items)(pb)

	validator := pb.getItemValidator()

	// Test the validator works
	err := validator(items[0])
	if err != nil {
		t.Errorf("Integration validator error: %v", err)
	}
}

// TestIntegrationRegistryWithBuilder tests registry integration with prompt builder
func TestIntegrationRegistryWithBuilder(t *testing.T) {
	registry := defaultRegistry()

	pb := &promptBuilder{
		promptType:       ptInput,
		settings:         make(map[any]any),
		defaultsRegistry: registry,
	}

	// Should get default value from registry
	title := pb.getTitle()
	if title == "" {
		t.Error("Integration: should get default title from registry")
	}

	// Custom value should override default
	WithTitle("Custom Title")(pb)
	customTitle := pb.getTitle()
	if customTitle != "Custom Title" {
		t.Error("Integration: custom title should override default")
	}
}
