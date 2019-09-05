package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"snapscale-api/config"
	"time"
)

var MongoClient *mongo.Client
var MongoDb *mongo.Database

type collection struct {
	p *mongo.Collection
}

func (c collection) FindOne(filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return c.p.FindOne(ctx, filter, opts...)
}

var Accounts *collection
var Transactions *collection
var TransactionTraces *collection
var ActionTraces *collection

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoConfig))

	if err == nil {
		MongoClient = client
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		_ = MongoClient.Connect(ctx)
	} else {
		log.Fatal(err)
	}

	MongoDb = MongoClient.Database("EOS")

	Accounts = &collection{}
	Accounts.p = MongoDb.Collection("accounts")

	Transactions = &collection{}
	Transactions.p = MongoDb.Collection("transactions")

	TransactionTraces = &collection{}
	TransactionTraces.p = MongoDb.Collection("transaction_traces")

	ActionTraces = &collection{}
	ActionTraces.p = MongoDb.Collection("action_traces")

	dailyInit()
}

func Start() {}
