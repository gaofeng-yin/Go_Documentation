package main

import (
	"fmt"
	"strconv"
)

func main() {
	//output = hypothermia
	fmt.Println(Solution("34.0"))
}

func Solution(T string) string {
	f, err := strconv.ParseFloat(T, 32)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	switch {
	case f < 35.0:
		return "hypothermia"
	case f >= 35.0 && f < 37.5:
		return "normal"
	case f >= 37.5 && f < 40.0:
		return "fever"
	case f >= 40.0:
		return "hyperpyrexia"
	default:
		return ""
	}
}
