package domain

import "market/internal/repo"

type ExchangeCoinDomain struct {
	exchangeCoinRepo repo.ExchangeCoinRepo
}

func NewExchangeCoinDomain(exchangeCoinRepo repo.ExchangeCoinRepo) *ExchangeCoinDomain {
	return &ExchangeCoinDomain{
		exchangeCoinRepo: exchangeCoinRepo,
	}
}
