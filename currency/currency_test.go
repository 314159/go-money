package currency_test

import (
	"testing"

	"github.com/314159/go-money/currency"
	"github.com/stretchr/testify/assert"
)

func TestCurrency(t *testing.T) {
	testCases := []struct {
		testName      string
		alphaCode     string
		numericCode   string
		decimalDigits int
	}{
		{
			testName:      "abc 010 3",
			alphaCode:     "abc",
			numericCode:   "010",
			decimalDigits: 3,
		},
		{
			testName:      "USD, 840 2",
			alphaCode:     "USD",
			numericCode:   "840",
			decimalDigits: 2,
		},
		{
			testName:      "JPY, 392 0",
			alphaCode:     "JPY",
			numericCode:   "392",
			decimalDigits: 0,
		},
		{
			testName:      "øøø, øøø -1",
			alphaCode:     "øøø",
			numericCode:   "øøø",
			decimalDigits: -1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			c := currency.New(tc.alphaCode, tc.numericCode, tc.decimalDigits)
			assert.Equal(t, tc.alphaCode, c.AlphaCode())
			assert.Equal(t, tc.numericCode, c.NumericCode())
			assert.Equal(t, tc.decimalDigits, c.DecimalDigits())
			assert.Equal(t, tc.alphaCode, c.String())
		})
	}
}
