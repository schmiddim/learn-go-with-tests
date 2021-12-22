package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {

	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration > bDuration {

		return b
	}

	return a
}

func measureResponseTime(url string) (t time.Duration) {
	start := time.Now()
	http.Get(url)
	t = time.Since(start)
	return

}
