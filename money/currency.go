package money

type currency struct {
	alphaCode     string
	numericCode   string
	decimalDigits int
}

func (c currency) AlphaCode() string {
	return c.alphaCode
}

func (c currency) NumericCode() string {
	return c.numericCode
}

func (c currency) DecimalDigits() int {
	return c.decimalDigits
}

func (c currency) String() string {
	return c.AlphaCode()
}

// NewCurrency constructs a currency object
func NewCurrency(alphaCode, numericCode string, decimalDigits int) Currency {
	c := currency{
		alphaCode:     alphaCode,
		numericCode:   numericCode,
		decimalDigits: decimalDigits,
	}
	return &c
}

// currencyList is a map of currencies
type currencyList map[string]Currency

// AddCurrency will add a currency to the list of currencies you can search for.
func AddCurrency(c Currency) {
	if c == nil {
		return
	}

	currencyMap[c.String()] = c
	currencyMap[c.NumericCode()] = c
}

// FindCurrency will return a currency based on the currency code (numeric or alpha)
func FindCurrency(code string) Currency {
	c, ok := currencyMap[code]

	if !ok {
		return nil
	}
	return c
}

func init() {
	for _, c := range list {
		currencyMap[c.String()] = c
		currencyMap[c.NumericCode()] = c
	}
}

