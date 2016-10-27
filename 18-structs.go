package main

import "fmt"

type animal struct {
	name string
	legs int
}

func main() {
	fmt.Println("Human:", animal{"Human", 2})
	fmt.Println("Cat:", animal{name: "Cat", legs: 4})
	fmt.Println("Dog:", animal{name: "Dog"})

	millipede := animal{"Millipede", 1000}
	fmt.Println("Name:", millipede.name)

	pointerToMillipede := &millipede
	fmt.Println("Legs:", pointerToMillipede.legs)

	millipede.legs = 750
	fmt.Println("Actually, it maxes out at:", pointerToMillipede.legs)
}
