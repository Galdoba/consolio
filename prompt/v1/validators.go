package v1

import (
	"errors"
	"strconv"
)

var defaultTextValiidatorFunc = func(string) error {
	return nil
}

var defaultItemValiidatorFunc = func(*Item) error {
	return nil
}

var defaultItemSetValiidatorFunc = func([]*Item) error {
	return nil
}

var Number = func(str string) error {
	_, err := strconv.Atoi(str)
	if err != nil {
		return errors.New("input must be integer")
	}
	return nil
}
