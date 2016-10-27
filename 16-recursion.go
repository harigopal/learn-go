package main

import "fmt"

func factorial(number int) int {
	if number == 0 {
		return 1
	}

	return number * factorial(number - 1)
}

func main() {
	fmt.Println("5 factorial =", factorial(5))
}
