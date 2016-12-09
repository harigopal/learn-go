package main

import "fmt"
import "time"

func main() {
	firstChannel := make(chan string)
	secondChannel := make(chan string)

	go func() {
		time.Sleep(time.Second)
		firstChannel <- "I am number one."
	}()

	go func() {
		time.Sleep(time.Second * 2)
		secondChannel <- "I am number two. Ew."
	}()

	for i := 0; i < 2; i++ {
		select {
		case messageFromFirstChannel := <-firstChannel:
			fmt.Println("First Channel Message:", messageFromFirstChannel)
		case messageFromSecondChannel := <-secondChannel:
			fmt.Println("Second Channel Message:", messageFromSecondChannel)
		}
	}
}
