package main

import (
	"sort"
	"fmt"
)

func main() {
	words := []string{"world", "go", "hello"}
	sort.Strings(words)
	fmt.Println("Sorted words:", words)

	numbers := []int{42, 2, 4}
	sort.Ints(numbers)
	fmt.Println("Numbers:", numbers)

	areNumbersSorted := sort.IntsAreSorted(numbers)
	fmt.Println("Numbers Sorted?", areNumbersSorted)
}
