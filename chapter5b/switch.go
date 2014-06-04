package main

import "fmt"

var c = "Advancing to"

func main() {
	for i := 1; i <= 11; i++ {
		switch i{
			case 0: fmt.Println("Zero")
			case 1: fmt.Println("One")
			case 2: fmt.Println("Two")
			case 3: fmt.Println("Three")
			case 4: fmt.Println("Four")
			case 5: fmt.Println("Five")
			case 6: fmt.Println("Six")
			case 7: fmt.Println("Seven")
			case 8: fmt.Println("Eight")
			case 9: fmt.Println("Nine")
			case 10: fmt.Println("Ten")
			case 11: endmsg()
			default: fmt.Println("Unknown Number")
		}
	}
}

func endmsg() {
	fmt.Println("You are all done counting, AND you have a better understanding of switches and functions then you thought.")
	
}