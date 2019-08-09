package main

import (
	"github.com/robfig/cron"
	"log"
)

func initCron() {
	c := cron.New()

	// wipe the ratelimit map every hour
	err := c.AddFunc("@hourly", func() { RateLimitMap = make(map[string]int) })
	if err != nil {
		log.Fatal(err)
	}

	c.Start()
}
