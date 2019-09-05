package common

import (
	"encoding/json"
	"errors"
	_type "snapscale-api/apiClient/type"
	"snapscale-api/libs/errorMsg"
	"snapscale-api/libs/log"
	"snapscale-api/libs/request"
	"snapscale-api/tools"
)

func Do(data []byte, base string, path string, in interface{}, logFlag bool) (*_type.Error, error) {
	body, err := request.Post(base+path, data)

	if err != nil {
		return nil, err
	}

	jsonTag := json.Valid(body)
	if !jsonTag {
		return nil, errors.New(errorMsg.ResNotJSON)
	}

	if logFlag {
		log.I.Println(tools.JsonDump(body))
	}

	//false
	f := &_type.Error{}
	err = json.Unmarshal(body, f)

	if err != nil {
		return nil, err
	}

	if f.Code != 0 {
		return f, nil
	}

	//true
	err = json.Unmarshal(body, in)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
