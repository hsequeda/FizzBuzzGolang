package app_test

import (
	"fizzbuzz/app"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("All three multiples (only) (range [1-100])", func(t *testing.T) {
		threeMultiples := []uint8{
			3, 6, 9, 12, 18, 21, 24, 27, 33, 36, 39, 42, 48, 51, 54, 57,
			63, 66, 69, 72, 78, 81, 84, 87, 93, 96, 99,
		}
		for _, threeMultiple := range threeMultiples {
			out, err := app.FizzBuzz(threeMultiple)
			require.NoError(t, err)
			assert.Equal(t, app.FizzCase, out)
		}
	})

	t.Run("All five multiples (only) (range [1-100])", func(t *testing.T) {
		fiveMultiples := []uint8{
			5, 10, 20, 25, 35, 40, 50, 55, 65, 70, 80, 85, 95, 100,
		}
		for _, fiveMultiple := range fiveMultiples {
			out, err := app.FizzBuzz(fiveMultiple)
			require.NoError(t, err)
			assert.Equal(t, app.BuzzCase, out)
		}
	})

	t.Run("All three && five multiples (only) (range [1-100])", func(t *testing.T) {
		fiveMultiples := []uint8{
			15, 30, 45, 60, 75, 90,
		}
		for _, fiveMultiple := range fiveMultiples {
			out, err := app.FizzBuzz(fiveMultiple)
			require.NoError(t, err)
			assert.Equal(t, app.FizzBuzzCase, out)
		}
	})

	t.Run("All number that isn't multiples of three or five (range [1-100])", func(t *testing.T) {
		noMultiplesOfFiveOrThree := []uint8{
			1, 2, 4, 7, 8, 11, 13, 14, 16, 17, 19, 22, 23, 26, 28, 29, 31, 32, 34, 37, 38, 41,
			43, 44, 46, 47, 49, 52, 53, 56, 58, 59, 61, 62, 64, 67, 68, 71, 73, 74, 76, 77,
			79, 82, 83, 86, 88, 89, 91, 92, 94, 97, 98,
		}
		for _, number := range noMultiplesOfFiveOrThree {
			out, err := app.FizzBuzz(number)
			require.NoError(t, err)
			assert.Equal(t, fmt.Sprint(number), out)
		}
	})

	t.Run("Error: Input == 0 || Input > 100", func(t *testing.T) {
		_, err := app.FizzBuzz(0)
		assert.NotNil(t, err)
		_, err = app.FizzBuzz(101)
		assert.NotNil(t, err)
		_, err = app.FizzBuzz(102)
		assert.NotNil(t, err)
	})
}
