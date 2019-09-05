package dashBoard

import (
	_type "snapscale-api/apiClient/type"
	"snapscale-api/libs/cron"
	"time"
)

var DataCenter = new(DataCenterS)
var BroadcastDashboard func()
var BroadcastBlock func(interface{})

func init() {
	DataCenter.Producers.ProducerMap = make(map[string]_type.ProducerInfo)
	DataCenter.Producers.ProducerLoop = make(map[string]string)
	cron.New(time.Minute, longSync)
	cron.New(time.Second, shortSync)
	SyncNow = make(chan int64)
	SyncFinish = make(chan bool)
	tpsaps = make(chan *tpsapsS)

	go blockSync()
	go tpsapsSync()
}

func Start() {}
