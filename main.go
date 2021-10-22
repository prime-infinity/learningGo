package main

import "fmt"

var score = 99.5 //to export, remember, cannot use := at package level

func main() {

	sayHello("mario")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()

}
