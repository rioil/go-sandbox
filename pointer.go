package main

import (
	"fmt"
)

func main() {
	var age int
	var pAge *int

	pAge = &age
	*pAge = 200
	fmt.Println(age)
	setValue(pAge, 45)
	fmt.Println(*pAge)
}

func setValue(target *int, value int) {
	*target = value
}
