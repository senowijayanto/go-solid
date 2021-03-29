package main

import (
	"fmt"
	"math"
)

// circle struct is a structure of circle with radius properties
type circle struct {
	radius float64
}

// area method with receiver of circle is a method to count the area of circle
func (c circle) area() {
	fmt.Printf("circle area: %f\n", math.Pi*c.radius*c.radius)
}

// square struct is a structure of square with length properties
type square struct {
	length float64
}

// area method with receiver of square is a mehtod to count area of square
func (s square) area() {
	fmt.Printf("square area: %f\n", s.length*s.length)
}

func main() {
	fmt.Println("Single Responsibility Principle Before")

	c := circle{radius: 5}
	c.area()

	s := square{length: 7}
	s.area()
}
