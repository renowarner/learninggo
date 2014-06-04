package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("The number", i, "is even.")
		} else {
			fmt.Println("The number", i, "is odd.")
		}
	}
}
