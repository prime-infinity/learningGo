package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {

	s := strings.ToUpper(n)

	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1]) //strings can also be treated as strings, so range can be used on them
	}

	if len(initials) > 1 {
		return initials[0], initials[1] //return more than one values
	}
	return initials[0], "_"

}

func main() {

	fn1, sn1 := getInitials("tifa lockhart") //getting the two return statements from strings

	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("cloud strife")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("osamede")
	fmt.Println(fn3, sn3)

}
