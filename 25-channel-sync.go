package main

import "fmt"
import "time"

func processSomething(done chan bool) {
	fmt.Println("I'm processing something...")
	time.Sleep(time.Second)
	fmt.Println("All right, I'm done.")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go processSomething(done)
	fmt.Println("main() has invoked processing.")
	<-done
	fmt.Println("main() has confirmed completion of processing.")
}
