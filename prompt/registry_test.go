package prompt

import (
	"testing"
)

// TestMapDefaultsRegistry tests the map-based defaults registry
func TestMapDefaultsRegistry(t *testing.T) {
	registry := newMapDefaultsRegistry()

	// Test SetDefault and GetDefault
	registry.SetDefault("testKey", ptInput, "testValue")

	value, exists := registry.GetDefault("testKey", ptInput)
	if !exists {
		t.Error("GetDefault() should find the set value")
	}
	if value != "testValue" {
		t.Errorf("GetDefault() = %v, want %v", value, "testValue")
	}

	// Test non-existent key
	_, exists = registry.GetDefault("nonExistent", ptInput)
	if exists {
		t.Error("GetDefault() should not find non-existent key")
	}

	// Test same key different prompt type
	registry.SetDefault("testKey", ptSelect, "differentValue")
	value, exists = registry.GetDefault("testKey", ptSelect)
	if !exists || value != "differentValue" {
		t.Errorf("GetDefault() for different prompt type = %v, want %v", value, "differentValue")
	}

	// Original value should still exist
	value, exists = registry.GetDefault("testKey", ptInput)
	if !exists || value != "testValue" {
		t.Errorf("GetDefault() original value = %v, want %v", value, "testValue")
	}
}

// TestRegistryClone tests the cloning functionality
func TestRegistryClone(t *testing.T) {
	original := newMapDefaultsRegistry()
	original.SetDefault("key1", ptInput, "value1")
	original.SetDefault("key2", ptSelect, "value2")

	clone := original.Clone()

	// Test that clone has same values
	val1, exists := clone.GetDefault("key1", ptInput)
	if !exists || val1 != "value1" {
		t.Error("Clone should have same values as original")
	}

	val2, exists := clone.GetDefault("key2", ptSelect)
	if !exists || val2 != "value2" {
		t.Error("Clone should have same values as original")
	}

	// Test that modifications to clone don't affect original
	clone.SetDefault("key1", ptInput, "modifiedValue")

	originalVal, _ := original.GetDefault("key1", ptInput)
	if originalVal == "modifiedValue" {
		t.Error("Modifying clone should not affect original")
	}
}

// TestDefaultRegistry tests the default registry initialization
func TestDefaultRegistry(t *testing.T) {
	registry := defaultRegistry()

	// Test that default values are set for all prompt types
	promptTypes := []promptType{ptInput, ptSelect, ptSelectMulti, ptConfirm, ptSearch}

	for _, pt := range promptTypes {
		// Test common defaults
		_, exists := registry.GetDefault(KeyTitle, pt)
		if !exists {
			t.Errorf("Default title not set for prompt type %v", pt)
		}

		_, exists = registry.GetDefault(KeyDescription, pt)
		if !exists {
			t.Errorf("Default description not set for prompt type %v", pt)
		}

		// Test type-specific defaults
		switch pt {
		case ptInput:
			validator, _ := registry.GetDefault(KeyStringValidatorFunc, pt)
			if validator == nil {
				t.Error("Default string validator not set for input prompt")
			}
		case ptSelect, ptSelectMulti:
			items, _ := registry.GetDefault(KeyItems, pt)
			if items == nil {
				t.Error("Default items not set for select prompt")
			}
		case ptConfirm:
			affirmative, _ := registry.GetDefault(KeyAffirmative, pt)
			if affirmative != "Yes" {
				t.Error("Default affirmative not set correctly for confirm prompt")
			}
		case ptSearch:
			caseSensitive, _ := registry.GetDefault(KeyCaseSensitiveFilter, pt)
			if caseSensitive != false {
				t.Error("Default case sensitivity not set correctly for search prompt")
			}
		}
	}
}

// TestNewDefaultsRegistry tests the public factory function
func TestNewDefaultsRegistry(t *testing.T) {
	registry := NewDefaultsRegistry()

	if registry == nil {
		t.Error("NewDefaultsRegistry() should not return nil")
	}

	// Verify it implements the interface
	if _, ok := registry.(DefaultsRegistry); !ok {
		t.Error("NewDefaultsRegistry() should return DefaultsRegistry implementation")
	}
}
