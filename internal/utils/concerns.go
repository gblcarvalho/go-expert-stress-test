package utils

import (
	"errors"
	"strings"
)

func checkAssert(value bool, errMsg string) error {
	if value == true {
		return nil
	}
	return errors.New(errMsg)
}

func AssertNotEmpty(value string, errMsg string) error {
	check := strings.TrimSpace(value) != ""
	return checkAssert(check, errMsg)
}

func AssertPositive(value int, errMsg string) error {
	return checkAssert(value > 0, errMsg)
}
