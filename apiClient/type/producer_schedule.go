package _type

type Producer struct {
	BlockSigningKey string `json:"block_signing_key"`
	ProducerName    string `json:"producer_name"`
}

type Active struct {
	Producers []Producer `json:"producers"`
	Version   int64      `json:"version"`
}

type ProducerScheduleResponse struct {
	Active Active `json:"active"`
}
