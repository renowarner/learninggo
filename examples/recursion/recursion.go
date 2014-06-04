package main

import "fmt"

func main() {
	countDownFrom(7)
	fmt.Println("Done")
}

func countDownFrom(n int) int {
	fmt.Println("Counting down, current number is:", n)
	if n == 0 {
		return n
	}
	return countDownFrom(n - 1)
}

// A function that calls itself.
