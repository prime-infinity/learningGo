package main

import "fmt"

func main() {

	//maps are like objects in other languages, but like really strict objects

	statePopulations := map[string]int{ //key value pairs must be of the same type
		"califonia":    23423123,
		"texas":        2324232,
		"florida":      342323,
		"new youk":     32312423,
		"pennsylvania": 323124232,
		"illinois":     333232,
		"ohio":         12323232,
	}

	statePopulations["califonia"] = 000000     //updating a value
	fmt.Println(statePopulations["califonia"]) //getting a single value from a key

	delete(statePopulations, "califonia") //deleting from map

	_, ok := statePopulations["new"] //cheching for the existence of a key, value pair
	fmt.Println(ok)

	//note, return order of a map is not guranteed

	/*for index, value := range statePopulations {
		fmt.Println(index)
		fmt.Println(value)
	}*/

}
