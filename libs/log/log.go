package log

import (
	"log"
	"os"
	"snapscale-api/config"
)

var I *log.Logger
var E *log.Logger
var S *log.Logger

var ErrorFile *os.File
var InfoFile *os.File

//mix file output & stderr output
type ErrorIo struct{}

func (ErrorIo) Write(p []byte) (n int, err error) {
	_, _ = ErrorFile.Write(p)
	return os.Stderr.Write(p)
}

type InfoIo struct{}

func (InfoIo) Write(p []byte) (n int, err error) {
	_, _ = InfoFile.Write(p)
	return os.Stdout.Write(p)
}

func init() {
	ErrorFile, _ = os.OpenFile(config.ErrorLogPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	InfoFile, _ = os.OpenFile(config.InfoLogPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)

	I = log.New(&InfoIo{}, "", log.LstdFlags)
	E = log.New(&ErrorIo{}, "", log.LstdFlags|log.Llongfile)
	S = log.New(os.Stdout, "", log.LstdFlags)
}
