package main

import "fmt"
// Variable option 1 Start
var(
	feet = "How many feet would you like to convert?"
	meterValue = 0.3048
	outCap = "feet, is equal to"
	outShoe = "meters."
)
// Variable option 1 End

func main(){
	fmt.Print(feet)
	// Variable option 2 Start
	var input float64
	// Variable option 2 End
	fmt.Scanf("%f", &input)
	// Variable option 3 Start
	output := input * meterValue
	// Variable option 3 End
	
	fmt.Println(input,outCap,output,outShoe)
}