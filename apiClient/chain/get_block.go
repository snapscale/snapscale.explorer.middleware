package chain

import (
	"encoding/json"
	"snapscale-api/apiClient/common"
	_type "snapscale-api/apiClient/type"
)

func GetBlock(index string) (*_type.BlockResponse, *_type.Error, error) {
	var path = "get_block"

	reqT := &_type.BlockRequest{}
	reqT.BlockNumOrId = index

	reqB, _ := json.Marshal(reqT)

	data := &_type.BlockResponse{}

	e1, e2 := common.Do(reqB, base, path, data, false)

	if e1 == nil && e2 == nil {
		return data, e1, e2
	} else {
		return nil, e1, e2
	}
}
