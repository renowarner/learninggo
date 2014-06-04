package main

import "fmt"

func main() {
    x := 5
    zero(&x)
    fmt.Println(x) // x is 0
}

func zero(test *int) {
    *test = 0
}