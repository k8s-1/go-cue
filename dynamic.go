package main

import (
	"fmt"
)

func dynamicstruct() {
	person := Person{
		Name: "Charlie Cartwright",
	}

  age := 99

	// Check if "Age" exists in dynamicValues and assign it
	if age != 0 {
		person.Age = age
	}

	fmt.Printf("%v\n", person)

}
