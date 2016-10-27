package main

import "fmt"

func main() {
	numbers := []int{42, 21, 10, 5, 2, 1}

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	fmt.Println("Sum:", sum)

	for index, number := range numbers {
		if number == 10 {
			fmt.Println("Index of 10 is", index)
		}
	}

	fruits := map[string]string{"a": "apple", "b": "banana", "c": "citrus"}

	for key, value := range fruits {
		fmt.Println(key, "for", value)
	}

	for index, character := range "rendezvous" {
		fmt.Printf("%d: %c\n", index, character)
	}
}
