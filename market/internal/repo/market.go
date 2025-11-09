package repo

import (
	"context"
	"market/internal/model"
)

type KlineRepo interface {
	FindSymbol(ctx context.Context, symbol, period, count string) ([]*model.Kline, error)
	FindBySymbolTime(ctx context.Context, symbol, period string, from, end int64) ([]*model.Kline, error)
}
