package prompt

import (
	"fmt"
	"maps"

	"github.com/charmbracelet/huh"
)

// DefaultsRegistry manages default values for prompt options across different prompt types.
// It provides methods to get, set, and clone default configurations.
// This interface allows for different storage implementations (in-memory, persistent, etc.).
// (ai generated comment)
type DefaultsRegistry interface {
	// GetDefault retrieves a default value for the given key and prompt type.
	// Returns the value and a boolean indicating if the key was found.
	// (ai generated comment)
	GetDefault(key any, pt promptType) (any, bool)

	// SetDefault registers a default value for a specific key and prompt type.
	// Overwrites existing values for the same key and prompt type combination.
	// (ai generated comment)
	SetDefault(key any, pt promptType, value any)

	// Clone creates a deep copy of the registry.
	// Useful for creating isolated configurations without affecting the original registry.
	// (ai generated comment)
	Clone() DefaultsRegistry
}

// mapDefaultsRegistry implements DefaultsRegistry using in-memory maps.
// This is the default implementation used by the library.
// It stores defaults in a nested map structure: key -> promptType -> value.
// (ai generated comment)
type mapDefaultsRegistry struct {
	defaults map[any]map[promptType]any
}

// newMapDefaultsRegistry creates a new registry instance with empty maps.
// This is the factory function for creating new map-based registries.
// (ai generated comment)
func newMapDefaultsRegistry() *mapDefaultsRegistry {
	return &mapDefaultsRegistry{
		defaults: make(map[any]map[promptType]any),
	}
}

// GetDefault retrieves a default value for the given key and prompt type.
// First looks for the key in the registry, then for the specific prompt type under that key.
// Returns nil and false if either the key or prompt type is not found.
// (ai generated comment)
func (r *mapDefaultsRegistry) GetDefault(key any, pt promptType) (any, bool) {
	typeDefaults, exists := r.defaults[key]
	if !exists {
		return nil, false
	}
	value, exists := typeDefaults[pt]
	return value, exists
}

// RegisterDefault registers a default value for a specific key and prompt type.
// Creates the nested map structure if it doesn't already exist.
// This method is aliased as SetDefault in the interface.
// (ai generated comment)
func (r *mapDefaultsRegistry) SetDefault(key any, pt promptType, value any) {
	if r.defaults[key] == nil {
		r.defaults[key] = make(map[promptType]any)
	}
	r.defaults[key][pt] = value
}

// Clone creates a deep copy of the registry.
// Uses the maps.Copy function to ensure proper copying of nested map structures.
// Returns a new DefaultsRegistry instance with the same default values.
// (ai generated comment)
func (r *mapDefaultsRegistry) Clone() DefaultsRegistry {
	newRegistry := newMapDefaultsRegistry()
	for key, typeMap := range r.defaults {
		newRegistry.defaults[key] = make(map[promptType]any)
		maps.Copy(newRegistry.defaults[key], typeMap)
	}
	return newRegistry
}

// NewDefaultsRegistry creates a new DefaultsRegistry instance.
// This is the public factory function that returns the interface type.
// Currently returns the same instance as defaultRegistry().
// (ai generated comment)
func NewDefaultsRegistry() DefaultsRegistry {
	return defaultRegistry()
}

// defaultRegistry creates and initializes the default registry with predefined values.
// Sets up sensible defaults for all supported prompt types.
// (ai generated comment)
func defaultRegistry() DefaultsRegistry {
	registry := newMapDefaultsRegistry()
	// Initialize defaults for all prompt types
	for _, ptType := range []promptType{
		ptInput,
		ptSelect,
		ptSelectMulti,
		ptConfirm,
		ptSearch,
	} {
		// Set type-specific defaults
		switch ptType {
		case ptInput:
			registry.SetDefault(KeyTitle, ptType, "user input:")
			registry.SetDefault(KeyPrompt, ptType, "")
			registry.SetDefault(KeyPlaceholder, ptType, "")
			registry.SetDefault(KeyStringValidatorFunc, ptType, defaultStringValidatorFunc)
		case ptSelect:
			registry.SetDefault(KeyTitle, ptType, "select one item:")
			registry.SetDefault(KeyItems, ptType, []*Item{})
			registry.SetDefault(KeyItemValidatorFunc, ptType, defaultItemValidatorFunc)
		case ptSelectMulti:
			registry.SetDefault(KeyTitle, ptType, "select item(s):")
			registry.SetDefault(KeyItems, ptType, []*Item{})
			registry.SetDefault(KeyItemListValidatorFunc, ptType, defaultItemListValidatorFunc)
		case ptConfirm:
			registry.SetDefault(KeyTitle, ptType, "confirm:")
			registry.SetDefault(KeyAffirmative, ptType, "Yes")
			registry.SetDefault(KeyNegative, ptType, "No")
		case ptSearch:
			registry.SetDefault(KeyTitle, ptType, "search item:")
			registry.SetDefault(KeyStringValidatorFunc, ptType, defaultStringValidatorFunc)
			registry.SetDefault(KeyItems, ptType, []*Item{})
			registry.SetDefault(KeyItemValidatorFunc, ptType, defaultItemValidatorFunc)
			registry.SetDefault(KeyCaseSensitiveFilter, ptType, false)
		}
		// Set common defaults that apply to all prompt types
		registry.SetDefault(KeyDescription, ptType, "")
		registry.SetDefault(KeyWidth, ptType, 0)
		registry.SetDefault(KeyHeight, ptType, 0)
		registry.SetDefault(KeyTheme, ptType, huh.ThemeBase16())
	}

	return registry
}

func validateRequiredFields(pb *promptBuilder) error {
	switch pb.promptType {
	case ptSelect, ptSelectMulti, ptSearch:
		if _, exists := pb.settings[KeyItems]; !exists {
			if _, exists := pb.defaultsRegistry.GetDefault(KeyItems, pb.promptType); !exists {
				return fmt.Errorf("items are required for %s prompt", pb.promptType)
			}
		}
	}
	return nil
}
