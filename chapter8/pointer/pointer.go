package main

import "fmt"

func zero(POINTERNAME *int) {
    *POINTERNAME = 0
}
func main() {
    x := 5
    zero(&x)
    fmt.Println(x)
}