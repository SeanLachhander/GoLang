package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	age int
}

type fbiAgent struct {
	person
	ltk bool
}

func main() {
	person_one := person{
		first: "Sam",
		last: "Smith",
		age: 34,
	}
	person_two := person{
		first: "John",
		last: "Doe",
		age: 18,
	}
	sec_agent := fbiAgent{
		person: person{
			first: "Sean",
			last: "Lachhander",
			age: 23,
		},
		ltk: true,
	}

	fmt.Println(person_one)
	fmt.Println(person_two)
	fmt.Println(person_one.first, person_one.last, person_one.age)
	fmt.Println(person_two.first, person_two.last, person_two.age)
	fmt.Println(sec_agent.first, sec_agent.last, sec_agent.age, sec_agent.ltk)

}


