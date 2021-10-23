package main

import "fmt"

//a simple struct can mix any type of data together
type testStruct struct {
	number     int
	actorName  string
	companions []string
}

func main() {

	//structs are like advanced maps
	//structs are custom types

	cStruct := testStruct{
		number:    3,
		actorName: "mayiam",
		companions: []string{
			"this", "that", "tes",
		},
	}

	fmt.Println(cStruct)
	fmt.Println(cStruct.companions)

}
