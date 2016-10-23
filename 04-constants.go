package main

import "fmt"
import "math"

const helloWorld string = "Hello world!"

func main() {
	fmt.Println(helloWorld)

	const theAnswer = 42

	const bigNumber = 42e10 / theAnswer
	fmt.Println(bigNumber)

	fmt.Println(int64(bigNumber))

	fmt.Println(math.Sin(theAnswer))
}
