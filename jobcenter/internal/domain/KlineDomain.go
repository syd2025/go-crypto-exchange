package domain

import (
	"context"
	"jobcenter/internal/dao"
	"jobcenter/internal/database"
	"jobcenter/internal/model"
	"jobcenter/internal/repo"
	"log"
)

type KlineDomain struct {
	klineRepo repo.KlineRepo
}

func NewKlineDomain(client *database.MongoClient) *KlineDomain {
	return &KlineDomain{
		klineRepo: dao.NewKlineDao(client.Db),
	}
}

func (d *KlineDomain) SaveBatch(data [][]string, symbol, period string) {
	kline := make([]*model.Kline, len(data))
	for i, v := range data {
		kline[i] = model.NewKline(v, "1m")
	}
	// 删除最老的数据
	err := d.klineRepo.DeleteGtTime(context.Background(), kline[len(data)-1].Time, symbol, period)
	if err != nil {
		log.Println("删除mongodb数据失败")
		return
	}
	err = d.klineRepo.SaveBatch(context.Background(), kline, symbol, period)
	if err != nil {
		log.Println("批量写入mongodb失败")
	}
}
