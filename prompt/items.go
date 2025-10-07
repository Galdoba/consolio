package prompt

// Item represents a selectable item in list-based prompts like select, multi-select, and search.
// It contains a display key and an optional payload for storing additional data.
// The payload can be any type, making Items flexible for various use cases.
// (ai generated comment)
type Item struct {
	key     string // Display text shown to the user in prompts
	payload any    // Optional associated data (can be any type)
}

// NewItem creates a new Item with the specified key and optional payload.
// If multiple payloads are provided, only the last one will be used.
// If no payload is provided, the key will be used as the payload.
// This factory function ensures proper Item initialization.
// (ai generated comment)
func NewItem(key string, payload ...any) *Item {
	i := Item{}
	i.key = key
	for _, pd := range payload {
		i.payload = pd
	}
	if i.payload == nil {
		i.payload = i.key
	}
	return &i
}

// Key returns the display key of the item.
// This is the text that will be visible to users in selection prompts.
// (ai generated comment)
func (i Item) Key() string {
	return i.key
}

// Payload returns the associated payload of the item.
// The payload can contain any additional data needed by the application.
// Returns nil if no payload was set during Item creation.
// (ai generated comment)
func (i Item) Payload() any {
	return i.payload
}
