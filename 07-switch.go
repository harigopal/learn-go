package main

import "fmt"

func main() {
	fmt.Print("Enter 1, 2, 3, 4 or something else: ")

	var inputValue int
	fmt.Scanf("%d", &inputValue)

	switch inputValue {
	case 1, 2:
		fmt.Println("One or two.")
	case 3, 4:
		fmt.Println("Three or four.")
	default:
		fmt.Println("Something else.")
	}
}
