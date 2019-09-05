package dashBoard

import _type "snapscale-api/apiClient/type"

type DataCenterS struct {
	Chain struct {
		TotalVoteWeight string `json:"total_vote_weight"`
	} `json:"chain"`
	Blocks struct {
		CurrentBlock          int64 `json:"current_block"`
		LastIrreversibleBlock int64 `json:"last_irreversible_block"`
	} `json:"blocks"`
	Producers struct {
		CurrentProducer    string                        `json:"current_producer"`
		NextProducer       string                        `json:"next_producer"`
		ActiveProducerList []_type.Producer              `json:"active_producer_list"`
		ProducerMap        map[string]_type.ProducerInfo `json:"producer_map"`
		ProducerLoop       map[string]string             `json:"producer_loop"`
	} `json:"producers"`
	Performance struct {
		Tps     int64 `json:"tps"`
		Aps     int64 `json:"aps"`
		TpsHigh int64 `json:"tps_high"`
		ApsHigh int64 `json:"aps_high"`
	} `json:"performance"`
	XST struct {
		VotedXST         int64 `json:"voted_xst"`
		StakedXST        int64 `json:"staked_xst"`
		TotalSupplyOfXST int64 `json:"total_supply_of_xst"`
	} `json:"xst"`
	IO struct {
		RamInChain           int64 `json:"ram_in_chain"`
		VirtualBlockCpuLimit int64 `json:"virtual_block_cpu_limit"`
		VirtualBlockNetLimit int64 `json:"virtual_block_net_limit"`
		BlockCpuLimit        int64 `json:"block_cpu_limit"`
		BlockNetLimit        int64 `json:"block_net_limit"`
	} `json:"io"`
	Count struct {
		VotedTotal  int64 `json:"voted_total"`
		StakedTotal int64 `json:"staked_total"`
	} `json:"count"`
}
