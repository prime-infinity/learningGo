package main

import "fmt"

func main() {

	//slices
	var scores = []int{100, 42, 21, 43}

	//fmt.Println(scores)

	scores[2] = 332

	//fmt.Println(scores)

	scores = append(scores, 85) //append something to the end of a slice

	//fmt.Println(scores)

	//slice ranges
	rangeOne := scores[1:3]  //range of 1 to 2
	rangeTwo := scores[2:]   //from 2 to end
	rangeThree := scores[:3] //from start to before 3
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

}
