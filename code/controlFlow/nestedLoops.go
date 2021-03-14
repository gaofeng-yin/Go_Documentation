package main

import (
	"fmt"
)

func main() {
	//slice of int
	xi := []int{2, 6, 3, 10, 8, 1, 5, 0}

	fmt.Println("unsorted:", xi)

	//nested for loop to sort slice of int
	for i := 0; i < len(xi); i++ {
		for j := i + 1; j < len(xi); j++ {
			if xi[i] > xi[j] {
				swap(&xi[i], &xi[j])
			}
		}
	}

	fmt.Println("sorted:", xi)
}

//function to swap 2 number
func swap(a, b *int) {
	temp := 0
	temp = *a
	*a = *b
	*b = temp
}
