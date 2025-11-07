package domain

type ExchangeRateDomain struct {
}

func NewExchangeRateDomain() *ExchangeRateDomain {
	return &ExchangeRateDomain{}
}

func (d *ExchangeRateDomain) UsdRate(unit string) float64 {
	switch unit {
	case "CNY":
		return 7
	case "JPY":
		return 1234.456
	}
	return 0
}
