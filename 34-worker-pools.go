package main

import "fmt"
import "time"

func pointlessWorker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("Worker", id, "is will 'work' on", job)
		time.Sleep(time.Millisecond * 250)
		fmt.Println("OK. Worker", id, "is done with", job)
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for worker := 1; worker < 3; worker++ {
		go pointlessWorker(worker, jobs, results)
	}

	for job := 1; job < 10; job++ {
		jobs <- job
	}

	close(jobs)

	for result := 1; result < 10; result++ {
		<-results
	}
}
