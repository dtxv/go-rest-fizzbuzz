package domain

import (
	"strconv"
)

// Interface
type IFizzBuzzService interface {
	FillFizzBuzz(request FizzBuzzModel) []string
}

// Struct
type fizzBuzzService struct{}

// Constructor
func NewFizzBuzzService() IFizzBuzzService {
	return &fizzBuzzService{}
}

// Fill fizz buzz result lists
func (f *fizzBuzzService) FillFizzBuzz(model FizzBuzzModel) []string {
	results := make([]string, 0)
	for i := 1; i < *model.Limit; i++ {
		results = append(results, computeFizzBuzz(i, model))
	}

	return results
}

// Compute fizz buzz result for one number
func computeFizzBuzz(i int, model FizzBuzzModel) string {
	if i%*model.Int1 == 0 && i%*model.Int2 == 0 {
		return *model.Str1 + *model.Str2
	}
	if i%*model.Int1 == 0 {
		return *model.Str1
	}
	if i%*model.Int2 == 0 {
		return *model.Str2
	}
	return strconv.Itoa(i)
}
