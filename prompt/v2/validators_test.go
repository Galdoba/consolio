package v2

import (
	"errors"
	"testing"
)

// TestIntegerValidator tests the Integer validation function
func TestIntegerValidator(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected error
	}{
		{"valid integer", "123", nil},
		{"valid negative integer", "-456", nil},
		{"valid zero", "0", nil},
		{"invalid float", "123.45", errors.New("input must be integer")},
		{"invalid string", "abc", errors.New("input must be integer")},
		{"empty string", "", errors.New("input must be integer")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Integer(tt.input)

			if (err == nil) != (tt.expected == nil) {
				t.Errorf("Integer(%q) error = %v, want %v", tt.input, err, tt.expected)
			}

			if err != nil && tt.expected != nil && err.Error() != tt.expected.Error() {
				t.Errorf("Integer(%q) error message = %v, want %v", tt.input, err.Error(), tt.expected.Error())
			}
		})
	}
}

// TestFloat64Validator tests the Float64 validation function
func TestFloat64Validator(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected error
	}{
		{"valid float", "123.45", nil},
		{"valid integer as float", "123", nil},
		{"valid negative float", "-456.78", nil},
		{"valid scientific notation", "1.23e4", nil},
		{"invalid string", "abc", errors.New("input must be float64")},
		{"empty string", "", errors.New("input must be float64")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Float64(tt.input)

			if (err == nil) != (tt.expected == nil) {
				t.Errorf("Float64(%q) error = %v, want %v", tt.input, err, tt.expected)
			}
		})
	}
}

// TestDefaultItemValidation tests the default item validation
func TestDefaultItemValidation(t *testing.T) {
	tests := []struct {
		name     string
		item     *Item
		expected error
	}{
		{"valid item", NewItem("test"), nil},
		{"nil item", nil, errors.New("item is nil")},
		{"item with empty key", &Item{key: ""}, errors.New("item has no key")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := defaultItemValidationFunc(tt.item)

			if (err == nil) != (tt.expected == nil) {
				t.Errorf("defaultItemValidationFunc() error = %v, want %v", err, tt.expected)
			}
		})
	}
}

// TestDefaultItemListValidation tests the default item list validation
func TestDefaultItemListValidation(t *testing.T) {
	tests := []struct {
		name     string
		items    []*Item
		expected error
	}{
		{
			"valid items",
			[]*Item{NewItem("a"), NewItem("b")},
			nil,
		},
		{
			"empty list",
			[]*Item{},
			nil,
		},
		{
			"list with nil item",
			[]*Item{NewItem("a"), nil, NewItem("b")},
			errors.New("item is nil"),
		},
		{
			"list with empty key",
			[]*Item{NewItem("a"), &Item{key: ""}},
			errors.New("item has no key"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := defaultItemListValidationFunc(tt.items)

			if (err == nil) != (tt.expected == nil) {
				t.Errorf("defaultItemListValidationFunc() error = %v, want %v", err, tt.expected)
			}
		})
	}
}
