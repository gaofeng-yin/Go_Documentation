package main

import (
	"fmt"
)

func main() {

	//basic switch -> switch expression { case condition: statements }
	a := 3

	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")

	}

	//switch without an expression
	number := 3452

	switch {
	case number > 0 && number < 100:
		fmt.Println("0-100")
		if number%2 == 0 {
			fmt.Println("even")
		}
	case number > 100 && number < 1000:
		fmt.Println("100-1000")
		if number%2 == 0 {
			fmt.Println("even")
		}
	case number > 1000 && number < 10000:
		fmt.Println("1000-10000")
		if number%2 == 0 {
			fmt.Println("even")
		}
	default:
		fmt.Println(">10000")
	}
}
