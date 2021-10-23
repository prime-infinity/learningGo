package main

import "fmt"

//note that the "biil" is now a type onits own and can be used as return in functions
type bill struct {
	name   string
	items  map[string]float64
	tip    float64
	status bool
}

//when defining a struct, the name and the type is needed

//make new bills function
func newBill(name string) bill {

	b := bill{
		name:   name,
		items:  map[string]float64{"pie": 5.99, "cake": 3.94},
		tip:    0,
		status: true,
	}

	return b

}

//a custom function that applies only to our cutsom type. #receiver functions
func (b bill) format() string {
	fs := "Biill breakdown: \n"

	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ... $%v \n", k+":", v)
		total += v
	}

	//total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)
	return fs
}

//the above functions should be implemented again using known methods
