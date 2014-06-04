package main

import "fmt"

func main() {
	x := [20]string{
		"i", "do", "house", "with", "mouse",
		"in", "not", "like", "them", "ham",
		"a", "anywhere", "green", "eggs", "and",
		"here", "or", "there", "sam", "am",
	}
	fmt.Println(x[0], x[1], x[6], x[7], x[8], x[5], x[10], x[2])
}
