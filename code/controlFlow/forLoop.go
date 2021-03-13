package main

import "fmt"

func main() {

	//for loop (initialization; termination; increment)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	//theres is no while or do while in Go. You can do it with For loop
	j := 0

	for j != 5 {
		fmt.Printf("j = %d, which is less than 5\n", j)
		j++
	}

	fmt.Println()

	//range to iterate over slice of int
	xi := []int{10, 20, 30, 40, 50}

	for i, v := range xi {
		fmt.Printf("Index %d \t\t Value %d\n", i, v)
	}

	fmt.Println()

	//range to iterate over map[key]value
	xii := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	for key, value := range xii {
		fmt.Printf("Key %d \t\t\t Value %s\n", key, value)
	}

	//infinite loop
	/*
		for {
			fmt.Println("Infinite loop!!!")
		}
	*/

}
