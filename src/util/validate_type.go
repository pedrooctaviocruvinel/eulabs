package util

import (
	"strings"

	"github.com/asaskevich/govalidator"
)

func ValidateType[T any](t T) (valid bool, errors []string) {
	result, err := govalidator.ValidateStruct(t)
	if !result {
		errors = strings.Split(err.Error(), ";")

		return result, errors
	}

	return result, nil
}
