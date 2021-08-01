package app

import "fmt"

const (
	FIZZ      string = "Fizz"
	BUZZ             = "Buzz"
	FIZZ_BUZZ        = "FizzBuzz"
)

// FizzBuzz return the input number as a string.
// If input is multiple of 3 Fizz.
func FizzBuzz(input uint8) string {
	switch {
	case isMultipleOfThree(input):
		return FIZZ
	default:
		return fmt.Sprint(input)
	}
}

func isMultipleOfThree(input uint8) bool {
	return input%3 == 0
}
