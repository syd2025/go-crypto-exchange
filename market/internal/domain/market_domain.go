package domain

import (
	"context"
	"grpc-common/market/types/market"
	"market/internal/dao"
	"market/internal/database"
	"market/internal/model"
	"mscoin-common/op"
	"mscoin-common/tools"
	"time"
)

type MarketDomain struct {
	klineRepo *dao.KlineDao
}

func NewMarketDomain(mongoClient *database.MongoClient) *MarketDomain {
	return &MarketDomain{
		klineRepo: dao.NewKlineDao(mongoClient.Db),
	}
}

func (d *MarketDomain) SymbolThumbTrend(coins []*model.ExchangeCoin) []*market.CoinThumb {

	coinThumbs := make([]*market.CoinThumb, len(coins))
	for i, v := range coins {
		from := tools.ZeroTime()
		end := time.Now().Unix()
		klines, err := d.klineRepo.FindBySymbolTime(context.Background(), v.Symbol, "1h", from, end)
		if err != nil {
			coinThumbs[i] = model.DefaultCoinThumb(v.Symbol)
			continue
		}
		if len(klines) <= 0 {
			coinThumbs[i] = model.DefaultCoinThumb(v.Symbol)
			continue
		}

		var high float64 = 0.0
		var low float64 = klines[0].LowestPrice
		var volumes float64 = 0
		var turnover float64 = 0
		length := len(klines)
		// 降序排列
		trend := make([]float64, length)
		// 每次循环获取数据时，去比较highest price 和 high作比较，如果highest price 大于high，则high = highest price
		// 同理，lowest price 和 low 作比较，如果lowest price 小于low，则low = lowest price
		for i := length - 1; i >= 0; i-- {
			trend[i] = klines[i].ClosePrice
			highestPrice := klines[i].HighestPrice
			if highestPrice > high {
				high = highestPrice
			}
			lowestPrice := klines[i].LowestPrice
			if lowestPrice < low {
				low = lowestPrice
			}
			volumes = op.AddN(volumes, klines[i].Volume, 8)
			turnover = op.AddN(turnover, klines[i].Turnover, 8)
		}

		newKline := klines[0]
		oldKline := klines[length-1]
		thumb := newKline.ToCoinThumb(v.Symbol, oldKline)
		thumb.Trend = trend
		thumb.High = high
		thumb.Low = low
		thumb.Volume = volumes
		thumb.Turnover = turnover
		coinThumbs[i] = thumb
	}
	return coinThumbs
}
