package main

import "fmt"

func main() {
	greet()
	greetTo("Yenni")

	fmt.Println(sum(2, 3))

}

func greet() {
	fmt.Println("Hello world!")
}

func greetTo(name string) {
	fmt.Println("Hello", name)
}

func sum(a, b int) int {
	return a + b
}

// Passing anonymous function as an argument(parameter)

// Returning anonymous function
