package main

import (
	"fmt"
)

type person struct {
	lName string
	fName string
}

type secretAgent struct {
	person        person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fName, p.lName, `says, Good morning!`)
}

func (sa1 secretAgent) speak() {
	fmt.Println(sa1.person.fName, sa1.person.lName, `says good morning`)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		"MoneyPenny",
		"Miss",
	}

	sa1 := secretAgent{
		person{
			"Bond",
			"James",
		},
		true,
	}
	saySomething(p)
	saySomething(sa1)
}
