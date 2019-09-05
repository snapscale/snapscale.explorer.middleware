package cron

import (
	"time"
)

func New(d time.Duration, f func()) {
	go func() {
		f()
		for range time.Tick(d) {
			f()
		}
	}()
}
