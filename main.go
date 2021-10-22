package main

import "fmt"

func main() {

	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 { //if conditions do not need brackets in go
		fmt.Println("age is less than 30")
	} else if age > 40 {
		fmt.Println("age is greater than 40")
	} else {
		//default catch clause
	}

}
