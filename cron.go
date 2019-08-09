package main

import (
	"github.com/robfig/cron"
)

func initCron() {
	c := cron.New()
	// wipe the ratelimit map every hour
	c.AddFunc("@hourly", func() { RateLimitMap = make(map[string]int) })
	c.Start()
}
