package http

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"net/http"
	"snapscale-api/apiClient/chain"
	"snapscale-api/apiClient/history"
	_type "snapscale-api/apiClient/type"
	"snapscale-api/database/mongodb"
	"strconv"
)

type router struct {
	M map[string]func(http.ResponseWriter, *http.Request)
}

var Router *router

func (x router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rt, ok := x.M[r.URL.Path]
	if ok {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		rt(w, r)
	} else {
		_ = r.Body.Close()
	}
}

func init() {
	Router = &router{}
	Router.M = make(map[string]func(http.ResponseWriter, *http.Request))
	Router.M["/api/account/"] = account
	Router.M["/api/block/"] = block
	Router.M["/api/transaction/"] = transaction
	Router.M["/api/charts/"] = charts
}

func account(w http.ResponseWriter, r *http.Request) {
	b, e := ioutil.ReadAll(r.Body)

	if e != nil {
		_ = r.Body.Close()
		return
	}

	req := &accountRequest{}
	e = json.Unmarshal(b, req)
	if e != nil {
		_ = r.Body.Close()
		return
	}

	res := &_type.AccountResponse{}
	res, e1, e2 := chain.GetAccount(req.Name)

	res2 := &_type.ActionsResponse{}
	res2, _, _ = history.GetActions(req.Name)
	if e1 == nil && e2 == nil {
		data, e3 := json.Marshal(struct {
			Info    *_type.AccountResponse `json:"info"`
			Actions *_type.ActionsResponse `json:"actions"`
		}{res, res2})
		if e3 != nil {
			_ = r.Body.Close()
			return
		}

		_, _ = w.Write(data)
	}
}

func block(w http.ResponseWriter, r *http.Request) {
	b, e := ioutil.ReadAll(r.Body)

	if e != nil {
		_ = r.Body.Close()
		return
	}

	req := &blockRequest{}
	e = json.Unmarshal(b, req)
	if e != nil {
		_ = r.Body.Close()
		return
	}

	res := &_type.BlockResponse{}
	res, e1, e2 := chain.GetBlock(strconv.FormatInt(req.Num, 10))
	if e1 == nil && e2 == nil {
		data, e3 := json.Marshal(res)
		if e3 != nil {
			_ = r.Body.Close()
			return
		}

		_, _ = w.Write(data)
	}
}

func transaction(w http.ResponseWriter, r *http.Request) {
	b, e := ioutil.ReadAll(r.Body)

	if e != nil {
		_ = r.Body.Close()
		return
	}

	req := &transactionRequest{}
	e = json.Unmarshal(b, req)
	if e != nil {
		_ = r.Body.Close()
		return
	}

	var result bson.M
	var result2 bson.M
	e1 := mongodb.Transactions.FindOne(bson.M{"trx_id": req.Hash}).Decode(&result)
	e2 := mongodb.TransactionTraces.FindOne(bson.M{"id": req.Hash}).Decode(&result2)

	if e1 == nil && e2 == nil {
		data, e3 := json.Marshal(struct {
			Trx   interface{} `json:"trx"`
			Trace interface{} `json:"trace"`
		}{result, result2})
		if e3 != nil {
			_ = r.Body.Close()
			return
		}

		_, _ = w.Write(data)
	}
}

func charts(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write(mongodb.DailyInfo())
}
