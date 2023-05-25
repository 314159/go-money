package currencylist_test

import (
	"testing"

	"github.com/314159/go-money/currency"
	"github.com/314159/go-money/currencylist"
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
			c:         currency.New("AAA", "000", 4),
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
			c:         currency.New("AAA", "111", 4),
			exists:    true,
			expectNil: false,
		},
	}

	l := currencylist.New()
	l.AddMany(currency.StandardList)

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			before := l.FindCurrency(tc.code)
			l.AddCurrency(tc.c)
			after := l.FindCurrency(tc.code)

			if tc.exists {
				assert.NotNil(t, before)
			} else {
				assert.Nil(t, before)
			}

			if tc.expectNil {
				assert.Nil(t, after)
			} else {
				assert.NotNil(t, after)
			}
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
			l := currencylist.New()
			l.AddMany(currency.StandardList)
			c := l.FindCurrency(tc.s)

			if tc.isNil {
				assert.Nil(t, c)
			} else {
				assert.NotNil(t, c)
				assert.Equal(t, tc.s, c.String())
				assert.Equal(t, tc.nc, c.NumericCode())
				assert.Equal(t, tc.dd, c.DecimalDigits())
			}
		})
	}
}
