package main

import "fmt"

var(
	fahr = "What is the current temperature in Fahrenheit?: "
	cels = "The current temperature in Celcius is: "
)

func main() {
	fmt.Print(fahr)
	var input float64
	fmt.Scanf("%f", &input)
	
	output := (input - 32) * 5/9
	
	fmt.Println(cels,output)
	
}