package main

import "fmt"

func main() {
	//initialize array of integer with zero value
	var num [5]int

	fmt.Println(num)

	//cap(num) return capacity of array
	fmt.Println(cap(num))

	//len(num) return length of array
	fmt.Println(len(num))

	//fill array with integer
	for i := 0; i < len(num); i++ {
		num[i] = i + 1
	}

	fmt.Println(num)

	fmt.Println()

	//in array we start counting from zero and so on.
	for index, value := range num {
		fmt.Println("Index", index, "value", value)
	}

	fmt.Println()

	//value at index 0
	fmt.Println("index 0", "value", num[0])
	//value at index 3
	fmt.Println("index 3", "value", num[3])

	fmt.Println()

	//create array in short hand
	names := [5]string{"James", "Tifanny", "Steven"}

	fmt.Println(names)

	names[3] = "George"

	names[4] = "Luther"

	fmt.Println(names)
}
