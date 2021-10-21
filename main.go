package main

import "fmt"

func main() {

	a := [...]int{1, 2, 3}
	b := a //copying arrays will literally copy the array

	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

}
