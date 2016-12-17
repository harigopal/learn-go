package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case message := <-messages:
		fmt.Println("Received message:", message)
	default:
		fmt.Println("Didn't get anything")
	}

	message := "Hello there!"

	select {
	case messages <- message:
		fmt.Println("Sent message: ", message)
	default:
		fmt.Println("Didn't send anything")
	}

	select {
	case message := <-messages:
		fmt.Println("Recieved message: ", message)
	case signal := <-signals:
		fmt.Println("Received signal: ", signal)
	default:
		fmt.Println("No activity, commander.")
	}
}
