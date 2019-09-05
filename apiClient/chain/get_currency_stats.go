package chain

import (
	"encoding/json"
	"snapscale-api/apiClient/common"
	_type "snapscale-api/apiClient/type"
)

func GetCurrencyStats(code string, symbol string) (*_type.CurrencyStatsResponse, *_type.Error, error) {
	var path = "get_currency_stats"

	reqT := &_type.CurrencyStatsRequest{}
	reqT.Code = code
	reqT.Symbol = symbol

	reqB, _ := json.Marshal(reqT)

	data := &_type.CurrencyStatsResponse{}

	e1, e2 := common.Do(reqB, base, path, data, false)

	if e1 == nil && e2 == nil {
		return data, e1, e2
	} else {
		return nil, e1, e2
	}
}
