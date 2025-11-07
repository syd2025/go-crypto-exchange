package dao

import (
	"context"
	"market/internal/model"
	"mscoin-common/msdb/gorms"
)

type ExchangeCoinDao struct {
	coin *gorms.GormConn
}

func NewExchangeCoinDao(coin *gorms.GormConn) *ExchangeCoinDao {
	return &ExchangeCoinDao{
		coin: coin,
	}
}

func (e *ExchangeCoinDao) FindVisible() (list []*model.ExchangeCoin, err error) {
	err = e.coin.Session(context.Background()).Model(&model.ExchangeCoin{}).Where("visible = ?", 1).Find(&list).Error
	return
}
