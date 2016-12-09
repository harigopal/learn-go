package main

import "fmt"

func accept(outgoingChannel chan<- string, message string) {
	outgoingChannel <- message
}

func transfer(incomingChannel <-chan string, outgoingChannel chan<- string) {
	message := <-incomingChannel
	outgoingChannel <- message
}

func main() {
	receiver := make(chan string, 1)
	sender := make(chan string, 1)

	accept(receiver, "Houston, we have a problem.")
	transfer(receiver, sender)

	fmt.Println(<-sender)
}
