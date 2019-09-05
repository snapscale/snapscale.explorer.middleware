package _type

type TransactionRequest struct {
	Id           string `json:"id"`
	BlockNumHint int32  `json:"block_num_hint"`
}

type TransactionResponse struct {
	Id                    string        `json:"id"`
	BlockTime             string        `json:"block_time"`
	BlockNum              int64         `json:"block_num"`
	LastIrreversibleBlock int64         `json:"last_irreversible_block"`
	Traces                []interface{} `json:"traces"`
	Trx                   struct {
		Receipt struct {
			Status        string      `json:"status"`
			CpuUsageUs    int64       `json:"cpu_usage_us"`
			NetUsageWords int64       `json:"net_usage_words"`
			Trx           interface{} `json:"trx"`
		} `json:"receipt"`
		Trx struct {
			Expiration         string        `json:"expiration"`
			RefBlockNum        int64         `json:"ref_block_num"`
			RefBlockPrefix     int64         `json:"ref_block_prefix"`
			MaxNetUsageWords   int64         `json:"max_net_usage_words"`
			MaxCpuUsageMs      int64         `json:"max_cpu_usage_ms"`
			DelaySec           int64         `json:"delay_sec"`
			ContextFreeActions []interface{} `json:"context_free_actions"`
			Actions            []struct {
				Account       string `json:"account"`
				Name          string `json:"name"`
				Authorization []struct {
					Actor      string `json:"actor"`
					Permission string `json:"permission"`
				} `json:"authorization"`
				Data struct {
					From     string `json:"from"`
					To       string `json:"to"`
					Quantity string `json:"quantity"`
					Memo     string `json:"memo"`
				} `json:"data"`
				HexData string `json:"hex_data"`
			} `json:"actions"`
			TransactionExtensions []interface{} `json:"transaction_extensions"`
			ContextFreeData       []interface{} `json:"context_free_data"`
			Signatures            []string      `json:"signatures"`
		} `json:"trx"`
	} `json:"trx"`
}
