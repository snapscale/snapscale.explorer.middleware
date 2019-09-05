package _type

type InfoResponse struct {
	ServerVersion            string `json:"server_version"`
	ChainId                  string `json:"chain_id"`
	HeadBlockNum             int64  `json:"head_block_num"`
	LastIrreversibleBlockNum int64  `json:"last_irreversible_block_num"`
	LastIrreversibleBlockId  string `json:"last_irreversible_block_id"`
	HeadBlockId              string `json:"head_block_id"`
	HeadBlockTime            string `json:"head_block_time"`
	HeadBlockProducer        string `json:"head_block_producer"`
	VirtualBlockCpuLimit     int64  `json:"virtual_block_cpu_limit"`
	VirtualBlockNetLimit     int64  `json:"virtual_block_net_limit"`
	BlockCpuLimit            int64  `json:"block_cpu_limit"`
	BlockNetLimit            int64  `json:"block_net_limit"`
	ServerVersionString      string `json:"server_version_string"`
}
