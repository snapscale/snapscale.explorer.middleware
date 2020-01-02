package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	createIndex()

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

	ctx1, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, _ = Daily.p.InsertOne(ctx1, bson.M{"type": "transactionAmount", "value": fmt.Sprintf("%.2f", sum), "time": tm.Unix(), "timeUTC": tm.Format("2006-01-02"), "xid": tm.Format("2006-01-02") + "_transactionAmount"})
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
	_, _ = Daily.p.InsertOne(ctx, bson.M{"type": "transactions", "value": transactions, "time": tm.Unix(), "timeUTC": tm.Format("2006-01-02"), "xid": tm.Format("2006-01-02") + "_transactions"})
}

func dailyAccount(tm time.Time) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	accounts, _ := Accounts.p.EstimatedDocumentCount(ctx)
	_, _ = Daily.p.InsertOne(ctx, bson.M{"type": "accounts", "value": accounts, "time": tm.Unix(), "timeUTC": tm.Format("2006-01-02"), "xid": tm.Format("2006-01-02") + "_accounts"})
}

func dailyContract(tm time.Time) {
	t1 := &_type.XAbiHashResponse{}
	_, _, _ = chain.GetTableRows("eosio", "abihash", "eosio", t1)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	length := len(t1.Rows)
	_, _ = Daily.p.InsertOne(ctx, bson.M{"type": "contracts", "value": length, "time": tm.Unix(), "timeUTC": tm.Format("2006-01-02"), "xid": tm.Format("2006-01-02") + "_contracts"})
}

func DailyInfo() []byte {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	now := time.Now()
	lastTarget := time.Date(now.Year(), now.Month(), now.Day()-10, 0, 0, 0, 0, now.Location())
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	optionx := options.Find()
	optionx.SetSort(bson.D{{"time", 1}})

	cur1, _ := Daily.p.Find(ctx, bson.D{{"$and",
		bson.A{
			bson.D{{
				"time",
				bson.D{{"$gte", lastTarget.Unix()}},
			}},
			bson.D{{
				"time",
				bson.D{{"$lte", today.Unix()}},
			}},
			bson.M{"type": "transactionAmount"},
		},
	}}, optionx)

	cur2, _ := Daily.p.Find(ctx, bson.D{{"$and",
		bson.A{
			bson.D{{
				"time",
				bson.D{{"$gte", lastTarget.Unix()}},
			}},
			bson.D{{
				"time",
				bson.D{{"$lte", today.Unix()}},
			}},
			bson.M{"type": "transactions"},
		},
	}}, optionx)

	cur3, _ := Daily.p.Find(ctx, bson.D{{"$and",
		bson.A{
			bson.D{{
				"time",
				bson.D{{"$gte", lastTarget.Unix()}},
			}},
			bson.D{{
				"time",
				bson.D{{"$lte", today.Unix()}},
			}},
			bson.M{"type": "accounts"},
		},
	}}, optionx)

	cur4, _ := Daily.p.Find(ctx, bson.D{{"$and",
		bson.A{
			bson.D{{
				"time",
				bson.D{{"$gte", lastTarget.Unix()}},
			}},
			bson.D{{
				"time",
				bson.D{{"$lte", today.Unix()}},
			}},
			bson.M{"type": "contracts"},
		},
	}}, optionx)

	type Normal struct {
		TimeUTC string `bson:"timeUTC"`
		Value   interface{}
		Time    int64
		Type    string
	}

	var result1 []string
	for cur1.Next(ctx) {
		var x Normal
		_ = cur1.Decode(&x)
		result1 = append(result1, x.Value.(string))
	}

	var result2 []int64
	for cur2.Next(ctx) {
		var x Normal
		_ = cur2.Decode(&x)
		result2 = append(result2, x.Value.(int64))
	}

	var result3 []int64
	for cur3.Next(ctx) {
		var x Normal
		_ = cur3.Decode(&x)
		result3 = append(result3, x.Value.(int64))
	}

	var result4 []int32

	for cur4.Next(ctx) {
		var x Normal
		_ = cur4.Decode(&x)
		result4 = append(result4, x.Value.(int32))
	}

	type results struct {
		RT1 []string `json:"dt1"`
		RT2 []int64  `json:"dt2"`
		RT3 []int64  `json:"dt3"`
		RT4 []int32  `json:"dt4"`
	}
	var result = results{}

	for i := len(result1); i < 10; i++ {
		result1 = append([]string{"0"}, result1...)
	}
	for i := len(result2); i < 10; i++ {
		result2 = append([]int64{0}, result2...)
	}
	for i := len(result3); i < 10; i++ {
		result3 = append([]int64{0}, result3...)
	}
	for i := len(result4); i < 10; i++ {
		result4 = append([]int32{0}, result4...)
	}

	result.RT1 = result1
	result.RT2 = result2
	result.RT3 = result3
	result.RT4 = result4

	zz, _ := json.Marshal(result)

	return zz
}

func createIndex() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	background := true

	mod := mongo.IndexModel{
		Keys: bson.M{
			"createdAt": 1, // index in ascending order
		},
		Options: &options.IndexOptions{
			Background: &background,
		},
	}
	ind, err := Transactions.p.Indexes().CreateOne(ctx, mod)

	fmt.Printf("createIndex index name %s, erro %v\n", ind, err)

	mod1 := mongo.IndexModel{
		Keys: bson.M{
			"xid": 1, // index in ascending order
		},
		Options: &options.IndexOptions{
			Background: &background,
		},
	}
	ind, err = Daily.p.Indexes().CreateOne(ctx, mod1)

	fmt.Printf("createIndex index name %s, erro %v\n", ind, err)
}
