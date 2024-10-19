package main

import (
	"cuelang.org/go/cue/cuecontext"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func parse() {
	// Read the YAML file
	content, err := os.ReadFile("values.yaml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))

	// Declare a generic map to hold the YAML data
	var data map[string]interface{}

	// Unmarshal the YAML into the map
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	// Print the parsed data
	fmt.Printf("Parsed YAML data: %+v\n", data)

	// Access specific keys without caring about types
	for key, value := range data {
		fmt.Printf("%s: %v\n", key, value) // %v will print the value in a generic way
	}

	// Access specific keys (dynamic content)
	if name, ok := data["name"]; ok {
		fmt.Println("Name:", name)
	}

	if address, ok := data["address"]; ok {
		fmt.Println("Address:", address)
	}

	if hobbies, ok := data["hobbies"]; ok {
		fmt.Println("Hobbies:", hobbies)
	}

	// Step 3: Construct CUE content dynamically
	cueContent := `import (
    "github.com/k8s-1/app/best"
)

#objects: best.#Def & {
    #vals: {
        `

	ctx := cuecontext.New()

	person := Person{
		Name: "Charlie Cartwright",
		Age:  999,
	}

	personAsCUE := ctx.Encode(person)
	cueContent += fmt.Sprintf("%v\n", personAsCUE)

	cueContent += `
    }
}
yaml.MarshalStream([ for _, o in #objects {o} ])
`
	fmt.Printf("CUE file: %+v\n", cueContent)

}
