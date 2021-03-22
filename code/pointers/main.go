package main

import "fmt"

func main() {

	phrase := "Hello world!"

	//ptr is a pointer to phrase address
	var ptr *string = &phrase

	fmt.Printf("%s\n", phrase)

	changeValue(phrase)

	fmt.Printf("%s\n", phrase)

	changeValuePointer(ptr)

	fmt.Printf("%s\n", phrase)

	fmt.Println("address:", ptr)
	fmt.Println("value pointed by address:", *ptr)

}

//changeValue has a string paramenter, so arguments will be passed to it by value. Function will get a copy of str instead of original variable.
func changeValue(str string) {
	str = "String updated"
}

//changeValuePointer has a *string paramenter which means a string pointer. Now funtcion will get address of the variable and we can manipulate their value.
func changeValuePointer(str *string) {
	*str = "String updated"
}
