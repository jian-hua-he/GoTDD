package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	aDur := measureResponseTime(a)
	bDur := measureResponseTime(b)

	if aDur < bDur {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
