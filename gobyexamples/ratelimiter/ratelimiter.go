package ratelimiter

import (
	"fmt"
	"time"
)

func SimpleLimiter() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.NewTicker(2 * time.Second)

	for req := range requests {
		<-limiter.C
		fmt.Println("request", req, time.Now())
	}
}

func BurstyLimiter() {
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(2 * time.Second) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
