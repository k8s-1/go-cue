package main

import (
	"fmt"
)

func dynamicstruct() {
	person := Person{
		Name: "Charlie Cartwright",
		Age:  999,
	}

	fmt.Printf("%v\n", person)

}
