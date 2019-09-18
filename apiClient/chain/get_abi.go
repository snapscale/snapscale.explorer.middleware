package chain

import (
	"encoding/json"
	"snapscale-api/apiClient/common"
	_type "snapscale-api/apiClient/type"
)

func GetAbi(name string) (*_type.AbiResponse, *_type.Error, error) {
	var path = "get_abi"

	reqT := &_type.AccountRequest{}
	reqT.AccountName = name

	reqB, _ := json.Marshal(reqT)

	data := &_type.AbiResponse{}

	e1, e2 := common.Do(reqB, base, path, data, false)

	if e1 == nil && e2 == nil {
		return data, e1, e2
	} else {
		return nil, e1, e2
	}
}
