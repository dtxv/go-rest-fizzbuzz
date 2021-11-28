package domain

import "testing"

var (
	int1  = 3
	int2  = 5
	limit = 100
	str1  = "fizz"
	str2  = "buzz"

	model = FizzBuzzModel{
		Int1:  &int1,
		Int2:  &int2,
		Limit: &limit,
		Str1:  &str1,
		Str2:  &str2,
	}
)

func TestReturnNumber(t *testing.T) {
	expected := "2"
	if result := computeFizzBuzz(2, model); result != expected {
		t.Fatalf("computeFizzBuzz(2)) = %v, want %v", result, expected)
	}
}

func TestReturnFizz(t *testing.T) {
	expected := "fizz"
	if result := computeFizzBuzz(9, model); result != expected {
		t.Fatalf("computeFizzBuzz(2)) = %v, want %v", result, expected)
	}
}

func TestReturnBuzz(t *testing.T) {
	expected := "buzz"
	if result := computeFizzBuzz(20, model); result != expected {
		t.Fatalf("computeFizzBuzz(2)) = %v, want %v", result, expected)
	}
}

func TestReturnFizzBuzz(t *testing.T) {
	expected := "fizzbuzz"
	if result := computeFizzBuzz(45, model); result != expected {
		t.Fatalf("computeFizzBuzz(2)) = %v, want %v", result, expected)
	}
}
