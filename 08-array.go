package main

import "fmt"

func main() {
	var someNumbers [5]int

	fmt.Println("What is this?", someNumbers)

	someNumbers[4] = 42
	fmt.Println("Modified!", someNumbers)
	fmt.Println("Last one is", someNumbers[4])
	fmt.Println("Length is", len(someNumbers))

	filledArray := [5]int{42, 21, 10, 5, 2}

	var twoDimensional [5][3]int

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			twoDimensional[i][j] = filledArray[i] * j
		}
	}

	fmt.Println(twoDimensional)
}
