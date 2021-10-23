package main

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
