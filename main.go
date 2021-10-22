package main

import "fmt"

func main() {

	statePopulations := map[string]int{ //key value pairs must be of the same type
		"califonia":    23423123,
		"texas":        2324232,
		"florida":      342323,
		"new youk":     32312423,
		"pennsylvania": 323124232,
		"illinois":     333232,
		"ohio":         12323232,
	}

	//fmt.Println(statePopulations)

	for index, value := range statePopulations {
		fmt.Println(index)
		fmt.Println(value)
	}

}
