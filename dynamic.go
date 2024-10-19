package main

import (
	"fmt"
)

func dynamicstruct() {
	person := Person{
		Name: "Charlie Cartwright",
	}

  age := 99

	if age != 0 {
		person.Age = age
	}

	fmt.Printf("%v\n", person.Name)
	fmt.Printf("%v\n", person.Age)

}
