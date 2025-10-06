package v2

import "maps"

// DefaultsRegistry manages default values for prompt options
type DefaultsRegistry interface {
	GetDefault(key any, pt promptType) (any, bool)
	SetDefault(key any, pt promptType, value any)
	Clone() DefaultsRegistry
}

// MapDefaultsRegistry implements DefaultsRegistry using in-memory maps
type mapDefaultsRegistry struct {
	defaults map[any]map[promptType]any
}

// NewMapDefaultsRegistry creates a new registry instance
func newMapDefaultsRegistry() *mapDefaultsRegistry {
	return &mapDefaultsRegistry{
		defaults: make(map[any]map[promptType]any),
	}
}

// GetDefault retrieves a default value for the given key and prompt type
func (r *mapDefaultsRegistry) GetDefault(key any, pt promptType) (any, bool) {
	typeDefaults, exists := r.defaults[key]
	if !exists {
		return nil, false
	}
	value, exists := typeDefaults[pt]
	return value, exists
}

// RegisterDefault registers a default value for a specific key and prompt type
func (r *mapDefaultsRegistry) SetDefault(key any, pt promptType, value any) {
	if r.defaults[key] == nil {
		r.defaults[key] = make(map[promptType]any)
	}
	r.defaults[key][pt] = value
}

// Clone creates a deep copy of the registry
func (r *mapDefaultsRegistry) Clone() DefaultsRegistry {
	newRegistry := newMapDefaultsRegistry()
	for key, typeMap := range r.defaults {
		newRegistry.defaults[key] = make(map[promptType]any)
		maps.Copy(newRegistry.defaults[key], typeMap)
	}
	return newRegistry
}

func NewDefaultsRegistry() DefaultsRegistry {
	return defaultRegistry()
}

func defaultRegistry() DefaultsRegistry {
	registry := newMapDefaultsRegistry()
	for _, ptType := range []promptType{
		ptInput,
		ptSelect,
	} {
		switch ptType {
		case ptInput:
			registry.SetDefault(KeyTitle, ptType, "user input:")
			registry.SetDefault(KeyPrompt, ptType, "")
			registry.SetDefault(KeyPlaceholder, ptType, "")
			registry.SetDefault(KeyStringValidatorFunc, ptType, defaultStringValidatorFunc)
		case ptSelect:
			registry.SetDefault(KeyTitle, ptType, "select one item:")
		}
		registry.SetDefault(KeyDescription, ptType, "")
		registry.SetDefault(KeyWidth, ptType, 0)
		registry.SetDefault(KeyHeight, ptType, 0)
	}

	return registry
}
