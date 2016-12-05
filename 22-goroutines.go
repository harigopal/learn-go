package main

import "fmt"

func printSomething(source string) {
	for i := 0; i < 4; i++ {
		fmt.Println(source, ":", i)
	}
}

func main() {
	printSomething("main")

	go printSomething("go-routine")
	go printSomething("go-routine-2")

	go func(message string) {
		fmt.Println("Message from anonymous Goroutine:", message)
	}("Foo!")

	var input string
	fmt.Scanln(&input)
	fmt.Println("We're done? Already?")
}
