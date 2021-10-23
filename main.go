package main

import "fmt"

//a simple struct can mix any type of data together

func main() {

	//structs are like advanced maps
	//structs are custom types

	myBill := newBill("osamede's bill")

	myBill.updateTip(10)
	myBill.addItems("sweet pie", 1000)
	myBill.addItems("pie", 10)
	myBill.addItems("onion soup", 100)
	myBill.addItems("cofee", 5)

	fmt.Println(myBill.format())

}
