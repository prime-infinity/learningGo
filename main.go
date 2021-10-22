package main

import "fmt"

func sayGreeting(n string) {
	fmt.Printf("good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {

	/*for i := 0; i < len(n); i++ {

	}*/ //the old way

	for _, v := range n {
		f(v)
	}

}

func main() {

	/*sayGreeting("mario")
	sayBye("luig")
	sayGreeting("osamede")*/

	sli := []string{"cloud", "tifa", "barret"}

	cycleNames(sli, sayGreeting) //the "saygereeting function is not invoked here, but inside the func callingit"
}
