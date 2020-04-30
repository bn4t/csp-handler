package main

import "time"

// rateLimitLoop resets the rate limit every hour
func rateLimitLoop() {
	for {
		time.Sleep(1 * time.Hour)
		RateLimitMap = make(map[string]int)
	}
}
