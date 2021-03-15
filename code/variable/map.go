package main

import "fmt"

func main() {

	//create atom map with built-in function make() -> make(map[key-type]val-type)
	atom := make(map[string]string)

	//set key/value
	atom["H"] = "Hydrogen"
	atom["He"] = "Helium"
	atom["Li"] = "Lithium"
	atom["Be"] = "Beryllium"
	atom["B"] = "Boron"

	fmt.Println(atom)

	fmt.Println(atom["H"])

	//get value from key
	H := atom["H"]

	fmt.Println(H)

	//built-in function delete() will remove key/value pairs from a map -> delete(map, key)
	delete(atom, "H")

	fmt.Println("after deleting Hydrohen", atom)

	fmt.Println()

	//other way to create a map
	crypto := map[string]string{"BTC": "bitcoin", "ETH": "ethereum", "LINK": "chainlink", "ADA": "cardano"}

	fmt.Println(crypto)

}
