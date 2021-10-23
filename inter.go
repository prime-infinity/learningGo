package main

import (
	"fmt"
	"math"
)

// an interface groups types together based on their methods
type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

//square methods
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

//circle mehtods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("area of %T is : %0.2f \n", s, s.area())
	fmt.Printf("circumfrence of %T is :%0.2f \n", s, s.circumf())
}

func main() {
	shapes := []shape{
		square{length: 15.3},
		square{length: 12.3},
		circle{radius: 31.2},
		circle{radius: 331.1},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}
}
