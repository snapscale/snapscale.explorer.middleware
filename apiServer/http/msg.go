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
