package main

import "fmt"

func main() {
	x := buildCrap("Hello")
	fmt.Println(x)
}

func buildCrap(crap string) string {
	return "You know what is crap: " + crap
}