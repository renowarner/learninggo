package main

import "fmt"

var (
	a = "One"
	b = "Two"
	c = "Three"
	d = "Four"
)

func main() {
	fmt.Println("The current variables outside of scope are:", a, b, c, d)
	a := 1
	b := 2
	c := 3
	d := 4
	fmt.Println("However, inside this function they have been re defined numerically as such:")
	fmt.Println(a, b, c, d)
	fmt.Println("This allows us to combine them mathematically.")
	fmt.Println(a, "+", b, "+", c, "+", d, "=", (a + b + c + d))
	main2()
}
func main2() {
	fmt.Println("Now inside this second function, if we attempted the same we get a very different result...")
	fmt.Println(a + b + c + d)
	fmt.Println("This is because of Scope.  The letter format of all", d, "variables is declared outside of any function making it global.")
	fmt.Println("But by defining the variables with new values, we can temporarily use them as numerical values.")

	fmt.Print("This proccess is what is known as? ")
	var input string
	fmt.Scanf("%v", &input)

	if input == "Clobbering" {
		fmt.Println("That is correct!")
	} else {
		fmt.Println("wat? ._.  is you stoopid?")

	}
}
