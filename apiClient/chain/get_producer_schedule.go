package chain

import (
	"snapscale-api/apiClient/common"
	_type "snapscale-api/apiClient/type"
)

func GetProducerSchedule() (*_type.ProducerScheduleResponse, *_type.Error, error) {
	var path = "get_producer_schedule"

	data := &_type.ProducerScheduleResponse{}

	e1, e2 := common.Do(nil, base, path, data, false)

	if e1 == nil && e2 == nil {
		return data, e1, e2
	} else {
		return nil, e1, e2
	}
}
