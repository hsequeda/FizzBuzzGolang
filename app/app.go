package app

import (
	"errors"
	"fmt"
)

const (
	// FizzCase represent the return of FizzBuzz function when the input is multiple of three.
	FizzCase string = "Fizz"
	// BuzzCase represent the return of FizzBuzz function when the input is multiple of five.
	BuzzCase = "Buzz"
	// FizzBuzzCase represent the return of FizzBuzz function when the input is multiple of three and five.
	FizzBuzzCase = "FizzBuzz"
)

// FizzBuzz return the input number as a string.
// If input is multiple of 3 Fizz.
// If input is multiple of 5 return Buzz.
// If input is multiple of 3 and 5 return FizzBuzz.
// Only works with inputs in range [1-100].
func FizzBuzz(input uint8) (string, error) {
	if input == 0 || input > 100 {
		return "", errors.New("FizzBuzz range is [1-100]")
	}
	switch {
	case isMultipleOfThree(input) && isMultipleOfFive(input):
		return FizzBuzzCase, nil
	case isMultipleOfThree(input):
		return FizzCase, nil
	case isMultipleOfFive(input):
		return BuzzCase, nil
	default:
		return fmt.Sprint(input), nil
	}
}

func isMultipleOfThree(input uint8) bool {
	return input%3 == 0
}

func isMultipleOfFive(input uint8) bool {
	return input%5 == 0
}
