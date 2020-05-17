package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	startA := time.Now()
	http.Get(a)
	aDur := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDur := time.Since(startB)

	if aDur < bDur {
		return a
	}

	return b
}
