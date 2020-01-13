package dashBoard

import (
	"snapscale-api/apiClient/chain"
	_type "snapscale-api/apiClient/type"
	"snapscale-api/database/mongodb"
	"strconv"
	"strings"
)

func longSync() {
	syncCurrencyStats()
	syncProducers()
	syncGlobal()
	syncStakedTotal()
	syncTransactions()
	syncUsers()
	syncContracts()
}

func syncCurrencyStats() {
	t, e1, e2 := chain.GetCurrencyStats("eosio.token", "XST")
	if e1 == nil && e2 == nil && t != nil {
		DataCenter.XST.TotalSupplyOfXST, _ = strconv.ParseInt(strings.Split(t.XST.Supply, ".")[0], 10, 64)
	}
}

func syncProducers() {
	t1 := &_type.XProducerResponse{}
	_, e1, e2 := chain.GetTableRows("eosio", "producers", "eosio", t1)
	if e1 == nil && e2 == nil && len(t1.Rows) > 0 {
		for _, one := range t1.Rows {
			DataCenter.Producers.ProducerMap[one.ProducerKey] = one
		}
	}

	t2, e1, e2 := chain.GetProducerSchedule()
	if e1 == nil && e2 == nil && t2 != nil {
		DataCenter.Producers.ActiveProducerList = t2.Active.Producers
		for index, one := range t2.Active.Producers {
			if (index + 1) == len(t2.Active.Producers) {
				DataCenter.Producers.ProducerLoop[one.ProducerName] = t2.Active.Producers[0].ProducerName
			} else {
				DataCenter.Producers.ProducerLoop[one.ProducerName] = t2.Active.Producers[index+1].ProducerName
			}
		}
	}
}

func syncGlobal() {
	t1 := &_type.XGlobalResponse{}
	_, e1, e2 := chain.GetTableRows("eosio", "global", "eosio", t1)
	if e1 == nil && e2 == nil && len(t1.Rows) > 0 {
		DataCenter.IO.RamInChain, _ = strconv.ParseInt(t1.Rows[0].MaxRamSize, 10, 64)
		DataCenter.XST.VotedXST, _ = strconv.ParseInt(t1.Rows[0].TotalActivatedStake, 10, 64)
		DataCenter.XST.VotedXST = DataCenter.XST.VotedXST / 10000
		DataCenter.Chain.TotalVoteWeight = t1.Rows[0].TotalProducerVoteWeight
	}
}

func syncStakedTotal() {
	t, e1, e2 := chain.GetAccount("eosio.stake")
	if e1 == nil && e2 == nil && t != nil {
		DataCenter.XST.StakedXST, _ = strconv.ParseInt(strings.Split(t.CoreLiquidBalance, ".")[0], 10, 64)
	}
}

func syncTransactions() {
	DataCenter.Count.Transactions = mongodb.Transactions.Count()
}

func syncUsers() {
	DataCenter.Count.Users = mongodb.Accounts.Count()
}

func syncContracts() {
	t1 := &_type.XAbiHashResponse{}
	_, _, _ = chain.GetTableRows("eosio", "abihash", "eosio", t1)

	length := len(t1.Rows)
	DataCenter.Count.Contracts = int64(length)
}
