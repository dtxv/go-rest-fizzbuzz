package domain

import (
	"fmt"
	"strings"
)

// Interface
type IValidator interface {
	Validate(model FizzBuzzModel) error
}

// Struct
type validator struct{}

// Constructor
func NewValidator() IValidator {
	return validator{}
}

// Validate FizzBuzzModel data
func (c validator) Validate(model FizzBuzzModel) error {
	validationErrors := validateModel(model)

	if len(validationErrors) > 0 {
		return fmt.Errorf("invalid params : %s", strings.Join(validationErrors, " ; "))
	}
	return nil
}

// ----- Utils ------

func validateModel(model FizzBuzzModel) []string {
	validationErrors := make([]string, 0)
	validationErrors = appendError(validationErrors, validateInt1(model.Int1))
	validationErrors = appendError(validationErrors, validateInt2(model.Int2))
	validationErrors = appendError(validationErrors, validateLimit(model.Limit))
	validationErrors = appendError(validationErrors, validateStr1(model.Str1))
	validationErrors = appendError(validationErrors, validateStr2(model.Str2))
	return validationErrors
}

func appendError(validationErrors []string, err error) []string {
	if err != nil {
		validationErrors = append(validationErrors, err.Error())
	}
	return validationErrors
}

func validateInt1(int1 *int) error {
	if int1 == nil {
		return fmt.Errorf("int1 is mandatory")
	}
	return nil
}

func validateInt2(int2 *int) error {
	if int2 == nil {
		return fmt.Errorf("int2 is mandatory")
	}
	return nil
}

func validateLimit(limit *int) error {
	if limit == nil {
		return fmt.Errorf("limit is mandatory")
	}
	return nil
}

func validateStr1(str1 *string) error {
	if str1 == nil {
		return fmt.Errorf("str1 is mandatory")
	}
	return nil
}

func validateStr2(str2 *string) error {
	if str2 == nil {
		return fmt.Errorf("str2 is mandatory")
	}
	return nil
}
