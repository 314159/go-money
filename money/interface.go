package money

// Money represents a monetary amount
type Money interface {
	// String returns the string value of the money value, with appropriate number of decimal digits.
	String() string

	// Currency returns the currency used by this amount
	Currency() Currency
}

// Currency represents a monetary currency
type Currency interface {
	// NumericCode returns the ISO numeric currency code as a string, e.g. "840"
	NumericCode() string

	// DecimalDigits returns the number of decimal digits used by the currency
	DecimalDigits() int

	// String returns the ISO alphabetic currency code, e.g. "GBP"
	String() string
}
