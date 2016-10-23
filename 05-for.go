package main

import "fmt"

func main() {
	variableOutside := 1
	for variableOutside <= 3 {
		fmt.Println(variableOutside)
		variableOutside += 1
	}

	for inlineVariable := 39; inlineVariable <= 42; inlineVariable++ {
		fmt.Println(inlineVariable)
	}

	anotherVariableOutside := 0

	for {
		fmt.Println("loop the loop")

		if anotherVariableOutside == 2 {
			break
		}

		anotherVariableOutside++
	}
}
