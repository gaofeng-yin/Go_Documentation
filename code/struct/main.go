package main

import "fmt"

//person strcut
type person struct {
	first string
	last  string
	age   int
}

//constructor for person. Create person with given firstName, lastName and agePerson. It returns a pointer.
func newPerson(firstName, lastName string, agePerson int) *person {
	p := person{
		first: firstName,
		last:  lastName,
		age:   agePerson,
	}
	return &p
}

//person function to get fullname. Only accessible for type person.
func (p person) fullname() string {
	return p.first + " " + p.last
}

//to modify person.age we use pointer or it won't be changed
func (p *person) incAge(i int) {
	p.age += i
}

func main() {

	//long way to create person
	p1 := person{
		first: "George",
		last:  "Roman",
		age:   45,
	}

	//short way to create person
	p2 := person{"Cristiano", "Messi", 34}

	//create person using constructor
	p3 := newPerson("Steve", "Liu", 24)

	//p1 increment age by 10
	p1.incAge(10)

	//print out person info
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	//get person full name
	fmt.Println(p1.fullname())
	fmt.Println(p2.fullname())
	fmt.Println(p3.fullname())

}
