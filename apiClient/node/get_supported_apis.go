package node

import (
	"snapscale-api/apiClient/common"
	_type "snapscale-api/apiClient/type"
)

func GetSupportedApis() (*_type.ApisResponse, *_type.Error, error) {
	var path = "get_supported_apis"
	data := &_type.ApisResponse{}
	e1, e2 := common.Do(nil, base, path, data, false)

	if e1 == nil && e2 == nil {
		return data, e1, e2
	} else {
		return nil, e1, e2
	}
}
