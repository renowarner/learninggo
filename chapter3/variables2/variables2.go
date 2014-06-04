package main

import "fmt"

const tag string = "is the current variable. "

func main() {
	var Var string
	Var = "First"
	fmt.Println(Var, tag)
	Var = "Second"
	fmt.Println(Var, tag)
}
