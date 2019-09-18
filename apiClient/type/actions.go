package _type

type ActionsRequest struct {
	AccountName string `json:"account_name"`
	Pos         int32  `json:"pos,omitempty"`
	Offset      int32  `json:"offset,omitempty"`
}

type Receipt struct {
	Receiver       string `json:"receiver"`
	ActDigest      string `json:"act_digest"`
	GlobalSequence int64  `json:"global_sequence"`
	RecvSequence   int64  `json:"recv_sequence"`
	CodeSequence   int64  `json:"code_sequence"`
	AbiSequence    int64  `json:"abi_sequence"`
}

type Authorization struct {
	Actor      string `json:"actor"`
	Permission string `json:"permission"`
}

type ActData struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Quantity string `json:"quantity"`
	Memo     string `json:"memo"`
}

type Act struct {
	Account       string          `json:"account"`
	Name          string          `json:"name"`
	Authorization []Authorization `json:"authorization"`
	Data          ActData         `json:"data"`
	HexData       string          `json:"hex_data"`
}

type ActionTrace struct {
	Receipt          Receipt       `json:"receipt"`
	Act              Act           `json:"act"`
	ContextFree      bool          `json:"context_free"`
	Elapsed          int64         `json:"elapsed"`
	Console          string        `json:"console"`
	TrxId            string        `json:"trx_id"`
	BlockNum         int64         `json:"block_num"`
	BlockTime        string        `json:"block_time"`
	ProducerBlockId  string        `json:"producer_block_id"`
	AccountRamDeltas interface{}   `json:"account_ram_deltas"`
	Except           interface{}   `json:"except"`
	InlineTraces     []ActionTrace `json:"inline_traces"`
}

type Action2 struct {
	GlobalActionSeq  int64       `json:"global_action_seq"`
	AccountActionSeq int64       `json:"account_action_seq"`
	BlockNum         int64       `json:"block_num"`
	BlockTime        string      `json:"block_time"`
	ActionTrace      ActionTrace `json:"action_trace"`
}

type ActionsResponse struct {
	Actions               []Action2 `json:"actions"`
	LastIrreversibleBlock int64     `json:"last_irreversible_block"`
}
