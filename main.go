package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("the price must be a number")
			promptOptions(b)
		}
		b.addItems(name, p)

		fmt.Println("item added - ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("enter tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("the tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("tip added - ", tip)
		promptOptions(b)
	case "s":
		fmt.Println("you choose save the bill", b)
	default:
		fmt.Println("that was not a valid option...")
		promptOptions(b)
	}
}

func main() {

	myBill := createBill()

	promptOptions(myBill)

}
