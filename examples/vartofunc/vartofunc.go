package main

import "fmt"

func main() {
	hello := buildCrap("Hello")
	d := buildCrap("ddd")
	x := buildCrap("xXx")
	x()
	d()
	hello()
}

func buildCrap(crap string) func() {
	return func() {
		fmt.Println("You know what is crap: " + crap)
	}
}
