package main

import "fmt"

func multipleValues() (int, int) {
	return 42, 21
}

func main() {
	firstReturn, secondReturn := multipleValues()
	fmt.Println(firstReturn, secondReturn)

	_, oneSkipped := multipleValues()
	fmt.Println(oneSkipped)
}
