package money_test

import (
	"testing"

	"github.com/314159/go-money/money"
	"github.com/stretchr/testify/assert"
)

func TestMoney(t *testing.T) {
	usDollar := money.NewCurrency("USD", "840", 2)
	jpYen := money.NewCurrency("JPY", "139", 0)

	testCases := []struct {
		Case     string
		Amount   string
		Currency money.Currency
		Expects  string
		ExpCurr  money.Currency
		ExpErr   bool
		ErrMsg   string
	}{
		{
			Case:     "123 USD",
			Amount:   "123",
			Currency: usDollar,
			Expects:  "123.00",
			ExpCurr:  usDollar,
			ExpErr:   false,
		},
		{
			Case:     "-123 USD",
			Amount:   "-123",
			Currency: usDollar,
			Expects:  "-123.00",
			ExpCurr:  usDollar,
			ExpErr:   false,
		},
		{
			Case:     "-.01 USD",
			Amount:   "-.01",
			Currency: usDollar,
			Expects:  "-0.01",
			ExpCurr:  usDollar,
			ExpErr:   false,
		},
		{
			Case:     "-.11 USD",
			Amount:   "-.11",
			Currency: usDollar,
			Expects:  "-0.11",
			ExpCurr:  usDollar,
			ExpErr:   false,
		},
		{
			Case:     "-0.11 USD",
			Amount:   "-0.11",
			Currency: usDollar,
			Expects:  "-0.11",
			ExpCurr:  usDollar,
			ExpErr:   false,
		},
		{
			Case:     "empty",
			Amount:   "",
			Currency: usDollar,
			Expects:  "0.00",
			ExpCurr:  usDollar,
			ExpErr:   false,
		},
		{
			Case:     "123.10 USD",
			Amount:   "123.10",
			Currency: usDollar,
			Expects:  "123.10",
			ExpCurr:  usDollar,
			ExpErr:   false,
		},
		{
			Case:     "123.1 USD",
			Amount:   "123.1",
			Currency: usDollar,
			Expects:  "123.10",
			ExpCurr:  usDollar,
			ExpErr:   false,
		},
		{
			Case:     "123.123 USD",
			Amount:   "123.123",
			Currency: usDollar,
			Expects:  "",
			ExpCurr:  nil,
			ExpErr:   true,
			ErrMsg:   "Invalid Money amount: 123.123",
		},
		{
			Case:     "123.12.3 USD",
			Amount:   "123.12.3",
			Currency: usDollar,
			Expects:  "",
			ExpCurr:  nil,
			ExpErr:   true,
			ErrMsg:   "Invalid Money amount: 123.12.3",
		},
		{
			Case:     "Non-numeric-amount",
			Amount:   "12.3X",
			Currency: usDollar,
			Expects:  "",
			ExpCurr:  nil,
			ExpErr:   true,
			ErrMsg:   "Invalid Money amount: 12.3X",
		},
		{
			Case:     "123 JPY",
			Amount:   "123",
			Currency: jpYen,
			Expects:  "123",
			ExpCurr:  jpYen,
			ExpErr:   false,
		},
		{
			Case:     "123.10 JPY",
			Amount:   "123.10",
			Currency: jpYen,
			Expects:  "",
			ExpCurr:  nil,
			ExpErr:   true,
			ErrMsg:   "Invalid Money amount: 123.10",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Case, func(t *testing.T) {
			var actual string
			var c money.Currency

			m, err := money.NewMoney(tc.Amount, tc.Currency)
			if m != nil {
				actual = m.String()
				c = m.Currency()
			}

			hadError := err != nil
			var es string

			if err != nil {
				es = err.Error()
			}

			assert.Equal(t, tc.Expects, actual)
			assert.Equal(t, tc.ExpCurr, c)
			assert.Equal(t, tc.ExpErr, hadError)
			assert.Equal(t, tc.ErrMsg, es)
		})
	}
}
