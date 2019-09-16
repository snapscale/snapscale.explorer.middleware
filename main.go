package main

import (
	"snapscale-api/apiServer/http"
	"snapscale-api/apiServer/ws"
	"snapscale-api/dashBoard"
	"snapscale-api/database/mongodb"
	"snapscale-api/libs/log"
)

func main() {
	log.S.Println("Snapscale explorer middleware start")
	dashBoard.Start()
	mongodb.Start()
	go http.Start()
	ws.Start()
}
