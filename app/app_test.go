package app_test

import (
	"fizzbuzz/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("All three multiples (only) (range [1-100])", func(t *testing.T) {
		threeMultiples := []uint8{
			3, 6, 9, 12, 18, 21, 24, 27, 33, 36, 39, 42, 48, 51, 54, 57,
			63, 66, 69, 72, 78, 81, 84, 87, 93, 96, 99,
		}
		for _, threeMultiple := range threeMultiples {
			out := app.FizzBuzz(threeMultiple)
			assert.Equal(t, app.FIZZ, out)
		}
	})

	t.Run("All five multiples (only) (range [1-100])", func(t *testing.T) {
		fiveMultiples := []uint8{
			5, 10, 20, 25, 35, 40, 50, 55, 65, 70, 80, 85, 95, 100,
		}
		for _, fiveMultiple := range fiveMultiples {
			out := app.FizzBuzz(fiveMultiple)
			assert.Equal(t, app.BUZZ, out)
		}
	})

	t.Run("All three && five multiples (only) (range [1-100])", func(t *testing.T) {
		fiveMultiples := []uint8{
			15, 30, 45, 60, 75, 90,
		}
		for _, fiveMultiple := range fiveMultiples {
			out := app.FizzBuzz(fiveMultiple)
			assert.Equal(t, app.FIZZ_BUZZ, out)
		}
	})
}
