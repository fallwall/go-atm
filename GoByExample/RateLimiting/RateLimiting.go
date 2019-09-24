package main

import (
	"fmt"
	"time"
)

func main() {

	// limit our handling of requests channel (incoming)
	requests := make(chan int, 5)
	for i :=1; i <=5; i++ {
		requests <- i
	}
	close(requests)

	// limiter channel to receive a value every 200ms
	// regulator in out rate limiting scheme
	limiter := time.Tick(200 * time.Millisecond)

	// blocking on a receive from the limiter before serving each request
	// 1 req/ 200ms
	for req := range requests {
		<-limiter
		fmt.Println("requests", req, time.Now())
	}

	// to preserbe the overall limite rate
	// by allowing short bursts of requests in out rate limiting scheme
	burstyLimiter := make(chan time.Time, 3)

	// fill up the channel to represnet allowed bursting
	for i := 0; i< 3; i++ {
		burstyLimiter <- time.Now()
	}

	// everey 200ms,add a new value to burstylimiter
	// up to its limit of 3
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// simulate 5 more incoming reqs
	burstyRequests := make(chan int, 5)
	for i :=1; i<=5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
