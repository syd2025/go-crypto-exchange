package dao

import (
	"context"
	"jobcenter/internal/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type KlineDao struct {
	db *mongo.Database
}

func NewKlineDao(db *mongo.Database) *KlineDao {
	return &KlineDao{
		db: db,
	}
}

func (d *KlineDao) DeleteGtTime(ctx context.Context, time int64, symbol string, period string) error {
	collection := d.db.Collection("exchange_kline_" + symbol + "_" + period)
	deleteResult, err := collection.DeleteMany(ctx, bson.D{{Key: "time", Value: bson.D{{Key: "$gte", Value: time}}}})
	if err != nil {
		log.Printf("删除表%s,数量: %d \n", "exchange_kline_"+symbol+"_"+period, deleteResult.DeletedCount)
	}
	return err
}

// 批量插入
func (d *KlineDao) SaveBatch(ctx context.Context, data []*model.Kline, symbol, period string) error {
	mk := &model.Kline{}
	collection := d.db.Collection(mk.Table(symbol, period))
	ds := make([]interface{}, len(data))
	for i, v := range data {
		ds[i] = v
	}
	_, err := collection.InsertMany(ctx, ds)
	if err != nil {
		return err
	}
	return nil
}
