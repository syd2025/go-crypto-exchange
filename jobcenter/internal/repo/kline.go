package repo

import (
	"context"
	"jobcenter/internal/model"
)

type KlineRepo interface {
	SaveBatch(ctx context.Context, data []*model.Kline, symbal, period string) error
	DeleteGtTime(background context.Context, time int64, symbol string, period string) error
}
