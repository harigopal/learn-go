package main

import "fmt"

func main() {
	firstEverMap := make(map[string]int)

	firstEverMap["keyOne"] = 1
	firstEverMap["keyTwo"] = 2

	fmt.Println(firstEverMap)
	fmt.Println("Length of map:", len(firstEverMap))

	firstEverMap["keyThree"] = 3
	delete(firstEverMap, "keyTwo")
	fmt.Println(firstEverMap)

	_, keyPresent := firstEverMap["keyTwo"]
	fmt.Println("Key Present?", keyPresent)

	secondMapEver := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(secondMapEver)
}
