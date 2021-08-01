package app_test

import (
	"fizzbuzz/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("All three multiples (1-100)", func(t *testing.T) {
		threeMultiples := []uint8{
			3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51, 54, 57, 60,
			63, 66, 69, 72, 75, 78, 81, 84, 87, 90, 93, 96, 99,
		}
		for _, threeMultiple := range threeMultiples {
			out := app.FizzBuzz(threeMultiple)
			assert.Equal(t, app.FIZZ, out)
		}
	})
}
