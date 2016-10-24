package main

import "fmt"

func main() {
	sliceString := make([]string, 3)
	fmt.Println("Empty?", sliceString)
	fmt.Println("Length:", len(sliceString))
	sliceString[2] = "f"
	fmt.Println("With value:", sliceString)

	sliceString = append(sliceString, "g", "h")
	fmt.Println("Appened:", sliceString)
	fmt.Println("Length:", len(sliceString))

	copiedString := make([]string, len(sliceString))
	copy(copiedString, sliceString)
	fmt.Println("Copied:", copiedString)

	sliceOne := sliceString[2:5]
	fmt.Println("Slice One:", sliceOne)

	prepopulated := []string{"h", "e", "l", "l", "o"}
	fmt.Println("Prepopulated:", prepopulated)

	twoDimensionalSlice := make([][]int, 3)

	for x := 0; x < 3; x++ {
		yMax := x * 2
		twoDimensionalSlice[x] = make([]int, yMax)

		for y := 0; y < yMax; y++ {
			twoDimensionalSlice[x][y] = x + y
		}
	}

	fmt.Println("Two Dimensional Slice:", twoDimensionalSlice)
}
