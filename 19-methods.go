package main

import "fmt"
import "math"

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	exampleCircle := circle{2.5}
	fmt.Println("Area:", exampleCircle.area())
	fmt.Println("Permieter:", exampleCircle.perimeter())

	pointerToCircle := &exampleCircle
	fmt.Println("Area:", pointerToCircle.area())
	fmt.Println("Permieter:", pointerToCircle.perimeter())
}
