package _type

type TableRowsRequest struct {
	Code  string `json:"code"`
	Table string `json:"table"`
	Scope string `json:"scope"`
	JSON  bool   `json:"json"`
}

type ProducerInfo struct {
	Owner         string `json:"owner"`
	TotalVotes    string `json:"total_votes"`
	ProducerKey   string `json:"producer_key"`
	IsActive      int64  `json:"is_active"`
	Url           string `json:"url"`
	UnpaidBlocks  int64  `json:"unpaid_blocks"`
	LastClaimTime string `json:"last_claim_time"`
	Location      int64  `json:"location"`
}

type GlobalInfo struct {
	MaxBlockNetUsage               int64  `json:"max_block_net_usage"`
	TargetBlockNetUsagePct         int64  `json:"target_block_net_usage_pct"`
	MaxTransactionNetUsage         int64  `json:"max_transaction_net_usage"`
	BasePerTransactionNetUsage     int64  `json:"base_per_transaction_net_usage"`
	NetUsageLeeway                 int64  `json:"net_usage_leeway"`
	ContextFreeDiscountNetUsageNum int64  `json:"context_free_discount_net_usage_num"`
	ContextFreeDiscountNetUsageDen int64  `json:"context_free_discount_net_usage_den"`
	MaxBlockCpuUsage               int64  `json:"max_block_cpu_usage"`
	TargetBlockCpuUsagePct         int64  `json:"target_block_cpu_usage_pct"`
	MaxTransactionCpuUsage         int64  `json:"max_transaction_cpu_usage"`
	MinTransactionCpuUsage         int64  `json:"min_transaction_cpu_usage"`
	MaxTransactionLifetime         int64  `json:"max_transaction_lifetime"`
	DeferredTrxExpirationWindow    int64  `json:"deferred_trx_expiration_window"`
	MaxTransactionDelay            int64  `json:"max_transaction_delay"`
	MaxInlineActionSize            int64  `json:"max_inline_action_size"`
	MaxInlineActionDepth           int64  `json:"max_inline_action_depth"`
	MaxAuthorityDepth              int64  `json:"max_authority_depth"`
	MaxRamSize                     string `json:"max_ram_size"`
	TotalRamBytesReserved          int64  `json:"total_ram_bytes_reserved"`
	TotalRamStake                  int64  `json:"total_ram_stake"`
	LastProducerScheduleUpdate     string `json:"last_producer_schedule_update"`
	LastPervoteBucketFill          string `json:"last_pervote_bucket_fill"`
	PervoteBucket                  int64  `json:"pervote_bucket"`
	PerblockBucket                 int64  `json:"perblock_bucket"`
	TotalUnpaidBlocks              int64  `json:"total_unpaid_blocks"`
	TotalActivatedStake            string `json:"total_activated_stake"`
	ThreshActivatedStakeTime       string `json:"thresh_activated_stake_time"`
	LastProducerScheduleSize       int64  `json:"last_producer_schedule_size"`
	TotalProducerVoteWeight        string `json:"total_producer_vote_weight"`
	LastNameClose                  string `json:"last_name_close"`
}

type AbiHashInfo struct {
	Owner string `json:"owner"`
	Hash  string `json:"hash"`
}

type XGlobalResponse struct {
	Rows []GlobalInfo `json:"rows"`
}

type XProducerResponse struct {
	Rows []ProducerInfo `json:"rows"`
}

type XAbiHashResponse struct {
	Rows []AbiHashInfo `json:"rows"`
}
