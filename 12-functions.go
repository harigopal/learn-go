package main

import "fmt"

func sum(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}

func sumThree(firstNumber, secondNumber, thirdNumber int) int {
	return firstNumber + secondNumber + thirdNumber
}

func main() {
	fmt.Println("Sum of 42 + 21:", sum(42, 21))
	fmt.Println("Sum of 42 + 21 + 10:", sumThree(42, 21, 10))
}
