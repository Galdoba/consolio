package v2

import (
	"testing"
)

// TestNewItem tests the creation of items with various payload configurations
func TestNewItem(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		payload  []any
		expected *Item
	}{
		{
			name:     "item with key only",
			key:      "test",
			payload:  []any{},
			expected: &Item{key: "test", payload: "test"},
		},
		{
			name:     "item with string payload",
			key:      "test",
			payload:  []any{"payload"},
			expected: &Item{key: "test", payload: "payload"},
		},
		{
			name:     "item with int payload",
			key:      "test",
			payload:  []any{42},
			expected: &Item{key: "test", payload: 42},
		},
		{
			name:     "item with struct payload",
			key:      "test",
			payload:  []any{struct{ Name string }{Name: "test"}},
			expected: &Item{key: "test", payload: struct{ Name string }{Name: "test"}},
		},
		{
			name:     "item with multiple payloads uses last",
			key:      "test",
			payload:  []any{"first", "second", "third"},
			expected: &Item{key: "test", payload: "third"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := NewItem(tt.key, tt.payload...)

			if item.Key() != tt.expected.key {
				t.Errorf("Key() = %v, want %v", item.Key(), tt.expected.key)
			}

			if item.Payload() != tt.expected.payload {
				t.Errorf("Payload() = %v, want %v", item.Payload(), tt.expected.payload)
			}
		})
	}
}

// TestItemMethods tests the getter methods of Item
func TestItemMethods(t *testing.T) {
	item := NewItem("test_key", "test_payload")

	if item.Key() != "test_key" {
		t.Errorf("Key() = %v, want %v", item.Key(), "test_key")
	}

	if item.Payload() != "test_payload" {
		t.Errorf("Payload() = %v, want %v", item.Payload(), "test_payload")
	}
}
