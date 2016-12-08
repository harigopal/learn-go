package main

import "fmt"

func main() {
	messageChannel := make(chan string, 2)

	messageChannel <- "First string"
	messageChannel <- "Second string"

	fmt.Println(<-messageChannel)
	fmt.Println(<-messageChannel)
}
