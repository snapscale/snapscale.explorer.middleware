package http

type accountRequest struct {
	Name string `json:"name"`
}

type blockRequest struct {
	Num int64 `json:"num"`
}

type transactionRequest struct {
	Hash string `json:"hash"`
}

type contractRequest struct {
	Code       string `json:"code"`
	Scope      string `json:"scope"`
	Limit      int32  `json:"limit"`
	LowerBound string `json:"lower_bound,omitempty"`
	UpperBound string `json:"upper_bound,omitempty"`
}

type actionsRequest struct {
	AccountName string `json:"account_name"`
	Pos         int32  `json:"pos"`
	Offset      int32  `json:"offset"`
}
