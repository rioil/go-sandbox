package main

import (
	"fmt"
)

type Printable interface {
	ToString() string
}

func PrintOutPrintable(p Printable) {
	fmt.Println(p.ToString())
}

func PrintOut(obj interface{}) {
	p, isPrintable := obj.(Printable)

	if isPrintable {
		PrintOutPrintable(p)
	} else {
		switch obj.(type) {
		case string:
			str := obj.(string)
			fmt.Println(str)
		default:
			fmt.Println("Can not print!")
		}
	}
}

type Person struct {
	name string
}

func (p Person) ToString() string {
	return p.name
}

type Book struct {
	title string
}

func (b Book) ToString() string {
	return b.title
}

func main() {
	person := Person{"Jonny"}
	book := Book{"Foo"}

	PrintOut(person)
	PrintOut(book)
	PrintOut("bar")
	PrintOut(1)
}
