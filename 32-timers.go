package main

import "fmt"
import "time"

func main() {
	twoSecondTimer := time.NewTimer(time.Second * 2)

	<-twoSecondTimer.C
	fmt.Println("Two seconds are up!")

	preemptedTimer := time.NewTimer(time.Second)

	go func() {
		<-preemptedTimer.C
		fmt.Println("Will I? Won't I? :-/")
	}()

	stopTimer := preemptedTimer.Stop()

	if stopTimer {
		fmt.Println("preemptedTimer has indeed been preempted.")
	}
}
