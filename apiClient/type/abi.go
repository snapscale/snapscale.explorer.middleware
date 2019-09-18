package _type

type AbiRequest struct {
	AccountName string `json:"account_name"`
}

type Abi struct {
	Version          string        `json:"version"`
	Types            []interface{} `json:"types"`
	Structs          []interface{} `json:"structs"`
	Actions          []interface{} `json:"actions"`
	Tables           []interface{} `json:"tables"`
	RicardianClauses []interface{} `json:"ricardian_clauses"`
	ErrorMessages    []interface{} `json:"error_messages"`
	AbiExtensions    []interface{} `json:"abi_extensions"`
	Variants         []interface{} `json:"variants"`
}

type AbiResponse struct {
	AccountName string `json:"account_name"`
	Abi         Abi    `json:"abi"`
}
