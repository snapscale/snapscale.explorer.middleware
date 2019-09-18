package chain

import (
	"encoding/json"
	"snapscale-api/apiClient/common"
	_type "snapscale-api/apiClient/type"
)

func GetTableRows(code string, table string, scope string, data interface{}) (interface{}, *_type.Error, error) {
	var path = "get_table_rows"

	reqT := &_type.TableRowsRequest{}
	reqT.Code = code
	reqT.Table = table
	reqT.Scope = scope
	reqT.JSON = true

	reqB, _ := json.Marshal(reqT)

	e1, e2 := common.Do(reqB, base, path, data, false)

	if e1 == nil && e2 == nil {
		return data, e1, e2
	} else {
		return nil, e1, e2
	}
}

func GetTableRows2(code string, table string, scope string, limit int32, upper_bound string, lower_bound string, data interface{}) (interface{}, *_type.Error, error) {
	var path = "get_table_rows"

	reqT := &_type.TableRowsRequest{}
	reqT.Code = code
	reqT.Table = table
	reqT.Scope = scope
	reqT.JSON = true
	reqT.Limit = limit
	reqT.UpperBound = upper_bound
	reqT.LowerBound = lower_bound

	reqB, _ := json.Marshal(reqT)

	e1, e2 := common.Do(reqB, base, path, data, false)

	if e1 == nil && e2 == nil {
		return data, e1, e2
	} else {
		return nil, e1, e2
	}
}
