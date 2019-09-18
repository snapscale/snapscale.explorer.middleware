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
	"snapscale-api/libs/log"
	"strconv"
)

type router struct {
	M map[string]func(http.ResponseWriter, *http.Request)
}

var Router *router

func (x router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.S.Printf("Request [%s]: %s", r.Method, r.URL.Path)
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
	Router.M["/api/abi/"] = abi
	Router.M["/api/contract/"] = contract
	Router.M["/api/actions/"] = actions
	Router.M["/health/"] = health
}

func health(w http.ResponseWriter, r *http.Request) {
	_ = r.Body.Close()
}

func actions(w http.ResponseWriter, r *http.Request) {
	b, e := ioutil.ReadAll(r.Body)

	if e != nil {
		_ = r.Body.Close()
		return
	}

	req := &actionsRequest{}
	e = json.Unmarshal(b, req)
	if e != nil {
		_ = r.Body.Close()
		return
	}

	t1, e1, e2 := history.GetActions2(req.AccountName, req.Pos, req.Offset)

	if e1 == nil && e2 == nil {
		data, e3 := json.Marshal(t1)
		if e3 == nil {
			_, _ = w.Write(data)
		}
	}
}

func contract(w http.ResponseWriter, r *http.Request) {
	b, e := ioutil.ReadAll(r.Body)

	if e != nil {
		_ = r.Body.Close()
		return
	}

	req := &contractRequest{}
	e = json.Unmarshal(b, req)
	if e != nil {
		_ = r.Body.Close()
		return
	}

	var t1 interface{}
	_, e1, e2 := chain.GetTableRows2(req.Code, req.Scope, req.Code, req.Limit, req.UpperBound, req.LowerBound, &t1)
	//_, e1, e2 := chain.GetTableRows2("eosio", "producers", "eosio", 100, "", "", &t1)
	if e1 == nil && e2 == nil {
		data, e3 := json.Marshal(t1)
		if e3 == nil {
			_, _ = w.Write(data)
		}
	}
}

func abi(w http.ResponseWriter, r *http.Request) {
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

	res := &_type.AbiResponse{}
	res, e1, e2 := chain.GetAbi(req.Name)
	if e1 == nil && e2 == nil {
		data, e3 := json.Marshal(res)
		if e3 != nil {
			_ = r.Body.Close()
			return
		}

		_, _ = w.Write(data)
	}
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

	var result struct {
		Abi map[string]interface{} `json:"abi"`
	}
	e4 := mongodb.Accounts.FindOne(bson.M{"name": req.Name}).Decode(&result)

	contract := false
	if e1 == nil && e2 == nil && e4 == nil {
		if len(result.Abi) > 0 {
			contract = true
		}

		data, e3 := json.Marshal(struct {
			Contract bool                   `json:"contract"`
			Info     *_type.AccountResponse `json:"info"`
			Actions  *_type.ActionsResponse `json:"actions"`
		}{contract, res, res2})
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
