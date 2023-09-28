package util

// 支持的货币常量
const (
	EUR = "EUR"
	USD = "USD"
	CAD = "CAD"
)

// 是否支持该货币
func IsSupportCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
