package history

import (
	"encoding/json"
	"snapscale-api/apiClient/common"
	_type "snapscale-api/apiClient/type"
)

func GetActions(name string) (*_type.ActionsResponse, *_type.Error, error) {
	var path = "get_actions"

	reqT := &_type.ActionsRequest{}
	reqT.AccountName = name

	reqB, _ := json.Marshal(reqT)

	data := &_type.ActionsResponse{}

	e1, e2 := common.Do(reqB, base, path, data, false)

	if e1 == nil && e2 == nil {
		return data, e1, e2
	} else {
		return nil, e1, e2
	}
}
