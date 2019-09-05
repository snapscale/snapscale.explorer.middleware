package http

import (
	"log"
	"net/http"
	"snapscale-api/config"
)

func Start() {
	server := &http.Server{}
	server.Addr = config.HttpPort
	server.Handler = Router
	log.Fatal(server.ListenAndServe())
}
