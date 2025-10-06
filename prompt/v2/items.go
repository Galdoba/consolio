package v2

// Item represents a selectable item in list-based prompts
type Item struct {
	key     string
	payload any
}

// NewItem creates a new Item with optional payload
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

// Key returns the display key of the item
func (i Item) Key() string {
	return i.key
}

// Payload returns the associated payload of the item
func (i Item) Payload() any {
	return i.payload
}