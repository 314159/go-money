package currencylist

import (
	"github.com/314159/go-money/currency"
	"github.com/314159/go-money/money"
)

// currencyList is a map of currencies
type currencyList struct {
	list map[string]money.Currency
}

func New() *currencyList {
	return &currencyList{
		list: make(map[string]money.Currency, 0),
	}
}

// AddCurrency will add a currency to the list of currencies you can search for.
func (l *currencyList) AddCurrency(c money.Currency) {
	if c == nil {
		return
	}

	l.list[c.String()] = c
	l.list[c.NumericCode()] = c
}

// FindCurrency will return a currency based on the currency code (numeric or alpha)
func (l *currencyList) FindCurrency(code string) money.Currency {
	c, ok := l.list[code]

	if !ok {
		return nil
	}
	return c
}

// AddMany will add a slice of Currencys to the list
func (l *currencyList) AddMany([]money.Currency) {
	for _, c := range currency.StandardList {
		l.AddCurrency(c)
	}
}
