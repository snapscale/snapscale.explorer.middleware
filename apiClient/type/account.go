package _type

type AccountRequest struct {
	AccountName string `json:"account_name"`
}

type AccountResponse struct {
	AccountName       string `json:"account_name"`
	CoreLiquidBalance string `json:"core_liquid_balance"`
	HeadBlockNum      int64  `json:"head_block_num"`
	HeadBlockTime     string `json:"head_block_time"`
	Privileged        bool   `json:"privileged"`
	LastCodeUpdate    string `json:"last_code_update"`
	Created           string `json:"created"`
	RamQuota          int64  `json:"ram_quota"`
	NetWeight         int64  `json:"net_weight"`
	CpuWeight         int64  `json:"cpu_weight"`
	NetLimit          struct {
		Used      int64       `json:"used"`
		Available interface{} `json:"available"`
		Max       interface{} `json:"max"`
	} `json:"net_limit"`
	CpuLimit struct {
		Used      int64       `json:"used"`
		Available interface{} `json:"available"`
		Max       interface{} `json:"max"`
	} `json:"cpu_limit"`
	RamUsage    int64        `json:"ram_usage"`
	Permissions []Permission `json:"permissions"`
}

type Permission struct {
	PermName     string `json:"perm_name"`
	Parent       string `json:"parent"`
	RequiredAuth struct {
		Threshold int64 `json:"threshold"`
		Keys      []Key `json:"keys"`
	} `json:"required_auth"`
}

type Key struct {
	Key    string `json:"key"`
	Weight int64  `json:"weight"`
}
