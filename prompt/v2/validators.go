package v2

import (
	"fmt"
	"strconv"
	"strings"
)

type StringValidatorFunc func(string) error

var defaultStringValidatorFunc StringValidatorFunc = defaultStringValidator

func Integer(s string) error {
	if _, err := strconv.Atoi(s); err != nil {
		return fmt.Errorf("input must be integer")
	}
	return nil
}

func Float64(s string) error {
	if _, err := strconv.ParseFloat(s, 64); err != nil {
		return fmt.Errorf("input must be float64")
	}
	return nil
}

type ItemValidationFunc func(*Item) error

var defaultItemValidatorFunc ItemValidationFunc = defaultItemValidationFunc

func defaultItemValidationFunc(i *Item) error {
	fmt.Println("validate", i)
	if i == nil {
		return fmt.Errorf("item is nil")
	}
	if i.key == "" {
		return fmt.Errorf("item has no key")
	}
	return nil
}

func NoNumbers(i *Item) error {
	for l := range strings.SplitSeq(i.key, "") {
		switch l {
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			return fmt.Errorf("no numbers allowed")
		default:
		}
	}
	return nil
}
