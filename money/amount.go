package money

import (
	"fmt"
	"strconv"
	"strings"
)

type amount struct {
	s string
}

func New(s string) (*amount, error) {
	s, i, d, err := splitParts(s)
	if err != nil {
		return nil, err
	}

	if i == "0" && len(d) == 0 {
		return &amount{s: "0"}, nil // We don't want "-0", just "0"
	}

	if len(i) > 0 {
		if _, err := strconv.ParseInt(i, 10, 64); err != nil {
			return nil, err
		}
	}
	if len(d) == 0 {
		return &amount{s: s + i}, nil
	}

	if _, err := strconv.ParseInt(d, 10, 64); err != nil {
		return nil, err
	}

	return &amount{s: s + i + decimalPoint + d}, nil
}

func (a *amount) String() string {
	if a == nil {
		return "0"
	}
	return a.s
}

func (a *amount) SetDecimalDigits(dd int) *amount {
	if a == nil {
		return a
	}
	s, i, d, err := splitParts(a.s)
	if err != nil {
		return nil
	}

	for dp := dd; dp < 0; dp++ {
		if i == "0" {
			i = ""
		}

		if len(d) == 0 {
			i = i + "0"
			continue
		}
		i = i + d[0:]

		d = d[1:]
	}

	switch {
	case dd < 0:
		dd = 0
	case len(d) < dd:
		d = d + strings.Repeat("0", dd-len(d))
	case len(d) > dd:
		d = d[:dd]
	}

	if d == "" {
		return &amount{s: s + i}
	}

	return &amount{s: s + i + decimalPoint + d}
}

func splitParts(a string) (string, string, string, error) {
	parts := strings.Split(a, decimalPoint)

	if len(parts) > 2 {
		return "", "", "", fmt.Errorf("too many decimal points (%d)", len(parts))
	}

	i := parts[0]
	d := ""

	if len(parts) > 1 {
		d = parts[1]
	}

	s := ""
	switch {
	case i == "":
		break
	case i[0:1] == "+":
		i = i[1:]
	case i[0:1] == "-":
		s = "-"
		i = i[1:]
	}
	if len(i) == 0 {
		i = "0"
	}

	return s, i, d, nil
}
