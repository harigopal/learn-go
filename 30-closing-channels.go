package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			job, more := <-jobs

			if more {
				fmt.Println("Received job: ", job)
			} else {
				fmt.Println("Received all jobs")
				done <- true
				return
			}
		}
	}()

	for count := 0; count <= 2; count++ {
		jobs <- count
		fmt.Println("Dispatched job: ", count)
	}

	close(jobs)
	fmt.Println("Dispatched all jobs, and closed jobs channel")

	<-done
	close(done)
}
