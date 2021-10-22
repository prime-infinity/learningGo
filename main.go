package main

import "fmt"

func circleArea(r float64) float64 {
	return 3.14 * r * r
}

func main() {

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1)
	fmt.Println(a2)

	fmt.Printf("circle 1 is %0.2f \n and circle 2 is %0.2f ", a1, a2)

}
