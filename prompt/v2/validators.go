package v2

import (
	"fmt"
	"strconv"
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
