package main

import "fmt"

func main() {

	num := 400

	//if else statement
	if num%2 == 0 {
		fmt.Println("num is even number")
	} else {
		fmt.Println("num is odd number")
	}

	//if statement
	if num >= 100 {
		fmt.Println("num is 3 digit number")
	}

	//else if statement
	if neg := -3; neg > 0 {
		fmt.Println("Positive number")
	} else if neg < 0 {
		fmt.Println("Negative number")
	} else {
		fmt.Println("Zero")
	}

}
