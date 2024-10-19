package main

import (
	"fmt"
)

func dynamicstruct() {
	person := Person{
		Name: "Charlie Cartwright",
	}

  age := 999

	// Check if "Age" exists in dynamicValues and assign it
	if &age != nil {
		person.Age = age
	}

	fmt.Printf("%v\n", person)

}
