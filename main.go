package main

import "fmt"

func main() {

	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue //break out of current iteration of loop
		}

		if index > 2 {
			fmt.Println("breaking at pos", index)
			break //breaks out of the loop entirely
		}

		fmt.Printf("the value at pos %v is %v \n", index, value)
	}

}
