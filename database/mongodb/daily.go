package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"snapscale-api/apiClient/chain"
	_type "snapscale-api/apiClient/type"
	"strconv"
	"strings"
	"time"
)

var Daily *collection

func dailyInit() {
	Daily = &collection{}
	Daily.p = MongoDb.Collection("daily")
	go dailyOnce()
}

func dailyOnce() {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	tomorrow := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	next := tomorrow.Unix() - now.Unix()
	for range time.Tick(time.Second * time.Duration(next)) {
		dailyDo(today)
		dailyOnce()
	}
}

func dailyDo(tm time.Time) {
	// 转账量
	dailyTransactionsAmount(tm)
	// 交易次数
	dailyTransactions(tm)
	// 用户数
	dailyAccount(tm)
	// 合约数
	dailyContract(tm)
}

func dailyTransactionsAmount(tm time.Time) {
	tomorrow := time.Date(tm.Year(), tm.Month(), tm.Day()+1, 0, 0, 0, 0, tm.Location())
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, _ := Transactions.p.Find(ctx, bson.D{{"$and",
		bson.A{
			bson.D{{
				"createdAt",
				bson.D{{"$gte", tm}},
			}},
			bson.D{{
				"createdAt",
				bson.D{{"$lt", tomorrow}},
			}},
		},
	}})

	var sum float64
	sum = 0

	type R struct {
		Actions []struct {
			Account string
			Name    string
			Data    struct {
				Quantity string
			}
		}
	}

	for cur.Next(ctx) {
		var result R
		_ = cur.Decode(&result)
		for _, action := range result.Actions {
			if action.Account == "eosio.token" && action.Name == "transfer" {
				arr := strings.Split(action.Data.Quantity, " ")
				N, _ := strconv.ParseFloat(arr[0], 64)
				sum += N
			}
		}
	}
	_, _ = Daily.p.InsertOne(ctx, bson.M{"type": "transactionAmount", "value": fmt.Sprintf("%.2f", sum), "time": tm.Unix(), "timeUTC": tm.Format("2006-01-02")})
}

func dailyTransactions(tm time.Time) {
	tomorrow := time.Date(tm.Year(), tm.Month(), tm.Day()+1, 0, 0, 0, 0, tm.Location())
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	transactions, _ := Transactions.p.CountDocuments(ctx, bson.D{{"$and",
		bson.A{
			bson.D{{
				"createdAt",
				bson.D{{"$gte", tm}},
			}},
			bson.D{{
				"createdAt",
				bson.D{{"$lt", tomorrow}},
			}},
		},
	}})
	_, _ = Daily.p.InsertOne(ctx, bson.M{"type": "transactions", "value": transactions, "time": tm.Unix(), "timeUTC": tm.Format("2006-01-02")})
}

func dailyAccount(tm time.Time) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	accounts, _ := Accounts.p.EstimatedDocumentCount(ctx)
	_, _ = Daily.p.InsertOne(ctx, bson.M{"type": "accounts", "value": accounts, "time": tm.Unix(), "timeUTC": tm.Format("2006-01-02")})
}

func dailyContract(tm time.Time) {
	t1 := &_type.XAbiHashResponse{}
	_, _, _ = chain.GetTableRows("eosio", "abihash", "eosio", t1)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	length := len(t1.Rows)
	_, _ = Daily.p.InsertOne(ctx, bson.M{"type": "contracts", "value": length, "time": tm.Unix(), "timeUTC": tm.Format("2006-01-02")})
}
