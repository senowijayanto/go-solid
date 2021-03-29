/***
** Liskov Substitution Principle :
** If S is a subtype of T, then objects of type T in a program may be replaced with objects of type S
** without altering any of the desirable properties of that program.
** Objects must be replaceable by instances of their subtypes without altering the correct functioning of our system
**/
package main

import "fmt"

type human struct {
	name string
}

func (h human) getName() string {
	return h.name
}

type teacher struct {
	human
	degree string
	salary float64
}

type student struct {
	human
	grades map[string]int
}

type person interface {
	getName() string
}

type printer struct{}

func (printer) info(p person) {
	fmt.Println("Name: ", p.getName())
}

func main() {
	fmt.Println("Liskov Substitution Principle")

	h := human{name: "Alex"}
	s := student{
		human: human{name: "Mike"},
		grades: map[string]int{
			"Math":    8,
			"English": 9,
		},
	}
	t := teacher{
		human:  human{name: "John"},
		degree: "CS",
		salary: 2000,
	}

	p := printer{}
	p.info(h)
	p.info(s)
	p.info(t)
}
