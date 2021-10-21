package main

import "fmt"

func main() {

	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j float32
	j = float32(i) //converts i to float and assigns it to float
	fmt.Printf("%v, %T\n", j, j)

}
