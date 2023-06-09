package money

import (
	"fmt"
	"strconv"
	"strings"
)

type amount struct {
	sign     string
	intPart  string
	fracPart string
}

// New will take a string containing a decimal number, potentially with a fraction part and/or a sign,
// and will validate it and return an "amount" struct with a normalized string value.
func New(sign string) (*amount, error) {
	sign, intPart, fracPart, err := splitParts(sign)
	if err != nil {
		return nil, err
	}

	if intPart == "0" && len(fracPart) == 0 {
		sign = ""
	}

	if len(intPart) > 0 {
		// Make sure integer part is numeric
		if _, err := strconv.ParseInt(intPart, 10, 64); err != nil {
			return nil, err
		}
	}

	if len(fracPart) > 0 {
		// Make sure fractional part is numerc
		if _, err := strconv.ParseInt(fracPart, 10, 64); err != nil {
			return nil, err
		}
	}

	// All good, save parts for use later
	return &amount{sign: sign, intPart: intPart, fracPart: fracPart}, nil
}

func (a *amount) String() string {
	if a == nil {
		return "0"
	}
	if len(a.fracPart) == 0 {
		return a.sign + a.intPart
	}

	return a.sign + a.intPart + decimalPoint + a.fracPart
}

func (a *amount) SetDecimalDigits(dd int) *amount {
	if a == nil {
		return a
	}
	sign := a.sign
	intPart := a.intPart
	fracPart := a.fracPart

	for dp := dd; dp < 0; dp++ {
		if intPart == "0" {
			intPart = ""
		}

		if len(fracPart) == 0 {
			intPart = intPart + "0"
			continue
		}
		intPart = intPart + fracPart[0:]

		fracPart = fracPart[1:]
	}

	switch {
	case dd < 0:
		dd = 0
	case len(fracPart) < dd:
		fracPart = fracPart + strings.Repeat("0", dd-len(fracPart))
	case len(fracPart) > dd:
		fracPart = fracPart[:dd]
	}

	return &amount{sign: sign, intPart: intPart, fracPart: fracPart}
}

func splitParts(a string) (string, string, string, error) {
	parts := strings.Split(a, decimalPoint)

	if len(parts) > 2 {
		return "", "", "", fmt.Errorf("too many decimal points (%d)", len(parts))
	}

	intPart := parts[0]
	fracPart := ""

	if len(parts) > 1 {
		fracPart = parts[1]
	}

	sign := ""
	switch {
	case intPart == "":
		break
	case intPart[0:1] == "+":
		intPart = intPart[1:]
	case intPart[0:1] == "-":
		sign = "-"
		intPart = intPart[1:]
	}
	if len(intPart) == 0 {
		intPart = "0"
	}

	return sign, intPart, fracPart, nil
}
