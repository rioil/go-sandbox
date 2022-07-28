package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p *Person) SetPerson(name string, age int) {
	p.name = name
	p.age = age
}

func (p *Person) GetPerson() (string, int) {
	return p.name, p.age
}

func main() {
	var p1 Person
	p1.SetPerson("Yamada", 26)
	name, age := p1.GetPerson()
	fmt.Printf("%s %d\n", name, age)

	p2 := Person{"Inoue", 65}
	name, age = p2.GetPerson()
	fmt.Printf("%s %d\n", name, age)
}
