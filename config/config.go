package config

import (
	"os"
	"path"
	"snapscale-api/libs/fileDir"
)

//var ApiBase = "http://192.168.1.201:30132/v1/"
//var ApiBase = "http://34.80.167.80:8888/v1/"
var ApiBase = os.Getenv("ApiBase")

//port
//const HttpPort = ":8090"
//const WsPort = ":8089"
var HttpPort = ":" + os.Getenv("HttpPort")
var WsPort = ":" + os.Getenv("WsPort")

//db
//const MongoConfig = "mongodb://xeniro:N0password@192.168.1.201:30017/?authSource=admin"
var MongoConfig = os.Getenv("MongoConfig")

//var
var FileDictionary = fileDir.ExecuteDirectory()
var LogDirBase = os.Getenv("SSE_LOG_PATH")
var LogDir string
var ErrorLogPath string
var InfoLogPath string

func init() {
	//logs
	if LogDirBase == "" {
		LogDirBase = FileDictionary
	}

	LogDir = path.Join(LogDirBase, "logs")
	InfoLogPath = path.Join(LogDir, "info.log")
	ErrorLogPath = path.Join(LogDir, "error.log")
}
