package ws

type Msg struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}
