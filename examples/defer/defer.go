package main

import "fmt"

func main() {
	defer printCrap("3") // adds to stack of things to do at end of function
	defer printCrap("2") // adds this on top of the stack, so it will be done first
	defer printCrap("1") // add this on top of that last one, so it will be done first
	printCrap("Hello")
	// when a function is over, it runs all the defers
}

func printCrap(crap string) {
	fmt.Println("Crap is:", crap)
}
