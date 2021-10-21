package main

import "fmt"

func main() {

	/*x := 0
	for x < 5 {
		fmt.Print(x, "\n")
		x++
	}*/

	/*for i := 0; i < 5; i++ {
		fmt.Println(i)
	}*/

	names := []string{"osamede", "akonofua", "omokhomion", "this", "that"}

	/*for i := 0; i < len(names); i++ {

		fmt.Println(names[i])

	}*/

	for _, value := range names {
		fmt.Println(value)
	}

	//dont forget the break and continue and label
}
