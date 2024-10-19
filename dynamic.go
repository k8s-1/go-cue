package main

import (
	"fmt"
)

func dynamicstruct() {
	person := Person{
		Name: "Charlie Cartwright",
		Age:  999,
	}

  // Check if "Age" exists in dynamicValues and assign it
  if age, ok := dynamicValues["Age"]; ok {
      if ageValue, valid := age.(int); valid { // Ensure the type is correct
          person.Age = ageValue
      }
  }

	fmt.Printf("%v\n", person)

}
