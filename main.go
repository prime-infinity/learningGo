package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	fmt.Println(opt)
}

func main() {

	myBill := createBill()

	promptOptions(myBill)

}
