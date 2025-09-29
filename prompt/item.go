package prompt

import (
	"fmt"
)

type Item struct {
	Key      string
	Selected bool
	Data     any
}

func NewItem(key string, data any) *Item {
	return &Item{
		Key:      key,
		Selected: false,
		Data:     data,
	}
}

func (i *Item) GetKey() string {
	return i.Key
}
func (i *Item) IsSelected() bool {
	return i.Selected
}
func (i *Item) PayData() any {
	return i.Data
}

func CreateItem(obj any) *Item {
	return NewItem(fmt.Sprintf("%v", obj), obj)
}
func (i *Item) WithKey(key string) *Item {
	i.Key = key
	return i
}

func NewItemList(objects ...any) []*Item {
	list := []*Item{}
	for _, obj := range objects {
		list = append(list, NewItem(fmt.Sprintf("%v", obj), obj))
	}
	return list
}
