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

	//passing anonymous function as argument
	a := func(name string) string {
		return name + " " + "say hello!"
	}

	greeting(a)

	//return anonymous function
	//function foo() returns a anonymous function
	boo := foo()
	//*if you just put str, it will print out address of boo
	fmt.Println(boo())
}

//anonymous function as parameter
func greeting(hey func(name string) string) {
	fmt.Println(hey("Jhon"))
	fmt.Println("Hello")
}

//return anonymous function
func foo() func() string {
	ano := func() string {
		return "boo"
	}
	return ano
}
