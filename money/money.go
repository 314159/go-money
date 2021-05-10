package money

import (
	"fmt"
	"strconv"
	"strings"
)

type money struct {
	value    int64
	currency Currency
}

func (m money) String() string {
	dp := m.currency.DecimalDigits()
	av, sign := func(v int64) (int64, string) {
		sign := ""
		if v < 0 {
			sign = "-"
			v = v * -1
		}
		return v, sign
	}(m.value)
	v := strconv.FormatInt(av, 10)

	switch {
	case dp == 0:
		return fmt.Sprintf("%s%s", sign, v)
	case len(v) < dp:
		return fmt.Sprintf("%s0.%0*s", sign, dp-len(v)+1, v)
	case len(v) == dp:
		return fmt.Sprintf("%s0.%s", sign, v)
	default:
		integerPart := v[:len(v)-dp]
		decimalPart := v[len(v)-dp:]
		return fmt.Sprintf("%s%s.%s", sign, integerPart, decimalPart)
	}
}
func (m money) Currency() Currency {
	return m.currency
}

// NewMoney constructs a Money object from the amount string and the currency.
func NewMoney(amount string, c Currency) (Money, error) {
	parts := strings.Split(amount, ".")

	var integerPart, decimalPart string

	switch {
	case len(parts) == 1:
		integerPart = parts[0]
		decimalPart = ""
	case len(parts) == 2:
		integerPart = parts[0]
		decimalPart = parts[1]
	default:
		return nil, fmt.Errorf("invalid Money amount: %s", amount)
	}

	// TODO: strip off trailing zeros from decimalPart so "1.00" dp 0 is okay

	l := len(decimalPart)
	dd := c.DecimalDigits()

	switch {
	case l < dd:
		decimalPart = decimalPart + fmt.Sprintf("%0*d", dd-l, 0)
	case l == dd:
		break
	default:
		// This would cause loss of precision!
		return nil, fmt.Errorf("invalid Money amount: %s", amount)
	}

	v, err := strconv.ParseInt(integerPart+decimalPart, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid Money amount: %s", amount)
	}

	m := money{value: v, currency: c}
	return &m, nil
}
