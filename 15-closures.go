package main

import "fmt"

func createCounter() func() int {
	counter := 0

	return func() int {
		counter += 1
		return counter
	}
}

func main() {
	newCounter := createCounter()

	fmt.Println(newCounter())
	fmt.Println(newCounter())
	fmt.Println(newCounter())

	anotherCounter := createCounter()

	fmt.Println(anotherCounter())
	fmt.Println(newCounter())
}
