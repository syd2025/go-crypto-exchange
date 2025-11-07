package database

import (
	"context"
	"log"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Url      string
	Username string
	Password string
	DataBase string
}

type MongoClient struct {
	cli *mongo.Client
	Db  *mongo.Database
}

func ConnectMongo(c MongoConfig) *MongoClient {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credential := options.Credential{
		Username: c.Username,
		Password: c.Password,
	}

	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI("mongodb://127.0.0.1:27017").SetAuth(credential),
	)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	database := client.Database(c.DataBase)
	return &MongoClient{cli: client, Db: database}
}

func (m *MongoClient) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := m.cli.Disconnect(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println("mongo 断开连接")
}
