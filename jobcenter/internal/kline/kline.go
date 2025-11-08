package kline

import (
	"encoding/json"
	"jobcenter/internal/config"
	"jobcenter/internal/domain"
	"jobcenter/internal/svc"
	"log"
	"mscoin-common/tools"
	"sync"
	"time"
)

type OkxResult struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}

type Kline struct {
	wg          sync.WaitGroup
	config      config.OkxConfig
	klineDomain *domain.KlineDomain
	queueDomain *domain.QueueDomain
}

func NewKline(cfg config.OkxConfig, ctx *svc.ServiceContext) *Kline {
	return &Kline{
		config:      cfg,
		klineDomain: domain.NewKlineDomain(ctx.MongoClient),
		queueDomain: domain.NewQueueDomain(ctx.KafkaClient),
	}
}

func (k *Kline) Do(period string) {
	k.wg.Add(2)
	// 获取某个币种的k线数据
	go k.getKlineData("BTC-USDT", "BTC/USDT", period)
	go k.getKlineData("ETH-USDT", "ETH/USDT", period)
	k.wg.Wait()
}

func (k *Kline) getKlineData(instId, symbol, period string) {
	api := k.config.Host + "/api/v5/market/candles?instId=" + instId + "&bar=" + period
	header := make(map[string]string)
	sign := tools.ComputeHmacSha256("", k.config.SecretKey)
	header["OK-ACCESS-KEY"] = k.config.ApiKey
	header["OK-ACCESS-PASSPHRASE"] = k.config.Pass
	header["OK-ACCESS-SIGN"] = sign
	header["OK-ACCESS-TIMESTAMP"] = tools.ISO(time.Now())
	resp, err := tools.GetWithHeader(api, header, k.config.Proxy)
	if err != nil {
		k.wg.Done()
		return
	}
	var result = &OkxResult{}
	err = json.Unmarshal(resp, result)
	if err != nil {
		k.wg.Done()
		return
	}
	if result.Code == "0" {
		// 代表成功
		k.klineDomain.SaveBatch(result.Data, symbol, period)
		if period == "1m" {
			if len(result.Data) > 0 {
				k.queueDomain.Send1mKline(result.Data[0], symbol)
			}
		}
	}
	k.wg.Done()
	log.Println("============end================")
}
