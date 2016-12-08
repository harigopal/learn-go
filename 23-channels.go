package main

import "fmt"

func main() {
	messageChannel := make(chan string)

	go func() {
		messageChannel <- "Hello there!"
	}()

	message := <-messageChannel

	fmt.Println(message)
}
