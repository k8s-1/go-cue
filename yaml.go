package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"log"
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

	// Access specific keys (dynamic content)
	if name, ok := data["name"]; ok {
		fmt.Println("Name:", name)
	}
	if address, ok := data["address"].(map[interface{}]interface{}); ok {
		fmt.Println("Address:", address)
	}
	if hobbies, ok := data["hobbies"].([]interface{}); ok {
		fmt.Println("Hobbies:", hobbies)
	}
}
