package main

import "fmt"

var (
	i = 1000
)

func main() {
	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}
}
