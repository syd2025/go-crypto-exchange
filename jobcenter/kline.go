package main

import (
	"encoding/json"
	"log"
	"mscoin-common/tools"
	"sync"
	"time"
)

var secretKey = "3u534j6khj565h76"

type OkxResult struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}

type Kline struct {
	wg sync.WaitGroup
}

func NewKline() *Kline {
	return &Kline{}
}

func (k *Kline) Do(period string) {
	k.wg.Add(2)
	// 获取某个币种的k线数据
	go k.getKlineData(period, "BTC-USDT")
	go k.getKlineData(period, "BTC-USDT")
	k.wg.Wait()
}

func (k *Kline) getKlineData(period, instId string) {
	api := "https://www.okx.com/api/v5/market/candles?instId=" + instId + "&bar=" + period
	header := make(map[string]string)
	sign := tools.ComputeHmacSha256("", secretKey)
	header["OK-ACCESS-KEY"] = "<KEY>"
	header["OK-ACCESS-PASSPHRASE"] = "<PASSWORD>"
	header["OK-ACCESS-SIGN"] = sign
	header["OK-ACCESS-TIMESTAMP"] = tools.ISO(time.Now())
	resp, err := tools.GetWithHeader(api, header, "http://127.0.0.1:8888")
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
	log.Println("获取到的K线数据")
	log.Println("instId: ", instId, "period: ", period)
	log.Println("result kline data: ", string(resp))
	log.Println("End")
}
