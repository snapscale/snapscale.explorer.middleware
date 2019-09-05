package chain

import (
	"snapscale-api/apiClient/common"
	_type "snapscale-api/apiClient/type"
)

func GetInfo() (*_type.InfoResponse, *_type.Error, error) {
	var path = "get_info"

	data := &_type.InfoResponse{}

	e1, e2 := common.Do(nil, base, path, data, false)

	if e1 == nil && e2 == nil {
		return data, e1, e2
	} else {
		return nil, e1, e2
	}
}
