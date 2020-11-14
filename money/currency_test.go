package money_test

import (
	"testing"

	"github.com/314159/go-money/money"
	"github.com/stretchr/testify/assert"
)

func TestAddCurrency(t *testing.T) {
	testCases := []struct {
		testName  string
		code      string
		c         money.Currency
		exists    bool
		expectNil bool
	}{
		{
			testName:  "New Currency",
			code:      "AAA",
			c:         money.NewCurrency("AAA", "000", 4),
			exists:    false,
			expectNil: false,
		},
		{
			testName:  "Try Add nil currency",
			code:      "NIL",
			c:         nil,
			exists:    false,
			expectNil: true,
		},
		{
			testName:  "Replace Currency",
			code:      "AAA",
			c:         money.NewCurrency("AAA", "111", 4),
			exists:    true,
			expectNil: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			before := money.FindCurrency(tc.code)
			money.AddCurrency(tc.c)
			after := money.FindCurrency(tc.code)

			assert.Equal(t, tc.exists, before != nil)
			assert.Equal(t, tc.expectNil, after == nil)
		})
	}
}
func TestFindCurrency(t *testing.T) {
	testCases := []struct {
		Case  string
		code  string
		s     string
		isNil bool
		nc    string
		dd    int
	}{
		{
			Case:  "US Dollars",
			code:  "USD",
			isNil: false,
			s:     "USD",
			nc:    "840",
			dd:    2,
		},
		{
			Case:  "Euro Dollars",
			code:  "EUR",
			isNil: false,
			s:     "EUR",
			nc:    "978",
			dd:    2,
		},
		{
			Case:  "Great Brittan Pounds",
			code:  "GBP",
			isNil: false,
			s:     "GBP",
			nc:    "826",
			dd:    2,
		},
		{
			Case:  "Japanese Yen",
			code:  "JPY",
			isNil: false,
			s:     "JPY",
			nc:    "392",
			dd:    0,
		},
		{
			Case:  "Romulan Ale",
			code:  "AAZ",
			isNil: true,
			s:     "",
			nc:    "",
			dd:    0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Case, func(t *testing.T) {
			var s string
			var nc string
			var dd int

			c := money.FindCurrency(tc.s)
			if c != nil {
				s = c.String()
				nc = c.NumericCode()
				dd = c.DecimalDigits()
			}

			assert.Equal(t, tc.isNil, c == nil)
			assert.Equal(t, tc.s, s)
			assert.Equal(t, tc.nc, nc)
			assert.Equal(t, tc.dd, dd)
		})
	}
}
