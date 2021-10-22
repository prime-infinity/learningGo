package main

import "fmt"

func updateName(x string) {
	x = "wedge"
}

func main() {

	name := "tifa"

	updateName(name)

	//fmt.Println("memory address of name is:", &name) //referencing pointers

	m := &name //referencing pointers
	fmt.Println("memory address ", m)

	fmt.Println("value at memory address is :", *m)

	//fmt.Println(name)

}
