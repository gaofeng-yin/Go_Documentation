package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func newPerson(firstName, lastName string, agePerson int) *person {
	p := person{
		first: firstName,
		last:  lastName,
		age:   agePerson,
	}
	return &p
}

func (p person) fullname() {
	fmt.Println("Full name:", p.first, p.last)
}

func main() {

	p1 := person{
		first: "George",
		last:  "Roman",
		age:   45,
	}

	p2 := person{"Cristiano", "Messi", 34}

	fmt.Println(p1)
	fmt.Println(p2)

	p1.fullname()
	p2.fullname()

}
