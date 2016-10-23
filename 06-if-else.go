package main

import "time"
import "fmt"
import "math/rand"

func main() {
	seed := time.Now().UnixNano()
	println("Seed is", seed)

	source := rand.NewSource(seed)
	newRand := rand.New(source)

	randomNumber := newRand.Intn(100)

	if randomNumber < 50 {
		fmt.Println("randomNumber is less than 50:", randomNumber)
	} else if randomNumber > 75 {
		fmt.Println("randomNumber is greater than 75:", randomNumber)
	} else {
		fmt.Println("randomNumber is between 50 and 75:", randomNumber)
	}
}
