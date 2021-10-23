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

	/*fmt.Print("create a new bill name: ")
	name = strings.TrimSpace(name)*/

	name, _ := getInput("create a new bill name:   ", reader)

	b := newBill(name)
	fmt.Println("created the bill - ", b.name)

	return b

}

func main() {

	myBill := createBill()

	fmt.Println(myBill)

}
