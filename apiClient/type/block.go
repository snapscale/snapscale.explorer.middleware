package _type

type BlockRequest struct {
	BlockNumOrId string `json:"block_num_or_id"`
}

type Auth struct {
	Actor      string `json:"actor"`
	Permission string `json:"permission"`
}

type ActionData struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Quantity string `json:"quantity"`
	Memo     string `json:"memo"`
}

type Action struct {
	Account       string     `json:"account"`
	Name          string     `json:"name"`
	Authorization []Auth     `json:"authorization"`
	Data          ActionData `json:"data"`
	HexData       string     `json:"hex_data"`
}

type Tx struct {
	Expiration       string   `json:"expiration"`
	RefBlockNum      int64    `json:"ref_block_num"`
	RefBlockPrefix   int64    `json:"ref_block_prefix"`
	DelaySec         int64    `json:"delay_sec"`
	Actions          []Action `json:"actions"`
	MaxNetUsageWords int64    `json:"max_net_usage_words"`
	MaxCpuUsageMs    int64    `json:"max_cpu_usage_ms"`
}

type Trx struct {
	Id          string   `json:"id"`
	Signatures  []string `json:"signatures"`
	Compression string   `json:"compression"`
	PackedTrx   string   `json:"packed_trx"`
	Transaction Tx       `json:"transaction"`
}

type Transaction struct {
	Status        string `json:"status"`
	Trx           Trx    `json:"trx"`
	CpuUsageUs    int64  `json:"cpu_usage_us"`
	NetUsageWords int64  `json:"net_usage_words"`
}

type BlockResponse struct {
	TimeStamp         string        `json:"timestamp"`
	Producer          string        `json:"producer"`
	Confirmed         int64         `json:"confirmed"`
	Previous          string        `json:"previous"`
	TransactionMroot  string        `json:"transaction_mroot"`
	ActionMroot       string        `json:"action_mroot"`
	ScheduleVersion   int64         `json:"schedule_version"`
	ProducerSignature string        `json:"producer_signature"`
	Id                string        `json:"id"`
	BlockNum          int64         `json:"block_num"`
	RefBlockPrefix    int64         `json:"ref_block_prefix"`
	Transactions      []Transaction `json:"transactions"`
}
