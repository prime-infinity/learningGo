package main

import "fmt"

//a simple struct can mix any type of data together

func main() {

	//structs are like advanced maps
	//structs are custom types

	myBill := newBill("osamede's bill")

	fmt.Println(myBill.format())

}
