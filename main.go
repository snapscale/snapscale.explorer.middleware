package main

import (
	"snapscale-api/apiServer/http"
	"snapscale-api/apiServer/ws"
	"snapscale-api/dashBoard"
	"snapscale-api/database/mongodb"
)

func main() {
	dashBoard.Start()
	mongodb.Start()
	go http.Start()
	ws.Start()
}
