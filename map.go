package main

import (
	"fmt"
)

func main() {
	members := map[string]int{
		"John": 20,
		"Teddy": 32,
		"Tommy": 13,
	}

	_, isMember := members["John"]
	if isMember {
		fmt.Println("John is a member")
	} else {
		fmt.Println("John is not a member")
	}

	var format string
	if len(members) > 1 {
		format = "There are %d members\n"
	}else{
		format = "There is %d member\n"
	}

	fmt.Printf(format, len(members))
	for name, age := range members {
		fmt.Printf("%s (%d)\n", name, age)
	}
}
