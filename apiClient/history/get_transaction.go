package history

import (
	"encoding/json"
	"snapscale-api/apiClient/common"
	_type "snapscale-api/apiClient/type"
)

func GetTransaction(index string, block_num int32) (*_type.TransactionResponse, *_type.Error, error) {
	var path = "get_transaction"

	reqT := &_type.TransactionRequest{}
	reqT.Id = index
	reqT.BlockNumHint = block_num

	reqB, _ := json.Marshal(reqT)

	data := &_type.TransactionResponse{}

	e1, e2 := common.Do(reqB, base, path, data, false)

	if e1 == nil && e2 == nil {
		return data, e1, e2
	} else {
		return nil, e1, e2
	}
}
