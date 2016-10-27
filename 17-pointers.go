package main

import "fmt"

func setZero(value int) {
	value = 0
}

func setZeroWithPointer(valuePointer *int) {
	*valuePointer = 0
}

func main() {
	number := 42
	fmt.Println("Intial value:", number)

	setZero(number)
	fmt.Println("After setZero:", number)

	setZeroWithPointer(&number)
	fmt.Println("After setZeroWithPointer:", number)

	fmt.Println("Pointer to number:", &number)
}
