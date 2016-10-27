package main

import "fmt"

func sum(numbers ...int) int {
	finalSum := 0

	for _, number := range numbers {
		finalSum += number
	}

	return finalSum
}

func main() {
	fmt.Println("Sum of 42, 21, 10, 5:", sum(42, 21, 10, 5))

	prepopulatedList := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Sum of prepopulated list:", sum(prepopulatedList...))
}
