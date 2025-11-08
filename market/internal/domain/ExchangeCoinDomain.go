package domain

import (
	"context"
	"market/internal/dao"
	"market/internal/model"
	"mscoin-common/msdb"
	"mscoin-common/msdb/gorms"
)

type ExchangeCoinDomain struct {
	db *msdb.MsDB
}

func NewExchangeCoinDomain(db *msdb.MsDB) *ExchangeCoinDomain {
	return &ExchangeCoinDomain{
		db: db,
	}
}

func (d *ExchangeCoinDomain) FindVisible(ctx context.Context) (list []*model.ExchangeCoin, err error) {
	gormConn := gorms.New(d.db.Conn)
	exchangeCoinDao := dao.NewExchangeCoinDao(gormConn)
	return exchangeCoinDao.FindVisible()
}
