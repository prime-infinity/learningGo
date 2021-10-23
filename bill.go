package main

import (
	"fmt"
	"os"
)

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
		items:  map[string]float64{},
		tip:    0,
		status: true,
	}

	return b

}

//a custom function that applies only to our cutsom type. #receiver functions
func (b *bill) format() string {
	fs := "Biill breakdown: \n"

	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ... $%v \n", k+":", v)
		total += v
	}

	//add tip
	fs += fmt.Sprintf("%-25v ... $%v \n", "tip:", b.tip)

	//total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)
	return fs
}

//the above functions should be implemented again using known methods

//function to update tip
/*func (b bill) updateTip(t float64) {

	b.tip = t //this wil not work cus the copy is updated, not the actual its
}*/

func (b *bill) updateTip(t float64) {

	b.tip = t

}

//function to add items to bill struct
func (b *bill) addItems(name string, price float64) {
	b.items[name] = price
}

//save bill
func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bill/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("bill was saved to file")
}
