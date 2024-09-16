package util

const (
	USD   = "USD"
	EUR   = "EUR"
	TENGE = "TENGE"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, TENGE:
		return true
	}
	return false
}
