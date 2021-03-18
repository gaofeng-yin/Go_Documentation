package main

import "fmt"

func main() {
	//function greet()
	greet()
	//function greetTo. "Yenni" is argument
	greetTo("Yenni")

	//you can also pass a variable as argument
	name := "Golang"

	greetTo(name)

	a := 3

	b := 5

	fmt.Println(sum(a, b))

	number := sum(10, 10)

	fmt.Println(number)
}

//how to define a function : func FUNCTIONAME(PARAMETER)RETURN{ STATEMENTS }

//function without parameter
func greet() {
	fmt.Println("Hello world!")
}

//function with parameter
func greetTo(name string) {
	fmt.Println("Hello", name)
}

//function with 2 parameter and returns a integer
func sum(a, b int) int {
	return a + b
}
