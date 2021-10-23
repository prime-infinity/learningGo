package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//remenber that user inputs are always stings
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	prompt, err := r.ReadString('\n')
	return strings.TrimSpace(prompt), err
}

func createBill() bill {

	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("create a new bill name:   ", reader)

	b := newBill(name)
	fmt.Println("created the bill - ", b.name)

	return b

}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("item name:", reader)
		price, _ := getInput("item price:", reader)
		fmt.Println(name, price)
	case "t":
		tip, _ := getInput("enter tip amount ($): ", reader)
		fmt.Println(tip)
	case "s":
		fmt.Println("you choose s")
	default:
		fmt.Println("that was not a valid option...")
		promptOptions(b)
	}
}

func main() {

	myBill := createBill()

	promptOptions(myBill)

}
