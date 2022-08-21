package util

const (
	USD = "USD"
	EUR = "EUR"
	AUD = "AUD"
	CAD = "CAD"
)

func IsCurrencySupported(currency string) bool {
	switch currency {
	case USD, EUR, AUD, CAD:
		return true
	}

	return false
}
