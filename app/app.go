package app

import (
	"errors"
	"fmt"
)

const (
	FIZZ      string = "Fizz"
	BUZZ             = "Buzz"
	FIZZ_BUZZ        = "FizzBuzz"
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
		return FIZZ_BUZZ, nil
	case isMultipleOfThree(input):
		return FIZZ, nil
	case isMultipleOfFive(input):
		return BUZZ, nil
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
