package main

import (
	_ "embed"
	"fmt"

	"cuelang.org/go/cue/cuecontext"
	// "cuelang.org/go/cue/load"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	//  ctx := cuecontext.New()
	// insts := load.Instances([]string{"."}, nil) // build max 1 package in current dir
	// v := ctx.BuildInstance(insts[0])
	// fmt.Printf("%v\n", v)
	ctx := cuecontext.New()

	person := Person{
		Name: "Charlie Cartwright",
		Age:  999,
	}

	personAsCUE := ctx.Encode(person)
	fmt.Printf("%personAsCUE\n", personAsCUE)
}
