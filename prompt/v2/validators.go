package v2

import (
	"fmt"
	"strconv"
)

// StringValidatorFunc is a function type that validates string input.
// It takes a string as input and returns an error if validation fails.
// Used for validating user input in text-based prompts.
// (ai generated comment)
type StringValidatorFunc func(string) error

// defaultStringValidatorFunc is the default string validator function.
// It uses the defaultStringValidator which accepts any input.
// (ai generated comment)
var defaultStringValidatorFunc StringValidatorFunc = defaultStringValidator

// Integer validates that a string can be parsed as an integer.
// Uses strconv.Atoi internally to check if the string is a valid integer.
// Returns an error if the string cannot be parsed as an integer.
// (ai generated comment)
func Integer(s string) error {
	if _, err := strconv.Atoi(s); err != nil {
		return fmt.Errorf("input must be integer")
	}
	return nil
}

// Float64 validates that a string can be parsed as a float64.
// Uses strconv.ParseFloat internally to check if the string is a valid float.
// Returns an error if the string cannot be parsed as a float64.
// (ai generated comment)
func Float64(s string) error {
	if _, err := strconv.ParseFloat(s, 64); err != nil {
		return fmt.Errorf("input must be float64")
	}
	return nil
}

// ItemValidationFunc is a function type that validates individual Item objects.
// Used to validate items in selection-based prompts.
// Returns an error if the item fails validation.
// (ai generated comment)
type ItemValidationFunc func(*Item) error

// defaultItemValidatorFunc is the default item validator function.
// It performs basic validation to ensure items are not nil and have keys.
// (ai generated comment)
var defaultItemValidatorFunc ItemValidationFunc = defaultItemValidationFunc

// defaultItemValidationFunc provides basic validation for Item objects.
// Checks that the item is not nil and has a non-empty key.
// This is the default validation used when no custom validator is provided.
// (ai generated comment)
func defaultItemValidationFunc(i *Item) error {
	if i == nil {
		return fmt.Errorf("item is nil")
	}
	if i.key == "" {
		return fmt.Errorf("item has no key")
	}
	return nil
}

// ItemListValidationFunc is a function type that validates lists of Item objects.
// Used to validate entire item collections in multi-select prompts.
// Returns an error if any item in the list fails validation.
// (ai generated comment)
type ItemListValidationFunc func([]*Item) error

// defaultItemListValidatorFunc is the default item list validator function.
// It validates each item in the list using the same rules as defaultItemValidationFunc.
// (ai generated comment)
var defaultItemListValidatorFunc ItemListValidationFunc = defaultItemListValidationFunc

// defaultItemListValidationFunc validates a list of Item objects.
// Iterates through all items and applies the same validation as defaultItemValidationFunc.
// Returns the first validation error encountered, if any.
// (ai generated comment)
func defaultItemListValidationFunc(items []*Item) error {
	for _, i := range items {
		if i == nil {
			return fmt.Errorf("item is nil")
		}
		if i.key == "" {
			return fmt.Errorf("item has no key")
		}
	}
	return nil
}
