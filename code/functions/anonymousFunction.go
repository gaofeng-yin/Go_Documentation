package main

import "fmt"

func main() {
	// Anonymous function
	func() {
		fmt.Println("Anonymous function: Hello world!")
	}()

	//assign a anonymous function to variable
	greet := func() {
		fmt.Println("Hello world!")
	}

	fmt.Printf("variable greet is type: %T\n", greet)

	greet()

	//anonymous functions with parameter
	func(name string) {
		fmt.Println("Welcome", name)
	}("James")
}
