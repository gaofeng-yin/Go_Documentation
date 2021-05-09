package main

import (
	"fmt"
	"strconv"
)

func main() {
	//output = hypothermia
	fmt.Println(temperature("34.0"))
}

func temperature(T string) string {
	//convert string to the float32 because temperature is decimal number
	f, err := strconv.ParseFloat(T, 32)
	//if there is a error during conversion, function will print the error and then return empty string
	if err != nil {
		fmt.Println(err)
		return ""
	}
	//switch instead of if because switch is more readable than nested if
	switch {
	case f < 35.0 && f >= 34.0: //*constraint: temperature should be bigger or equals to 34.0
		return "hypothermia"
	case f > 35.0 && f < 37.5:
		return "normal"
	case f > 37.5 && f < 40.0:
		return "fever"
	case f > 40.0 && f < 42: //*constraint: temperature should be smaller or equals to 42
		return "hyperpyrexia"
	default:
		return "Please insert a valid human temperature!" //if temperature == 43, it will be considered invalid
	}
}
