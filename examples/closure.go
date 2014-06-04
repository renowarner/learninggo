package main

import "fmt"

func main() {
	counterA := buildCounter("A", 5)
	counterB := buildCounter("B", 19)
	counterA()
	counterA()
	counterB()
	counterA()
	counterB()
}

func buildCounter(counterName string, stepping int) func() {
	counter := 1
	return func() {
		counter = counter * stepping
		fmt.Println("Counter ", counterName, " is: ", counter)
	}
}
