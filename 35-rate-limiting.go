package main

import "fmt"
import "time"

func main() {
	requests := make(chan int, 5)

	for i := 1; i < 5; i++ {
		requests <- i
	}

	close(requests)

	limiter := time.Tick(time.Millisecond * 500)

	for request := range requests {
		<-limiter
		fmt.Println("Received request", request, "at", time.Now())
	}

	burstLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 500) {
			burstLimiter <- t
		}
	}()

	burstRequests := make(chan int, 5)

	for i := 0; i < 5; i++ {
		burstRequests <- i
	}

	close(burstRequests)

	for request := range burstRequests {
		<-burstLimiter
		fmt.Println("Received (burst) request", request, "at", time.Now())
	}
}