var currencyMap currencyList = map[string]Currency{}
var list = []Currency{
	// list retrieved 2020-11-23 from https://en.wikipedia.org/wiki/ISO_4217
	currency{"AED", "784", 2},
	currency{"AFN", "971", 2},
	currency{"ALL", "008", 2},
	currency{"AMD", "051", 2},
	currency{"ANG", "532", 2},
	currency{"AOA", "973", 2},
	currency{"ARS", "032", 2},
	currency{"AUD", "036", 2},
	currency{"AWG", "533", 2},
	currency{"AZN", "944", 2},
	currency{"BAM", "977", 2},
	currency{"BBD", "052", 2},
	currency{"BDT", "050", 2},
	currency{"BGN", "975", 2},
	currency{"BHD", "048", 3},
	currency{"BIF", "108", 0},
	currency{"BMD", "060", 2},
	currency{"BND", "096", 2},
	currency{"BOB", "068", 2},
	currency{"BOV", "984", 2},
	currency{"BRL", "986", 2},
	currency{"BSD", "044", 2},
	currency{"BTN", "064", 2},
	currency{"BWP", "072", 2},
	currency{"BYN", "933", 2},
	currency{"BZD", "084", 2},
	currency{"CAD", "124", 2},
	currency{"CDF", "976", 2},
	currency{"CHE", "947", 2},
	currency{"CHF", "756", 2},
	currency{"CHW", "948", 2},
	currency{"CLF", "990", 4},
	currency{"CLP", "152", 0},
	currency{"CNY", "156", 2},
	currency{"COP", "170", 2},
	currency{"COU", "970", 2},
	currency{"CRC", "188", 2},
	currency{"CUC", "931", 2},
	currency{"CUP", "192", 2},
	currency{"CVE", "132", 2},
	currency{"CZK", "203", 2},
	currency{"DJF", "262", 0},
	currency{"DKK", "208", 2},
	currency{"DOP", "214", 2},
	currency{"DZD", "012", 2},
	currency{"EGP", "818", 2},
	currency{"ERN", "232", 2},
	currency{"ETB", "230", 2},
	currency{"EUR", "978", 2},
	currency{"FJD", "242", 2},
	currency{"FKP", "238", 2},
	currency{"GBP", "826", 2},
	currency{"GEL", "981", 2},
	currency{"GHS", "936", 2},
	currency{"GIP", "292", 2},
	currency{"GMD", "270", 2},
	currency{"GNF", "324", 0},
	currency{"GTQ", "320", 2},
	currency{"GYD", "328", 2},
	currency{"HKD", "344", 2},
	currency{"HNL", "340", 2},
	currency{"HRK", "191", 2},
	currency{"HTG", "332", 2},
	currency{"HUF", "348", 2},
	currency{"IDR", "360", 2},
	currency{"ILS", "376", 2},
	currency{"INR", "356", 2},
	currency{"IQD", "368", 3},
	currency{"IRR", "364", 2},
	currency{"ISK", "352", 0},
	currency{"JMD", "388", 2},
	currency{"JOD", "400", 3},
	currency{"JPY", "392", 0},
	currency{"KES", "404", 2},
	currency{"KGS", "417", 2},
	currency{"KHR", "116", 2},
	currency{"KMF", "174", 0},
	currency{"KPW", "408", 2},
	currency{"KRW", "410", 0},
	currency{"KWD", "414", 3},
	currency{"KYD", "136", 2},
	currency{"KZT", "398", 2},
	currency{"LAK", "418", 2},
	currency{"LBP", "422", 2},
	currency{"LKR", "144", 2},
	currency{"LRD", "430", 2},
	currency{"LSL", "426", 2},
	currency{"LYD", "434", 3},
	currency{"MAD", "504", 2},
	currency{"MDL", "498", 2},
	currency{"MGA", "969", 2},
	currency{"MKD", "807", 2},
	currency{"MMK", "104", 2},
	currency{"MNT", "496", 2},
	currency{"MOP", "446", 2},
	currency{"MRU", "929", 2},
	currency{"MUR", "480", 2},
	currency{"MVR", "462", 2},
	currency{"MWK", "454", 2},
	currency{"MXN", "484", 2},
	currency{"MXV", "979", 2},
	currency{"MYR", "458", 2},
	currency{"MZN", "943", 2},
	currency{"NAD", "516", 2},
	currency{"NGN", "566", 2},
	currency{"NIO", "558", 2},
	currency{"NOK", "578", 2},
	currency{"NPR", "524", 2},
	currency{"NZD", "554", 2},
	currency{"OMR", "512", 3},
	currency{"PAB", "590", 2},
	currency{"PEN", "604", 2},
	currency{"PGK", "598", 2},
	currency{"PHP", "608", 2},
	currency{"PKR", "586", 2},
	currency{"PLN", "985", 2},
	currency{"PYG", "600", 0},
	currency{"QAR", "634", 2},
	currency{"RON", "946", 2},
	currency{"RSD", "941", 2},
	currency{"RUB", "643", 2},
	currency{"RWF", "646", 0},
	currency{"SAR", "682", 2},
	currency{"SBD", "090", 2},
	currency{"SCR", "690", 2},
	currency{"SDG", "938", 2},
	currency{"SEK", "752", 2},
	currency{"SGD", "702", 2},
	currency{"SHP", "654", 2},
	currency{"SLL", "694", 2},
	currency{"SOS", "706", 2},
	currency{"SRD", "968", 2},
	currency{"SSP", "728", 2},
	currency{"STN", "930", 2},
	currency{"SVC", "222", 2},
	currency{"SYP", "760", 2},
	currency{"SZL", "748", 2},
	currency{"THB", "764", 2},
	currency{"TJS", "972", 2},
	currency{"TMT", "934", 2},
	currency{"TND", "788", 3},
	currency{"TOP", "776", 2},
	currency{"TRY", "949", 2},
	currency{"TTD", "780", 2},
	currency{"TWD", "901", 2},
	currency{"TZS", "834", 2},
	currency{"UAH", "980", 2},
	currency{"UGX", "800", 0},
	currency{"USD", "840", 2},
	currency{"USN", "997", 2},
	currency{"UYI", "940", 0},
	currency{"UYU", "858", 2},
	currency{"UYW", "927", 4},
	currency{"UZS", "860", 2},
	currency{"VES", "928", 2},
	currency{"VND", "704", 0},
	currency{"VUV", "548", 0},
	currency{"WST", "882", 2},
	currency{"XAF", "950", 0},
	currency{"XCD", "951", 2},
	currency{"XOF", "952", 0},
	currency{"XPF", "953", 0},
	currency{"YER", "886", 2},
	currency{"ZAR", "710", 2},
	currency{"ZMW", "967", 2},
	currency{"ZWL", "932", 2},
}
